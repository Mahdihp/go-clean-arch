package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	WsServer             BybitWSServer             `mapstructure:"server"`
	TradeGrpcServer      BybitTradGrpcServer       `mapstructure:"server"`
	HistoricalGrpcServer BybitHistoricalGrpcServer `mapstructure:"server"`
	Postgres             Postgres                  `mapstructure:"postgres"`
	ByBit                ByBitWS
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
type ScyllaDBConfig struct {
	User     string `mapstructure:"user"`
	Pass     string `mapstructure:"pass"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Keyspace string `mapstructure:"keyspace"`
}

type BybitWSServer struct {
	Host                    string        `mapstructure:"BYBIT_WS_HOST"`
	HttpPort                int           `mapstructure:"BYBIT_WS_HOST"`
	GracefulShutdownTimeout time.Duration `mapstructure:"BYBIT_WS_GRACEFUL_SHUTDOWN_TIMEOUT"`
}

type BybitHistoricalGrpcServer struct {
	Host                    string        `mapstructure:"BYBIT_GRPC_HOST"`
	HttpPort                int           `mapstructure:"BYBIT_GRPC_PORT"`
	GracefulShutdownTimeout time.Duration `mapstructure:"BYBIT_GRPC_GRACEFUL_SHUTDOWN_TIMEOUT"`
}
type BybitTradGrpcServer struct {
	Host                    string        `mapstructure:"BYBIT_HTTP_HOST"`
	HttpPort                int           `mapstructure:"BYBIT_HTTP_PORT"`
	GracefulShutdownTimeout time.Duration `mapstructure:"BYBIT_HTTP_GRACEFUL_SHUTDOWN_TIMEOUT"`
}

func LoadConfig() Config {

	var c Config
	c.WsServer.Host = os.Getenv("BYBIT_WS_HOST")
	c.WsServer.HttpPort, _ = strconv.Atoi(os.Getenv("BYBIT_WS_HTTP_PORT"))
	c.WsServer.GracefulShutdownTimeout, _ = time.ParseDuration(os.Getenv("BYBIT_WS_GRACEFUL_SHUTDOWN_TIMEOUT"))

	c.TradeGrpcServer.Host = os.Getenv("BYBIT_TRADE_GRPC_HOST")
	c.TradeGrpcServer.HttpPort, _ = strconv.Atoi(os.Getenv("BYBIT_TRADE_GRPC_PORT"))
	c.TradeGrpcServer.GracefulShutdownTimeout, _ = time.ParseDuration(os.Getenv("BYBIT_TRADE_GRPC_GRACEFUL_SHUTDOWN_TIMEOUT"))

	c.HistoricalGrpcServer.Host = os.Getenv("BYBIT_HIS_GRPC_HOST")
	c.HistoricalGrpcServer.HttpPort, _ = strconv.Atoi(os.Getenv("BYBIT_HIS_GRPC_PORT"))
	c.HistoricalGrpcServer.GracefulShutdownTimeout, _ = time.ParseDuration(os.Getenv("BYBIT_HIS_GRPC_GRACEFUL_SHUTDOWN_TIMEOUT"))

	c.Postgres.PostgresHost = os.Getenv("POSTGRES_HOST")
	c.Postgres.PostgresDB = os.Getenv("POSTGRES_DB")
	c.Postgres.PostgresUser = os.Getenv("POSTGRES_USER")
	c.Postgres.PostgresPassword = os.Getenv("POSTGRES_PASSWORD")
	c.Postgres.PostgresPort, _ = strconv.Atoi(os.Getenv("POSTGRES_PORT"))

	c.ByBit.ApiKey = os.Getenv("APIKEY")
	c.ByBit.ApiSecret = os.Getenv("APISECRET")
	c.ByBit.WsSocketSpot = os.Getenv("WS_SOCKET_SPOT")
	c.ByBit.WsSocketLinear = os.Getenv("WS_SOCKET_LINEAR")

	return c

}
