package repository

import (
	"context"

	userDomain "github.com/o-ga09/tutorial-ec-backend/app/domain/user"
	"gorm.io/gorm"
)

type userRepository struct {
	conn *gorm.DB
}

type User struct {
	id string
	email string
	phoneNumber string
	lastName string
	firstName string
	pref string
	city string
	extra string
}

// FindAll implements user.UserRepository.
func (u *userRepository) FindAll(ctx context.Context) ([]*userDomain.User, error) {
	res := []*userDomain.User{}
	users := []*User{}
	u.conn.Find(users)

	for _, user := range users {
		u, err := userDomain.Reconstract(user.id,user.email,user.phoneNumber,user.lastName,user.firstName,user.pref,user.city,user.extra)
		if err != nil {
			return nil, err
		}

		res = append(res, u)
	}

	return res , nil
}

// FindById implements user.UserRepository.
func (u *userRepository) FindById(ctx context.Context, id string) (*userDomain.User, error) {
	user := User{}

	u.conn.Where("id = ?",id).Find(user)
	res, err := userDomain.Reconstract(user.id,user.email,user.phoneNumber,user.lastName,user.firstName,user.pref,user.city,user.extra)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Save implements user.UserRepository.
func (u *userRepository) Save(ctx context.Context, user *userDomain.User) error {
	repoUser := User{
		id: user.ID(),
		email: user.Email(),
		phoneNumber: user.PhoneNumber(),
		lastName: user.LastName(),
		firstName: user.FirstName(),
		pref: user.Pref(),
		city: user.City(),
		extra: user.Extra(),
	}

	u.conn.Create(repoUser)
	return nil
}

func NewUserRepository(conn *gorm.DB) userDomain.UserRepository {
	return &userRepository{conn: conn}
}
