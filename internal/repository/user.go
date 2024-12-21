package repository

import (
	"context"

	"github.com/xianlianghe0123/go-web-mvc-tpl/internal/model/do"
	"github.com/xianlianghe0123/go-web-mvc-tpl/internal/repository/po"
	"github.com/xianlianghe0123/go-web-mvc-tpl/internal/utils/helper"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *do.UserDO) error
	Update(ctx context.Context, user *do.UserDO) error
	Get(ctx context.Context, id int64) (*do.UserDO, error)
}

func NewUserRepository(
	r *Repository,
) UserRepository {
	return &userRepository{
		Repository: r,
	}
}

type userRepository struct {
	*Repository
}

func (repo *userRepository) Create(ctx context.Context, userDO *do.UserDO) error {
	userPO := repo.po(userDO)
	err := repo.DB(ctx).User.WithContext(ctx).Create(userPO)
	if err != nil {
		return err
	}
	userDO.Id = userPO.ID
	return nil
}

func (repo *userRepository) Update(ctx context.Context, userDO *do.UserDO) error {
	userPO := repo.po(userDO)
	err := repo.DB(ctx).User.WithContext(ctx).Save(userPO)
	if err != nil {
		return err
	}
	return nil
}

func (repo *userRepository) Get(ctx context.Context, id int64) (*do.UserDO, error) {
	userPO, err := repo.DB(ctx).User.WithContext(ctx).
		Where(repo.db.User.ID.Eq(id)).
		Take()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, helper.UserNotFound.WithCause(err)
		}
		return nil, err
	}
	return repo.do(userPO), nil
}

func (repo *userRepository) do(po *po.User) *do.UserDO {
	return &do.UserDO{
		Id:       po.ID,
		Username: po.Username,
		Nickname: po.Nickname,
		Password: po.Password,
	}
}

func (repo *userRepository) po(do *do.UserDO) *po.User {
	return &po.User{
		ID:       do.Id,
		Username: do.Username,
		Nickname: do.Nickname,
		Password: do.Password,
	}
}
