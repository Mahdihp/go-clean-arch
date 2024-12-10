package db

import (
	"database/sql"
	"fmt"
	"github.com/bxcodec/go-clean-arch/config"
	"github.com/bxcodec/go-clean-arch/internal/bybit_history_service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

type PostgresDB struct {
	db    *gorm.DB
	SqlDB *sql.DB
	cfg   config.Postgres
}

func NewPostgres(cfg config.Postgres) *PostgresDB {
	//dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	ConnetionString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresDB, cfg.PostgresPassword)

	//db, err := gorm.Open(postgres.Open(ConnetionString), &gorm.Config{})

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: ConnetionString,

		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	// Migrate the schema
	db.AutoMigrate(&models.ByBitUser{})
	db.AutoMigrate(&models.BybitFutureOrderHistory{})
	db.AutoMigrate(&models.BybitFutureTradeHistory{})
	db.AutoMigrate(&models.BybitFuturePnlHistory{})
	db.AutoMigrate(&models.BybitSpotOrderHistory{})
	db.AutoMigrate(&models.BybitSpotTradelHistory{})

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("Error getting *sql.DB object:", err)
	}

	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	defer sqlDB.Close()
	return &PostgresDB{cfg: cfg, db: db, SqlDB: sqlDB}
}

func (m *PostgresDB) Conn() *gorm.DB {
	return m.db
}
