package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go-clean-arch/config"
	"go-clean-arch/internal/bybit_ws"
	"go-clean-arch/internal/delivery"
	"log"
	"os"
	"os/signal"
)

func init() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

}

func main() {
	cfg := config.LoadConfig()
	bybitWsService := setupServices(cfg)

	server := delivery.NewWSServer(cfg, bybitWsService)
	go func() {
		server.Serve()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx := context.Background()
	ctxWithTimeout, cancel := context.WithTimeout(ctx, cfg.WsOrderBookServer.GracefulShutdownTimeout)
	defer cancel()

	if err := server.Router.Shutdown(ctxWithTimeout); err != nil {
		fmt.Println("http server shutdown error", err)
	}
	fmt.Println("received interrupt signal, shutting down gracefully..")
	<-ctxWithTimeout.Done()
}
func setupServices(cfg config.Config) user_service.ByBitWSService {
	//postgresDB := db.NewPostgres(cfg)
	//userRepository := account_repository.New(postgresDB)

	userService := user_service.NewByBitWSService(cfg)
	//userValidator := user_validator.New(userRepository)

	return userService
}
