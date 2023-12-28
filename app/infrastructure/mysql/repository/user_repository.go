package repository

import (
	"context"

	userDomain "github.com/o-ga09/tutorial-ec-backend/app/domain/user"
	"github.com/o-ga09/tutorial-ec-backend/app/infrastructure/mysql/schema"
	"gorm.io/gorm"
)

type userRepository struct {
	conn *gorm.DB
}

// FindAll implements user.UserRepository.
func (u *userRepository) FindAll(ctx context.Context) ([]*userDomain.User, error) {
	res := []*userDomain.User{}
	users := []*schema.User{}
	u.conn.Find(users)

	for _, user := range users {
		u, err := userDomain.Reconstract(user.UserID,user.Email,user.PhoneNumber,user.LastName,user.FirstName,user.Pref,user.City,user.Extra)
		if err != nil {
			return nil, err
		}

		res = append(res, u)
	}

	return res , nil
}

// FindById implements user.UserRepository.
func (u *userRepository) FindById(ctx context.Context, id string) (*userDomain.User, error) {
	user := schema.User{}

	u.conn.Where("id = ?",id).Find(user)
	res, err := userDomain.Reconstract(user.UserID,user.Email,user.PhoneNumber,user.LastName,user.FirstName,user.Pref,user.City,user.Extra)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Save implements user.UserRepository.
func (u *userRepository) Save(ctx context.Context, user *userDomain.User) error {
	repoUser := schema.User{
		UserID: user.ID(),
		Email: user.Email(),
		PhoneNumber: user.PhoneNumber(),
		LastName: user.LastName(),
		FirstName: user.FirstName(),
		Pref: user.Pref(),
		City: user.City(),
		Extra: user.Extra(),
	}

	u.conn.Create(repoUser)
	return nil
}

func NewUserRepository(conn *gorm.DB) userDomain.UserRepository {
	return &userRepository{conn: conn}
}
