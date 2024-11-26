package main

import (
	"context"
	"fmt"
	user_service "github.com/bxcodec/go-clean-arch/account"
	"github.com/bxcodec/go-clean-arch/config"
	"github.com/bxcodec/go-clean-arch/internal/delivery"
	account_repository "github.com/bxcodec/go-clean-arch/internal/repository/account"
	db "github.com/bxcodec/go-clean-arch/internal/repository/postgres"
	user_validator "github.com/bxcodec/go-clean-arch/internal/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
)

func init() {
	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

}

func main() {
	cfg := config.LoadConfig()
	fmt.Println(cfg)
	userService, userValidator := setupServices(cfg)

	server := delivery.NewServer(cfg, userValidator, userService)
	go func() {
		server.Serve()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx := context.Background()
	ctxWithTimeout, cancel := context.WithTimeout(ctx, cfg.Server.GracefulShutdownTimeout)
	defer cancel()

	if err := server.Router.Shutdown(ctxWithTimeout); err != nil {
		fmt.Println("http server shutdown error", err)
	}
	fmt.Println("received interrupt signal, shutting down gracefully..")
	<-ctxWithTimeout.Done()
}
func setupServices(cfg config.Config) (user_service.UserService, user_validator.UserValidator) {
	postgresDB := db.NewPostgres(cfg)
	userRepository := account_repository.New(postgresDB)

	userService := user_service.NewUserService(cfg, userRepository)
	userValidator := user_validator.New(userRepository)

	return userService, userValidator
}
