package bybit_history_service

import (
	"github.com/bxcodec/go-clean-arch/db/postgres"
	"github.com/bxcodec/go-clean-arch/internal/bybit_history_service/models"
)

type UserRepository interface {
	FindByApiKey(apiKey string) (models.ByBitUser, error)
	FindByUsername(username string) (models.ByBitUser, error)
	FindByEmail(email string) (models.ByBitUser, error)
	FindByPhoneNumber(phoneNumber string) (models.ByBitUser, error)
}
type UserRepositoryImpl struct {
	db *db.PostgresDB
}

func NewUser(db *db.PostgresDB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (s *UserRepositoryImpl) FindByApiKey(apiKey string) (models.ByBitUser, error) {
	var user models.ByBitUser
	tx := s.db.Conn().Where("api_key = ?", apiKey).First(&user)
	if tx.Error != nil {
		return models.ByBitUser{}, tx.Error
	}
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
