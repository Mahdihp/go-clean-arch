package bybit_history_service

import (
	"context"
	"errors"
	"github.com/bxcodec/go-clean-arch/db/mongodb"
	"github.com/bxcodec/go-clean-arch/internal/bybit_history_service/models"
	"github.com/bxcodec/go-clean-arch/params"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
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
	filter := bson.D{{params.Field_Search_ApiKey, bson.D{{params.Equal_Opt, apiKey}}}}

	var user models.ByBitUser
	err := collection.FindOne(ctx, filter).Decode(&user)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return models.ByBitUser{}, err
	}
	return user, nil
}
