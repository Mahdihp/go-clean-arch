package delivery

import (
	"fmt"
	"github.com/bxcodec/go-clean-arch/config"
	"github.com/bxcodec/go-clean-arch/internal/bybit_ws"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type Server struct {
	config         config.Config
	bybitWsService user_service.ByBitWSService
	Router         *echo.Echo
}

func NewServer(cfg config.Config, userService user_service.ByBitWSService) Server {
	return Server{
		config:         cfg,
		Router:         echo.New(),
		bybitWsService: userService,
	}
}

func (s Server) Serve() {
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
	s.Router.GET("/health-check", s.healthCheck)
	s.bybitWsService.SetRoutes(s.Router)

	s.Router.Use(middleware.RequestID())
	s.Router.Use(middleware.Recover())

	// Start server
	address := fmt.Sprintf(":%d", s.config.Server.HttpPort)
	fmt.Printf("start echo server on %s\n", address)
	if err := s.Router.Start(address); err != nil {
		fmt.Println("router start error", err)
	}
}
