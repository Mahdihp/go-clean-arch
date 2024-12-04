package bybit_history_service

import (
	"context"
	"github.com/bxcodec/go-clean-arch/db/postgres"
	"github.com/bxcodec/go-clean-arch/internal/bybit_history_service/models/ent"
	"github.com/bxcodec/go-clean-arch/internal/bybit_history_service/models/ent/bybituser"
)

type UserRepository interface {
	FindByApiKey(ctx context.Context, apiKey string) (ent.ByBitUser, error)
	FindByUsername(ctx context.Context, username string) (ent.ByBitUser, error)
	FindByEmail(ctx context.Context, email string) (ent.ByBitUser, error)
	FindByPhoneNumber(ctx context.Context, phoneNumber string) (ent.ByBitUser, error)
}
type UserRepositoryImpl struct {
	db *db.PostgresDB
}

func NewUser(db *db.PostgresDB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}
func (s *UserRepositoryImpl) FindByApiKey(ctx context.Context, apiKey string) (ent.ByBitUser, error) {
	found, err := s.db.Conn().ByBitUser.
		//Query().Where(sql.FieldEQ(bybituser.FieldAPIKey, apiKey)).First(ctx)
		Query().Where(bybituser.APIKey(apiKey)).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return ent.ByBitUser{}, err
		}
	}
	return *found, nil
}

func (s *UserRepositoryImpl) FindByUsername(ctx context.Context, username string) (ent.ByBitUser, error) {
	found, err := s.db.Conn().ByBitUser.
		//Query().Where(sql.FieldEQ(bybituser.FieldAPIKey, apiKey)).First(ctx)
		Query().Where(bybituser.Username(username)).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return ent.ByBitUser{}, err
		}
	}
	return *found, nil
}

func (s *UserRepositoryImpl) FindByEmail(ctx context.Context, email string) (ent.ByBitUser, error) {
	found, err := s.db.Conn().ByBitUser.
		//Query().Where(sql.FieldEQ(bybituser.FieldAPIKey, apiKey)).First(ctx)
		Query().Where(bybituser.Email(email)).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return ent.ByBitUser{}, err
		}
	}
	return *found, nil
}

func (s *UserRepositoryImpl) FindByPhoneNumber(ctx context.Context, phoneNumber string) (ent.ByBitUser, error) {
	found, err := s.db.Conn().ByBitUser.
		//Query().Where(sql.FieldEQ(bybituser.FieldAPIKey, apiKey)).First(ctx)
		Query().Where(bybituser.PhoneNumber(phoneNumber)).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return ent.ByBitUser{}, err
		}
	}
	return *found, nil
}
