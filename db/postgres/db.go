package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/bxcodec/go-clean-arch/config"
	"github.com/bxcodec/go-clean-arch/internal/bybit_history_service/models"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

type PostgresDB struct {
	mongoDb *mongo.Database
	db      *gorm.DB
	SqlDB   *sql.DB
	config  config.Postgres
	cfg     config.MongoDb
}

func NewMongoDb(cfg config.MongoDb) *PostgresDB {
	//ConnetionString := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", cfg.Username, cfg.Password, cfg.Host, cfg.HttpPort, cfg.DBName)
	ConnetionString := fmt.Sprintf("mongodb://%s:%d/?directConnection=false", cfg.Host, cfg.HttpPort)
	fmt.Println(ConnetionString)

	clientOptions := options.Client().ApplyURI(ConnetionString)
	client, err := mongo.Connect(clientOptions)

	if err != nil {
		logrus.Info("Connect Error ", err)
	}

	// Check the connection
	if err = client.Ping(context.TODO(), nil); err != nil {
		logrus.Info("Ping Error ", err)
	}
	database := client.Database(cfg.DBName)

	database.CreateCollection(context.Background(), models.Coll_ByBitUser)
	database.CreateCollection(context.Background(), models.Coll_BybitFutureOrderHistory)
	database.CreateCollection(context.Background(), models.Coll_BybitFutureTradeHistory)
	database.CreateCollection(context.Background(), models.Coll_BybitFuturePnlHistory)
	database.CreateCollection(context.Background(), models.Coll_BybitSpotOrderHistory)
	database.CreateCollection(context.Background(), models.Coll_BybitSpotTradelHistory)

	//filter := bson.D{{}}
	//collections, _ := database.ListCollectionNames(context.Background(), filter)
	//fmt.Println(collections)

	logrus.Info("mongoClient connected")
	return &PostgresDB{cfg: cfg, mongoDb: database}
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
	return &PostgresDB{config: cfg, db: db, SqlDB: sqlDB}
}

func (m *PostgresDB) Conn() *gorm.DB {
	return m.db
}
func (m *PostgresDB) MongoConn() *mongo.Database {
	return m.mongoDb
}
