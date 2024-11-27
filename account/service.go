package user_service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bxcodec/go-clean-arch/config"
	account_repository "github.com/bxcodec/go-clean-arch/internal/repository/account"
	"github.com/bxcodec/go-clean-arch/params"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	bybit "github.com/wuhewuhe/bybit.go.api"
	"log"
	"net/http"
	"time"
)

type AccountServiceRepository interface {
	FindByUsernameAndPassword(ctx context.Context, username string, password string) (params.UserInfo, error)
}

type UserService struct {
	UserRepo    account_repository.AccountRepository
	config      config.Config
	ByBitClient *bybit.Client
	StopChannel chan bool
}

func NewUserService(cfg config.Config, user account_repository.AccountRepository) UserService {
	return UserService{
		UserRepo:    user,
		ByBitClient: bybit.NewBybitHttpClient(cfg.ByBit.ApiKey, cfg.ByBit.ApiSecret, bybit.WithBaseURL(bybit.MAINNET)),
	}
}

func (h UserService) SetRoutes(e *echo.Echo) {

	userGroup := e.Group("/orderbook")
	//userGroup.GET("/GetAccountWallet", h.GetAccountWallet)
	userGroup.GET("/StartSubscribe", h.StartSubscribe)
	userGroup.GET("/StopSubscribe", h.StopSubscribe)

	//userGroup.Use(middleware.Protected(h.authConfig))

	//userGroup.Post("/all", middleware.Protected(h.authConfig, h.authSvc),
	//	middleware.AccessCheck(h.userSvc, h.authSvc), h.getAll)
	//userGroup.Post("/register", h.userRegister)
}

func (s *UserService) StopSubscribe(ctx echo.Context) error {
	inputOrderBook := params.InputOrderBook{
		Op:   "subscribe",
		Args: []string{"orderbook.1.BTCUSDT"},
	}
	defer handlePanic()
	s.StopChannel <- true
	return ctx.JSON(http.StatusOK, "Stop Subscribe form `"+inputOrderBook.Args[0]+" Args` "+time.Now().Format("2006-01-02 15:04:05"))
}
func (s *UserService) StartSubscribe(ctx echo.Context) error {
	inputOrderBook := params.InputOrderBook{
		Op:   "subscribe",
		Args: []string{"orderbook.1.BTCUSDT"},
	}
	s.StopChannel = make(chan bool)
	defer handlePanic()
	go getOrderBook(inputOrderBook, s.StopChannel)
	return ctx.JSON(http.StatusOK, "Started Subscribe form `"+inputOrderBook.Args[0]+" Args` "+time.Now().Format("2006-01-02 15:04:05"))

}

func getOrderBook(iob params.InputOrderBook, stopChannel chan bool) {
	inputMsg, err2 := json.Marshal(iob)
	if err2 != nil {
		log.Println("Marshal err:", err2)
	}
	c, _, err := websocket.DefaultDialer.Dial("wss://stream.bybit.com/v5/public/spot", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	c.WriteMessage(websocket.TextMessage, inputMsg)

	defer c.Close()
	var snapShot params.SpotOrderBook
	for {
		select {
		case <-stopChannel:
			fmt.Println("stopping my goroutine")
			return
		default:
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read :", err)
			}
			if len(message) > 0 {
				var delta params.SpotOrderBook
				err = json.Unmarshal(message, &delta)
				if err != nil {
					log.Println("unmarshal err:", err)
				}
				if delta.Type == "snapshot" {
					log.Printf("recv: %s\n", message)
					snapShot = delta
				} else if delta.Type == "delta" {
					snapShot.UpdateSnapShot(delta.Data)
					log.Printf("recv: %s\n", snapShot)
				}
			}
			//fmt.Println("My Goroutine is running :( ")
			//time.Sleep(time.Second)
		}
	}
}

func handlePanic() {
	if panicInfo := recover(); panicInfo != nil {
		fmt.Println(panicInfo)
	}
}
func (s *UserService) GetAccountWallet(ctx echo.Context) error {

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
