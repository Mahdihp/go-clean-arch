package mongodb

import (
	"context"
	"fmt"
	"github.com/bxcodec/go-clean-arch/config"
	models_grpc "github.com/bxcodec/go-clean-arch/internal/bybit_grpc_service/models"
	"github.com/bxcodec/go-clean-arch/internal/bybit_history_service/models"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"time"
)

type MongoDb struct {
	mongoDb *mongo.Database
}

func NewMongoDb(cfg config.MongoDb) *MongoDb {
	//ConnetionString := "mongodb://username:password@host:port"
	ConnetionString := fmt.Sprintf("mongodb://%s:%s@%s:%d", cfg.Username, cfg.Password, cfg.Host, cfg.HttpPort)
	//ConnetionString := fmt.Sprintf("mongodb://%s:%d/?directConnection=false", cfg.Host, cfg.HttpPort)
	fmt.Println(ConnetionString)

	clientOptions := options.Client().ApplyURI(ConnetionString)
	client, err := mongo.Connect(clientOptions)

	if err != nil {
		log.Info("Connect Error ", err)
	}

	// Check the connection
	if err = client.Ping(context.Background(), nil); err != nil {
		log.Info("Ping Error ", err)
	}
	var database = existDb(client, cfg)

	//seedData(database)

	log.Info("mongoClient connected")
	return &MongoDb{mongoDb: database}
}

func seedData(database *mongo.Database) {
	database.Collection(models_grpc.Coll_ByBitMarketGetInstrumentsInfoLinear).
		InsertOne(context.TODO(), models_grpc.ByBitMarketGetInstrumentsInfoLinear{

			Symbol:    "BTC-2",
			Status:    "Status-2",
			CreatedAt: time.Now(),
		})

}
func existDb(client *mongo.Client, cfg config.MongoDb) *mongo.Database {
	var d *mongo.Database

	filter := bson.D{{"name", bson.D{{"$eq", cfg.DBName}}}}
	names, err := client.ListDatabaseNames(context.Background(), filter)

	if err != nil {
		log.Info("ListDatabaseNames Error ", err)
	}
	if len(names) == 0 {
		d = client.Database(cfg.DBName)
		filter := bson.D{{}}
		collections, err := d.ListCollections(context.Background(), filter)
		if err != nil {
			log.Info("ListCollections Error ", err)
		}
		var colls []string
		if err = collections.All(context.TODO(), &colls); err != nil {
			log.Info("ListCollections All Error ", err)
		}
		if len(colls) == 0 {
			d.CreateCollection(context.Background(), models.Coll_ByBitUser)
			d.CreateCollection(context.Background(), models.Coll_BybitFutureOrderHistory)
			d.CreateCollection(context.Background(), models.Coll_BybitFutureTradeHistory)
			d.CreateCollection(context.Background(), models.Coll_BybitFuturePnlHistory)
			d.CreateCollection(context.Background(), models.Coll_BybitSpotOrderHistory)
			d.CreateCollection(context.Background(), models.Coll_BybitSpotTradelHistory)

			d.CreateCollection(context.Background(), models_grpc.Coll_ByBitMarketGetInstrumentsInfoLinear)
			d.CreateCollection(context.Background(), models_grpc.Coll_ByBitMarketGetInstrumentsInfoInverse)
			d.CreateCollection(context.Background(), models_grpc.Coll_ByBitMarketGetInstrumentsInfoSpot)
		}
	}
	if d == nil {
		d = client.Database(cfg.DBName)
	}
	return d
}
func (m *MongoDb) MongoConn() *mongo.Database {
	return m.mongoDb
}
