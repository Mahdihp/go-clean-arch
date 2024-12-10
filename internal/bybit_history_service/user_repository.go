package bybit_history_service

import (
	"context"
	"fmt"
	"github.com/bxcodec/go-clean-arch/db/mongodb"
	"github.com/bxcodec/go-clean-arch/internal/bybit_history_service/models"
	"go.mongodb.org/mongo-driver/bson"
)

type UserRepository interface {
	FindByApiKey(ctx context.Context, apiKey string) (models.ByBitUser, error)
	//FindByUsername(username string) (models.ByBitUser, error)
	//FindByEmail(email string) (models.ByBitUser, error)
	//FindByPhoneNumber(phoneNumber string) (models.ByBitUser, error)
}
type UserRepositoryImpl struct {
	db *mongodb.MongoDb
}

func NewUser(db *mongodb.MongoDb) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (s *UserRepositoryImpl) FindByApiKey(ctx context.Context, apiKey string) (models.ByBitUser, error) {
	collection := s.db.MongoConn().Collection(models.Coll_ByBitUser)
	//filter := bson.D{primitive.E{Key: "api_key", Value: apiKey}}
	filter := bson.D{{"api_key", bson.D{{"$eq", apiKey}}}}

	var user models.ByBitUser
	collection.FindOne(ctx, filter).Decode(user)
	fmt.Println(user)

	//tx := s.db.Conn().Where("api_key = ?", apiKey).First(&user)
	//if tx.Error != nil {
	//	return models.ByBitUser{}, tx.Error
	//}
	return user, nil
}
