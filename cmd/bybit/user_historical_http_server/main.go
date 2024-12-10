package main

import (
	"context"
	"fmt"
	"github.com/bxcodec/go-clean-arch/config"
	"github.com/bxcodec/go-clean-arch/db/mongodb"
	"github.com/bxcodec/go-clean-arch/internal/bybit_history_service"
	"github.com/bxcodec/go-clean-arch/internal/delivery"
	"github.com/joho/godotenv"
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
	bybitHisService := setupServices(cfg)
	server := delivery.NewHistoricalServer(cfg, bybitHisService)
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

func setupServices(cfg config.Config) bybit_history_service.ByBitHistoricalServic {
	//postgresDB := db.NewPostgres(cfg.Postgres)
	mongoDb := mongodb.NewMongoDb(cfg.MongoDb)

	userRepo := bybit_history_service.NewUser(mongoDb)
	historyRepo := bybit_history_service.NewHistory(mongoDb)

	service := bybit_history_service.NewByBitHistoricalServic(cfg, userRepo, historyRepo)

	return service
}
