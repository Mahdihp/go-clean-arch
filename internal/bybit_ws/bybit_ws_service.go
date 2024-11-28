package user_service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bxcodec/go-clean-arch/config"
	"github.com/bxcodec/go-clean-arch/internal/bybit_ws/params"
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
	snapShot    params.SpotOrderBook
	stopChan    chan bool
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
func (s *ByBitWSService) startSpotSubscribe(ctx echo.Context) error {
	inputOrderBook := params.InputOrderBook{
		Op:   "subscribe",
		Args: []string{"orderbook." + ctx.QueryParam("count") + "." + ctx.QueryParam("symbol")},
	}
	s.stopChan = make(chan bool)
	s.snapShot = params.SpotOrderBook{}
	go s.sendRequest(inputOrderBook)
	return ctx.JSON(http.StatusOK, inputOrderBook)
}

func (s *ByBitWSService) stopSpotSubscribe(ctx echo.Context) error {
	s.stopChan <- true
	return ctx.JSON(http.StatusOK, "Stop Subscribe form "+time.Now().Format("2006-01-02 15:04:05"))
}

func (s *ByBitWSService) sendRequest(iob params.InputOrderBook) {
	wsConn, _, err := websocket.DefaultDialer.Dial(s.config.ByBit.WsSocketSpot, nil)
	if err != nil {
		log.Println("WebSocket Upgrade Error: ", err)
		return
	}
	defer wsConn.Close()
	wsConn.SetCloseHandler(func(code int, text string) error {
		log.Printf("Received connection close request. Closing connection .....")
		return errors.New("Received connection close request. Closing connection .....")
	})
	inputMsg := marshalOperatore(iob)
	err = wsConn.WriteMessage(websocket.TextMessage, inputMsg)
	if err != nil {
		log.Println("Error writing message to WS connection: ", err)
		return
	}
	for {
		select {
		case <-s.stopChan:
			fmt.Println("Thanks for stopping my goroutine :)")
			return
		default:
			s.readData(wsConn)
		}
		time.Sleep(150 * time.Microsecond)
	}
}

func (s *ByBitWSService) readData(c *websocket.Conn) {
	mt, message, err := c.ReadMessage()
	if err != nil {
		log.Println("read :", err, mt)
	}
	//log.Printf("recv: %s", message)
	if len(message) > 0 {
		var delta params.SpotOrderBook
		err = json.Unmarshal(message, &delta)
		if err != nil {
			log.Println("unmarshal err:", err)
		}
		if delta.Type == "snapshot" {
			log.Printf("recv: %s\n", message)
			s.snapShot = delta
			log.Printf("recv snapShot: %d %d \n", len(s.snapShot.Data.A), len(s.snapShot.Data.B))
		} else if delta.Type == "delta" {
			s.snapShot.UpdateSnapShot(delta.Data, 1000)
			log.Printf("recv: %d %d \n", len(s.snapShot.Data.A), len(s.snapShot.Data.B))
			//log.Printf("recv2: %d %d \n", len(snapShot.Data.A), len(snapShot.Data.B))
			//snapShot.UpdateSnapShot(delta.Data, 1000)
			//marshal, _ := json.Marshal(delta)
			//log.Printf("recv: %s", marshal)

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

	//s.ByBitClient.Debug = true
	//bybit.WithDebug(true)
	//accountResult, err := s.ByBitClient.NewClassicalBybitServiceWithParams(
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
