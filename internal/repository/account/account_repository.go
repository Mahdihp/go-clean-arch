package account

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/bxcodec/go-clean-arch/domain/ent"
	"github.com/bxcodec/go-clean-arch/domain/ent/account"
	db "github.com/bxcodec/go-clean-arch/internal/repository/postgres"
	"github.com/bxcodec/go-clean-arch/pkg/errmsg"
	"github.com/bxcodec/go-clean-arch/pkg/richerror"
)

type AccountRepository interface {
	//Insert(ctx context.Context, req dto.UserInsertRequest) (dto.UserInfo, error)
	//Update(ctx context.Context, req dto.UserUpdateRequest) error
	//DeleteById(ctx context.Context, userId int64) error
	FindByUsernameAndPassword(ctx context.Context, username string, password string) (*ent.Account, error)
	//IsMobileUnique(ctx context.Context, mobile string) (bool, error)
	//GetById(ctx context.Context, userId int64) (dto.UserInfo, error)
	//GetAll(ctx context.Context, req dto.GetAllUserRequest) ([]dto.UserInfo, error)
}

type AccountRepositoryImpl struct {
	db *db.PostgresDB
}

func New(db *db.PostgresDB) *AccountRepositoryImpl {
	return &AccountRepositoryImpl{
		db: db,
	}
}

func (uri AccountRepositoryImpl) FindByUsernameAndPassword(ctx context.Context, username string, password string) (*ent.Account, error) {
	const op = "AccountRepositoryImpl.FindByUsernameAndPassword"

	found, err := uri.db.Conn().Account.Query().
		Where(sql.FieldEQ(account.FieldUsername, username), sql.FieldEQ(account.FieldPassword, password)).
		First(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, richerror.New(op).WithErr(err).
				WithMessage(errmsg.ErrorMsgNotFound).WithKind(richerror.KindNotFound)
		}
		return nil, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}

	return found, nil
}
