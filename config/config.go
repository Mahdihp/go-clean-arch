package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	Server   ServerConfig `mapstructure:"server"`
	Postgres Postgres     `mapstructure:"postgres"`
	ByBit    ByBit        `mapstructure:"postgres"`
}
type ByBit struct {
	ApiKey    string `mapstructure:"APIKEY"`
	ApiSecret string `mapstructure:"APISECRET"`
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

type ServerConfig struct {
	Host                    string        `mapstructure:"HOST"`
	HttpPort                int           `mapstructure:"HTTP_PORT"`
	GracefulShutdownTimeout time.Duration `mapstructure:"GRACEFUL_SHUTDOWN_TIMEOUT"`
}

func LoadConfig() Config {

	var c Config
	c.Server.Host = os.Getenv("HOST")
	c.Server.HttpPort, _ = strconv.Atoi(os.Getenv("HTTP_PORT"))

	c.Postgres.PostgresHost = os.Getenv("POSTGRES_HOST")
	c.Postgres.PostgresDB = os.Getenv("POSTGRES_DB")
	c.Postgres.PostgresUser = os.Getenv("POSTGRES_USER")
	c.Postgres.PostgresPassword = os.Getenv("POSTGRES_PASSWORD")
	c.Postgres.PostgresPort, _ = strconv.Atoi(os.Getenv("POSTGRES_PORT"))

	c.ByBit.ApiKey = os.Getenv("APIKEY")
	c.ByBit.ApiSecret = os.Getenv("APISECRET")

	return c

}
