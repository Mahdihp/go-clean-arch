package main

import (
	"fmt"
	"github.com/bxcodec/go-clean-arch/adapter/grpc-proto/market"
	"github.com/bxcodec/go-clean-arch/adapter/grpc-proto/order"
	"github.com/bxcodec/go-clean-arch/adapter/grpc-proto/position"
	"github.com/bxcodec/go-clean-arch/config"
	"github.com/bxcodec/go-clean-arch/db/mongodb"
	"github.com/bxcodec/go-clean-arch/internal/bybit_grpc_service"
	"github.com/bxcodec/go-clean-arch/internal/bybit_grpc_service/jobs"
	"github.com/go-co-op/gocron/v2"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
	"time"
)

func init() {
	err := godotenv.Load(".env")
	//err := godotenv.Load(filepath.Join("./config", ".env"))

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

	db := mongodb.NewMongoDb(cfg.MongoDb)

	s := grpc.NewServer()
	grpcServer_Order := bybit_grpc_service.NewByBitHttpServerOrder(cfg)
	grpcServer_Position := bybit_grpc_service.NewByBitHttpServerPosition(cfg)
	grpcServer_Market := bybit_grpc_service.NewByBitHttpServerMarket(cfg, db)

	position.RegisterPositionServiceServer(s, &grpcServer_Position)
	order.RegisterOrderServiceServer(s, &grpcServer_Order)
	market.RegisterMarketServiceServer(s, &grpcServer_Market)

	go setupCronJob(cfg, db)

	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func setupCronJob(cfg config.Config, db *mongodb.MongoDb) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("setupCronJob. Error:\n", r)
		}
	}()

	s, err := gocron.NewScheduler()
	if err != nil {
		log.Fatalf("failed to gocron start: %v", err)
	}
	jobs.UpdateInstrumentInfoSpot(cfg, db)
	//duration := time.Duration(cfg.CronJob.DurationBySecond) * time.Second

	//s.NewJob(gocron.DurationJob(duration), jobs.UpdateInstrumentInfoLinear(cfg, db))

	s.Start()

	// block until you are ready to shut down
	select {
	case <-time.After(time.Minute):
	}

	// when you're done, shut it down
	err = s.Shutdown()
	if err != nil {
		fmt.Println("Shutdown setupCronJob....")
	}
}
