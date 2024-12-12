package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	WsOrderBookServer BybitWSServer
	TradeGrpcServer   BybitTradGrpcServer
	HistoricalServer  BybitHistoricalServer
	Postgres          Postgres
	ByBitWs           ByBitWS
	MongoDbMarket     MongoDbMarket
	MongoDbHistory    MongoDbHistory
	CronJob           CronJob
	MaxCountMarket    int
}
type ByBitWS struct {
	ApiKey         string `mapstructure:"APIKEY"`
	ApiSecret      string `mapstructure:"APISECRET"`
	WsSocketSpot   string `mapstructure:"WS_SOCKET_SPOT"`
	WsSocketLinear string `mapstructure:"WS_SOCKET_LINEAR"`
}

type Postgres struct {
	PostgresHost     string `mapstructure:"POSTGRES_HOST"`
	PostgresDB       string `mapstructure:"POSTGRES_DB"`
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresPort     int    `mapstructure:"POSTGRES_PORT"`
}

type BybitWSServer struct {
	Host                    string        `mapstructure:"BYBIT_WS_HOST"`
	HttpPort                int           `mapstructure:"BYBIT_WS_HOST"`
	GracefulShutdownTimeout time.Duration `mapstructure:"BYBIT_WS_GRACEFUL_SHUTDOWN_TIMEOUT"`
}

type BybitHistoricalServer struct {
	Host                    string        `mapstructure:"BYBIT_GRPC_HOST"`
	HttpPort                int           `mapstructure:"BYBIT_GRPC_PORT"`
	GracefulShutdownTimeout time.Duration `mapstructure:"BYBIT_GRPC_GRACEFUL_SHUTDOWN_TIMEOUT"`
}
type BybitTradGrpcServer struct {
	Host                    string        `mapstructure:"BYBIT_HTTP_HOST"`
	HttpPort                int           `mapstructure:"BYBIT_HTTP_PORT"`
	GracefulShutdownTimeout time.Duration `mapstructure:"BYBIT_HTTP_GRACEFUL_SHUTDOWN_TIMEOUT"`
}
type MongoDbMarket struct {
	Host     string `mapstructure:"MONGODB_MARKET_HOST"`
	HttpPort int    `mapstructure:"MONGODB_MARKET_PORT"`
	DBName   string `mapstructure:"MONGODB_MARKET_DB"`
	Username string `mapstructure:"MONGODB_MARKET_USER"`
	Password string `mapstructure:"MONGODB_MARKET_PASS"`
}

type MongoDbHistory struct {
	Host     string `mapstructure:"MONGODB_HISTORY_HOST"`
	HttpPort int    `mapstructure:"MONGODB_HISTORY_PORT"`
	DBName   string `mapstructure:"MONGODB_HISTORY_DB"`
	Username string `mapstructure:"MONGODB_HISTORY_USER"`
	Password string `mapstructure:"MONGODB_HISTORY_PASS"`
}
type CronJob struct {
	DurationBySecond int `mapstructure:"GOCRON_DURATION_JOB_SECOND"`
}

func LoadConfig() Config {

	var c Config
	c.WsOrderBookServer.Host = os.Getenv("BYBIT_WS_HOST")
	c.WsOrderBookServer.HttpPort, _ = strconv.Atoi(os.Getenv("BYBIT_WS_HTTP_PORT"))
	c.WsOrderBookServer.GracefulShutdownTimeout, _ = time.ParseDuration(os.Getenv("BYBIT_WS_GRACEFUL_SHUTDOWN_TIMEOUT"))

	c.TradeGrpcServer.Host = os.Getenv("BYBIT_TRADE_GRPC_HOST")
	c.TradeGrpcServer.HttpPort, _ = strconv.Atoi(os.Getenv("BYBIT_TRADE_GRPC_PORT"))
	c.TradeGrpcServer.GracefulShutdownTimeout, _ = time.ParseDuration(os.Getenv("BYBIT_TRADE_GRPC_GRACEFUL_SHUTDOWN_TIMEOUT"))

	c.HistoricalServer.Host = os.Getenv("BYBIT_HIS_HOST")
	c.HistoricalServer.HttpPort, _ = strconv.Atoi(os.Getenv("BYBIT_HIS_PORT"))
	c.HistoricalServer.GracefulShutdownTimeout, _ = time.ParseDuration(os.Getenv("BYBIT_HIS_GRACEFUL_SHUTDOWN_TIMEOUT"))

	c.Postgres.PostgresHost = os.Getenv("POSTGRES_HOST")
	c.Postgres.PostgresDB = os.Getenv("POSTGRES_DB")
	c.Postgres.PostgresUser = os.Getenv("POSTGRES_USER")
	c.Postgres.PostgresPassword = os.Getenv("POSTGRES_PASSWORD")
	c.Postgres.PostgresPort, _ = strconv.Atoi(os.Getenv("POSTGRES_PORT"))

	c.ByBitWs.ApiKey = os.Getenv("APIKEY")
	c.ByBitWs.ApiSecret = os.Getenv("APISECRET")
	c.ByBitWs.WsSocketSpot = os.Getenv("WS_SOCKET_SPOT")
	c.ByBitWs.WsSocketLinear = os.Getenv("WS_SOCKET_LINEAR")

	c.MongoDbMarket.Host = os.Getenv("MONGODB_MARKET_HOST")
	c.MongoDbMarket.HttpPort, _ = strconv.Atoi(os.Getenv("MONGODB_MARKET_PORT"))
	c.MongoDbMarket.DBName = os.Getenv("MONGODB_MARKET_DB")
	c.MongoDbMarket.Username = os.Getenv("MONGODB_MARKET_USER")
	c.MongoDbMarket.Password = os.Getenv("MONGODB_MARKET_PASS")

	c.MongoDbHistory.Host = os.Getenv("MONGODB_HISTORY_HOST")
	c.MongoDbHistory.HttpPort, _ = strconv.Atoi(os.Getenv("MONGODB_HISTORY_PORT"))
	c.MongoDbHistory.DBName = os.Getenv("MONGODB_HISTORY_MARKET_DB")
	c.MongoDbHistory.Username = os.Getenv("MONGODB_HISTORY_USER")
	c.MongoDbHistory.Password = os.Getenv("MONGODB_HISTORY_PASS")

	c.CronJob.DurationBySecond, _ = strconv.Atoi(os.Getenv("GOCRON_DURATION_JOB_SECOND"))
	c.MaxCountMarket, _ = strconv.Atoi(os.Getenv("MAX_COUNT_MARKET"))

	return c

}
