package main

import (
	"github.com/bxcodec/go-clean-arch/adapter/grpc-proto/account"
	"github.com/bxcodec/go-clean-arch/adapter/grpc-proto/market"
	"github.com/bxcodec/go-clean-arch/adapter/grpc-proto/order"
	"github.com/bxcodec/go-clean-arch/adapter/grpc-proto/position"
	"github.com/bxcodec/go-clean-arch/config"
	"github.com/bxcodec/go-clean-arch/internal/bybit_grpc_service"
	"github.com/bxcodec/go-clean-arch/internal/bybit_grpc_service/jobs"
	"github.com/bxcodec/go-clean-arch/internal/bybit_grpc_service/repository"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

func init() {
	err := godotenv.Load("env")
	//err := godotenv.Load(filepath.Join("./config", "env"))

	if err != nil {
		log.Fatalf("Error loading env file: %v", err)
	}
}

func main() {
	cfg := config.LoadConfig()

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(cfg.TradeGrpcServer.HttpPort))
	if err != nil {
		log.Fatalf("failed to listen on port %d: %v", cfg.TradeGrpcServer.HttpPort, err)
	}

	marketRepository := repository.New(cfg)

	s := grpc.NewServer()

	grpcServer_Order := bybit_grpc_service.NewByBitHttpServerOrder(cfg)
	grpcServer_Position := bybit_grpc_service.NewByBitHttpServerPosition(cfg)
	grpcServer_Market := bybit_grpc_service.NewByBitHttpServerMarket(cfg, marketRepository)
	grpcserver_Account := bybit_grpc_service.NewByBitHttpServerAccount(cfg)

	position.RegisterPositionServiceServer(s, &grpcServer_Position)
	order.RegisterOrderServiceServer(s, &grpcServer_Order)
	market.RegisterMarketServiceServer(s, &grpcServer_Market)
	account.RegisterAccountServicServer(s, &grpcserver_Account)

	go jobs.SetupCronJob(cfg, marketRepository)

	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
