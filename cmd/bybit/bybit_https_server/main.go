package main

import (
	"fmt"
	"github.com/bxcodec/go-clean-arch/adapter/grpc-proto/order"
	"github.com/bxcodec/go-clean-arch/config"
	"github.com/bxcodec/go-clean-arch/internal/bybit_http"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
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

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(cfg.HttpServer.HttpPort))
	if err != nil {
		log.Fatalf("failed to listen on port %d: %v", cfg.HttpServer.HttpPort, err)
	}

	s := grpc.NewServer()
	httpServer := bybit_http.NewByBitHttpService(cfg)
	order.RegisterOrderServiceServer(s, &httpServer)
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
