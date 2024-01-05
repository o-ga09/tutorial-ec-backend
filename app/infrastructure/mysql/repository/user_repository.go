package repository

import (
	"context"
	"errors"
	"log/slog"

	userDomain "github.com/o-ga09/tutorial-ec-backend/app/domain/user"
	"github.com/o-ga09/tutorial-ec-backend/app/infrastructure/mysql/schema"
	"github.com/o-ga09/tutorial-ec-backend/app/server/middleware"
	"gorm.io/gorm"
)

type userRepository struct {
	conn *gorm.DB
}

// Delete implements user.UserRepository.
func (u *userRepository) Delete(ctx context.Context, id string) error {
	user := schema.User{}
	err := u.conn.Where("user_id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}

	return nil
}

// FindAll implements user.UserRepository.
func (u *userRepository) FindAll(ctx context.Context) ([]*userDomain.User, error) {
	res := []*userDomain.User{}
	users := []*schema.User{}
	err := u.conn.Find(&users).Error
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		u, err := userDomain.Reconstract(user.UserID, user.Email, user.PhoneNumber, user.LastName, user.FirstName, user.Pref, user.City, user.Extra)
		if err != nil {
			return nil, err
		}

		res = append(res, u)
	}

	return res, nil
}

// FindById implements user.UserRepository.
func (u *userRepository) FindById(ctx context.Context, id string) (*userDomain.User, error) {
	user := schema.User{}

	err := u.conn.Where("user_id = ?", id).Find(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		slog.Log(ctx, middleware.SeverityError, "record not found", "err msg", id)
		return nil, gorm.ErrRecordNotFound
	}

	slog.Log(ctx, middleware.SeverityInfo, "Debug", "user", user)
	res, err := userDomain.Reconstract(user.UserID, user.Email, user.PhoneNumber, user.LastName, user.FirstName, user.Pref, user.City, user.Extra)
	if err != nil {
		slog.Log(ctx, middleware.SeverityError, "repository error", "err msg", err)
		return nil, err
	}
	return res, nil
}

// Save implements user.UserRepository.
func (u *userRepository) Save(ctx context.Context, user *userDomain.User) error {
	repoUser := schema.User{
		UserID:      user.ID(),
		Email:       user.Email(),
		PhoneNumber: user.PhoneNumber(),
		LastName:    user.LastName(),
		FirstName:   user.FirstName(),
		Pref:        user.Pref(),
		City:        user.City(),
		Extra:       user.Extra(),
	}

	err := u.conn.Save(&repoUser).Error
	return err
}

func NewUserRepository(conn *gorm.DB) userDomain.UserRepository {
	return &userRepository{conn: conn}
}
