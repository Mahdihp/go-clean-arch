package main

import (
	"github.com/bxcodec/go-clean-arch/config"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

func init() {

}
func main() {
	cfg := config.LoadConfig()

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(cfg.HistoricalGrpcServer.HttpPort))
	if err != nil {
		log.Fatalf("failed to listen on port %d: %v", cfg.HistoricalGrpcServer.HttpPort, err)
	}

	s := grpc.NewServer()
	//grpcServer_Order := bybit_historical_grpc_service.NewByBitHisGrpcServer(cfg)
	//grpcServer_Position := bybit_grpc_service.NewByBitHttpServerPosition(cfg)
	//
	//position.RegisterPositionServiceServer(s, &grpcServer_Order)
	//order.RegisterOrderServiceServer(s, &grpcServer_Order)
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
