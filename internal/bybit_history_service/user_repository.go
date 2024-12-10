package bybit_history_service

import (
	"context"
	"fmt"
	"github.com/bxcodec/go-clean-arch/db/postgres"
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
	db *db.PostgresDB
}

func NewUser(db *db.PostgresDB) *UserRepositoryImpl {
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

func (s *UserRepositoryImpl) FindByUsername(username string) (models.ByBitUser, error) {
	var user models.ByBitUser
	tx := s.db.Conn().Where("username = ?", username).First(&user)
	if tx.Error != nil {
		return models.ByBitUser{}, tx.Error
	}
	return user, nil
}

func (s *UserRepositoryImpl) FindByEmail(email string) (models.ByBitUser, error) {
	var user models.ByBitUser
	tx := s.db.Conn().Where("email = ?", email).First(&user)
	if tx.Error != nil {
		return models.ByBitUser{}, tx.Error
	}
	return user, nil
}

func (s *UserRepositoryImpl) FindByPhoneNumber(phoneNumber string) (models.ByBitUser, error) {
	var user models.ByBitUser
	tx := s.db.Conn().Where("phoneNumber = ?", phoneNumber).First(&user)
	if tx.Error != nil {
		return models.ByBitUser{}, tx.Error
	}
	return user, nil
}
