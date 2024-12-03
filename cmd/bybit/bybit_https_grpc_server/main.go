package main

import (
	"github.com/bxcodec/go-clean-arch/adapter/grpc-proto/order"
	"github.com/bxcodec/go-clean-arch/adapter/grpc-proto/position"
	"github.com/bxcodec/go-clean-arch/config"
	"github.com/bxcodec/go-clean-arch/internal/bybit_grpc_service"
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

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(cfg.TradeGrpcServer.HttpPort))
	if err != nil {
		log.Fatalf("failed to listen on port %d: %v", cfg.TradeGrpcServer.HttpPort, err)
	}

	s := grpc.NewServer()
	grpcServer_Order := bybit_grpc_service.NewByBitHttpServerOrder(cfg)
	grpcServer_Position := bybit_grpc_service.NewByBitHttpServerPosition(cfg)

	position.RegisterPositionServiceServer(s, &grpcServer_Position)
	order.RegisterOrderServiceServer(s, &grpcServer_Order)
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
