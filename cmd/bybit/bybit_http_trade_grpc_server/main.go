package main

import (
	"context"
	"fmt"
	"github.com/bxcodec/go-clean-arch/adapter/grpc-proto/market"
	"github.com/bxcodec/go-clean-arch/adapter/grpc-proto/order"
	"github.com/bxcodec/go-clean-arch/adapter/grpc-proto/position"
	"github.com/bxcodec/go-clean-arch/config"
	"github.com/bxcodec/go-clean-arch/internal/bybit_grpc_service"
	"github.com/bxcodec/go-clean-arch/internal/bybit_grpc_service/jobs"
	models_grpc "github.com/bxcodec/go-clean-arch/internal/bybit_grpc_service/models"
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

	marketRepository := bybit_grpc_service.New(cfg)

	s := grpc.NewServer()
	grpcServer_Order := bybit_grpc_service.NewByBitHttpServerOrder(cfg)
	grpcServer_Position := bybit_grpc_service.NewByBitHttpServerPosition(cfg)
	grpcServer_Market := bybit_grpc_service.NewByBitHttpServerMarket(cfg, marketRepository)

	position.RegisterPositionServiceServer(s, &grpcServer_Position)
	order.RegisterOrderServiceServer(s, &grpcServer_Order)
	market.RegisterMarketServiceServer(s, &grpcServer_Market)

	go setupCronJob(cfg, marketRepository)

	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// Start Cron Jobs for Update Linear,Spot,Inverse Mongo Collection Db
func setupCronJob(cfg config.Config, svc bybit_grpc_service.ByBitMarketRepository) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("setupCronJob. Error:\n", r)
		}
	}()

	s, err := gocron.NewScheduler()
	if err != nil {
		log.Fatalf("failed to gocron start: %v", err)
	}

	duration := time.Duration(cfg.CronJob.DurationBySecond) * time.Second
	if svc.GetCountCollecton(context.Background(), models_grpc.Collection_ByBit_MGIIL) <= 0 {
		_, _ = s.NewJob(gocron.OneTimeJob(gocron.OneTimeJobStartImmediately()),
			jobs.UpdateInstrumentInfoLinear(cfg, svc),
		)
	} else {
		s.NewJob(gocron.DurationJob(duration), jobs.UpdateInstrumentInfoLinear(cfg, svc))
	}
	duration = time.Duration(cfg.CronJob.DurationBySecond+300) * time.Second
	if svc.GetCountCollecton(context.Background(), models_grpc.Collection_ByBit_MGIIS) <= 0 {
		_, _ = s.NewJob(gocron.OneTimeJob(gocron.OneTimeJobStartImmediately()),
			jobs.UpdateInstrumentInfoSpot(cfg, svc))
	} else {
		s.NewJob(gocron.DurationJob(duration), jobs.UpdateInstrumentInfoSpot(cfg, svc))
	}

	duration = time.Duration(cfg.CronJob.DurationBySecond+600) * time.Second
	if svc.GetCountCollecton(context.Background(), models_grpc.Collection_ByBit_MGIII) <= 0 {
		_, _ = s.NewJob(gocron.OneTimeJob(gocron.OneTimeJobStartImmediately()),
			jobs.UpdateInstrumentInfoInverse(cfg, svc))
	} else {
		s.NewJob(gocron.DurationJob(duration), jobs.UpdateInstrumentInfoInverse(cfg, svc))
	}

	s.Start()
	log.Printf("Cron Job Is Started...")
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
