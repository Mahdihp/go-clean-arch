package delivery

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-clean-arch/config"
	"go-clean-arch/internal/bybit_history_service"
	"go.uber.org/zap"
)

type HistoricalServer struct {
	config  config.Config
	service bybit_history_service.ByBitHistoricalServic
	Router  *echo.Echo
}

func NewHistoricalServer(cfg config.Config, service bybit_history_service.ByBitHistoricalServic) HistoricalServer {
	return HistoricalServer{
		config:  cfg,
		Router:  echo.New(),
		service: service,
	}
}

func (s HistoricalServer) Serve() {
	// Middleware

	logger, _ := zap.NewProduction()
	s.Router.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info("request",
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
			)

			return nil
		},
	}))

	// Routes
	s.Router.GET("/health-check", healthCheck)
	s.service.SetRoutes(s.Router)

	s.Router.Use(middleware.RequestID())
	s.Router.Use(middleware.Recover())

	// Start server
	address := fmt.Sprintf(":%d", s.config.HistoricalServer.HttpPort)
	fmt.Printf("start echo server on %s\n", address)
	if err := s.Router.Start(address); err != nil {
		fmt.Println("router start error", err)
	}
}
