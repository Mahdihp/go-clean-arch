package delivery

import (
	"fmt"
	user_service "github.com/bxcodec/go-clean-arch/account"
	"github.com/bxcodec/go-clean-arch/config"
	user_validator "github.com/bxcodec/go-clean-arch/internal/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type Server struct {
	config        config.Config
	userValidator user_validator.UserValidator
	userService   user_service.UserService
	Router        *echo.Echo
}

func NewServer(cfg config.Config, userValidator user_validator.UserValidator, userService user_service.UserService) Server {
	return Server{
		config:        cfg,
		Router:        echo.New(),
		userValidator: userValidator,
		userService:   userService,
	}
}

func (s Server) Serve() {
	// Middleware

	s.Router.Use(middleware.RequestID())

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

	//s.Router.Use(middleware.Logger())

	s.Router.Use(middleware.Recover())

	// Routes
	s.Router.GET("/health-check", s.healthCheck)

	s.userService.SetRoutes(s.Router)
	//s.backofficeUserHandler.SetRoutes(s.Router)
	//s.matchingHandler.SetRoutes(s.Router)

	// Start server
	address := fmt.Sprintf(":%d", s.config.Server.HttpPort)
	fmt.Printf("start echo server on %s\n", address)
	if err := s.Router.Start(address); err != nil {
		fmt.Println("router start error", err)
	}
}
