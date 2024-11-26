package user_service

import (
	"context"
	"fmt"
	"github.com/bxcodec/go-clean-arch/config"
	account_repository "github.com/bxcodec/go-clean-arch/internal/repository/account"
	"github.com/bxcodec/go-clean-arch/params"
	"github.com/labstack/echo/v4"
	bybit "github.com/wuhewuhe/bybit.go.api"
	"net/http"
)

type AccountServiceRepository interface {
	FindByUsernameAndPassword(ctx context.Context, username string, password string) (params.UserInfo, error)
}

type UserService struct {
	UserRepo    account_repository.AccountRepository
	config      config.Config
	ByBitClient *bybit.Client
}

func NewUserService(cfg config.Config, user account_repository.AccountRepository) UserService {
	return UserService{
		UserRepo:    user,
		ByBitClient: bybit.NewBybitHttpClient(cfg.ByBit.ApiKey, cfg.ByBit.ApiSecret, bybit.WithBaseURL(bybit.MAINNET)),
	}
}

func (h UserService) SetRoutes(e *echo.Echo) {

	userGroup := e.Group("/users")
	userGroup.GET("/GetAccountWallet", h.GetAccountWallet)

	//userGroup.Use(middleware.Protected(h.authConfig))

	//userGroup.Post("/all", middleware.Protected(h.authConfig, h.authSvc),
	//	middleware.AccessCheck(h.userSvc, h.authSvc), h.getAll)
	//userGroup.Post("/register", h.userRegister)
}

func (s *UserService) GetAccountWallet(ctx echo.Context) error {

	s.ByBitClient.Debug = true
	param := make(map[string]interface{})
	param["accountType"] = "CONTRACT"
	bybit.WithDebug(true)

	accountResult, err := s.ByBitClient.NewClassicalBybitServiceWithParams(
		map[string]interface{}{"accountType": "CONTRACT"}).
		GetAccountWallet(context.Background())

	//userFound, err := s.UserRepo.FindByUsernameAndPassword(ctx, username, password)
	//if err != nil {
	//
	//}
	if err != nil {
		fmt.Println(err)
	}
	return ctx.JSON(http.StatusOK, accountResult)
}
