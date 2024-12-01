package user_service

import (
	"encoding/json"
	"fmt"
	"github.com/bxcodec/go-clean-arch/config"
	"github.com/bxcodec/go-clean-arch/internal/bybit_ws/params"
	"github.com/bxcodec/go-clean-arch/internal/bybit_ws/pkg"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	bybit "github.com/wuhewuhe/bybit.go.api"
	"log"
	"net/http"
	"time"
)

type ByBitWSService struct {
	config      config.Config
	ByBitClient *bybit.Client
	snapShot    *params.SpotOrderBook
	stoperChan  chan bool
	isErrorChan chan bool
}

func NewByBitWSService(cfg config.Config) ByBitWSService {
	return ByBitWSService{
		config:      cfg,
		ByBitClient: bybit.NewBybitHttpClient(cfg.ByBit.ApiKey, cfg.ByBit.ApiSecret, bybit.WithBaseURL(bybit.MAINNET)),
	}
}

func (s *ByBitWSService) SetRoutes(e *echo.Echo) {

	userGroup := e.Group("/orderbook")
	//userGroup.GET("/GetAccountWallet", s.GetAccountWallet)
	userGroup.GET("/start_spot_subscribe", s.startSpotSubscribe)
	userGroup.GET("/stop_spot_subscribe", s.stopSpotSubscribe)
}
func (s *ByBitWSService) stopSpotSubscribe(ctx echo.Context) error {
	s.stoperChan <- true
	return ctx.JSON(http.StatusOK, "Stop Subscribe form "+time.Now().Format("2006-01-02 15:04:05"))
}

func (s *ByBitWSService) startSpotSubscribe(ctx echo.Context) error {
	inputOrderBook := params.InputOrderBook{
		Op:   "subscribe",
		Args: []string{"orderbook." + ctx.QueryParam("count") + "." + ctx.QueryParam("symbol")},
	}
	s.stoperChan = make(chan bool, 1)
	s.isErrorChan = make(chan bool, 1)
	s.snapShot = new(params.SpotOrderBook)
	go s.sendRequest(inputOrderBook)
	go func() {
		for {
			select {
			case <-s.isErrorChan:
				fmt.Println("new worker Running..." + time.Now().Format("2006-01-02 15:04:05"))
				if <-s.stoperChan == false {
					go s.sendRequest(inputOrderBook)
				}
			case <-s.stoperChan:
				fmt.Println("Stop worker Running...", time.Now().Format("2006-01-02 15:04:05"))
				return
			default:
				fmt.Println("worker Running...", time.Now().Format("2006-01-02 15:04:05"), <-s.stoperChan)
				time.Sleep(3 * time.Second)
			}
		}
	}()
	return ctx.JSON(http.StatusOK, inputOrderBook)
}

func (s *ByBitWSService) sendRequest(iob params.InputOrderBook) {
	wsConn, _, err := websocket.DefaultDialer.Dial(s.config.ByBit.WsSocketSpot, nil)
	if err != nil {
		log.Println("WebSocket Upgrade Error: ", err)
		s.isErrorChan <- true
	}
	defer wsConn.Close()
	wsConn.SetCloseHandler(func(code int, text string) error {
		log.Printf("Received connection close request. Closing connection .....")
		s.isErrorChan <- true
		return err
	})
	inputMsg := marshalOperatore(iob)
	err = wsConn.WriteMessage(websocket.TextMessage, inputMsg)
	if err != nil {
		log.Println("Error writing message to WS connection: ", err)
		s.isErrorChan <- true
	}
	for {
		select {
		case <-s.isErrorChan:
			return
		case <-s.stoperChan:
			fmt.Println("Stop sendRequest...")
			return
		default:
			s.readData(wsConn)
		}
		time.Sleep(1 * time.Microsecond)
	}
}

func (s *ByBitWSService) readData(c *websocket.Conn) {
	mt, message, err := c.ReadMessage()
	if err != nil {
		log.Println("read :", err, mt)
		s.isErrorChan <- true
	}
	if len(message) > 0 {
		delta := params.SpotOrderBook{}
		err = json.Unmarshal(message, &delta)
		if err != nil {
			log.Println("unmarshal err:", err)
			s.isErrorChan <- true
		}
		if delta.Type == "snapshot" {
			log.Printf("Type: %s\n", delta.Type)
			s.snapShot = &delta
			log.Printf("Type: %s\n", len(s.snapShot.Data.B), len(s.snapShot.Data.A))
		} else if delta.Type == "delta" {
			UpdateSnapShot(s.snapShot, delta.Data, 1000)
			//log.Printf("recv: %d %d \n", len(s.snapShot.Data.A), len(s.snapShot.Data.B))
			marshal, _ := json.Marshal(s.snapShot)
			log.Printf("recv: %s \n", marshal)

		}

	}
}
func marshalOperatore(inputOrderBook params.InputOrderBook) []byte {
	inputMsg, err2 := json.Marshal(inputOrderBook)
	if err2 != nil {
		log.Println("Marshal err:", err2)
	}
	return inputMsg
}
func (s *ByBitWSService) GetAccountWallet(ctx echo.Context) error {

	//s.byBitClient.Debug = true
	//bybit.WithDebug(true)
	//accountResult, err := s.byBitClient.NewClassicalBybitServiceWithParams(
	//	map[string]interface{}{"accountType": "CONTRACT"}).
	//	GetAccountWallet(context.Background())

	//userFound, err := s.UserRepo.FindByUsernameAndPassword(ctx, username, password)
	//if err != nil {
	//
	//}
	//if err != nil {
	//	fmt.Println(err)
	//}

	return ctx.JSON(http.StatusOK, "accountResult")
}

func UpdateSnapShot(snapShot *params.SpotOrderBook, data params.BidAsk, limitOrderBook int) {
	pkg.UpdateBids(&snapShot.Data, data)
	pkg.UpdateAsks(&snapShot.Data, data)
	pkg.RemoveDuplicatesBids(&snapShot.Data)
	pkg.RemoveDuplicatesAsks(&snapShot.Data)

	if len(snapShot.Data.A) > limitOrderBook {
		snapShot.Data.A = snapShot.Data.A[:limitOrderBook]
	}
	if len(snapShot.Data.B) > limitOrderBook {
		snapShot.Data.B = snapShot.Data.B[:limitOrderBook]
	}
}
