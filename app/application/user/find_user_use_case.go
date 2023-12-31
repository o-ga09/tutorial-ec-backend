package user

import (
	"context"
	"log/slog"

	userDomain "github.com/o-ga09/tutorial-ec-backend/app/domain/user"
	"github.com/o-ga09/tutorial-ec-backend/app/server/middleware"
)

type FindUserUsecase struct {
	userRepo userDomain.UserRepository
}

func NewFindUserUsecase(userRepo userDomain.UserRepository) *FindUserUsecase {
	return &FindUserUsecase{userRepo: userRepo}
}

type FindUserUsecaseDto struct {
	ID string
	Email string
	Phonenumber string
	Lastname string
	Firstname string
	Address string
}

func(u *FindUserUsecase) Run(ctx context.Context, id string) (*FindUserUsecaseDto, error) {
	user, err := u.userRepo.FindById(ctx, id)
	if err != nil {
		slog.Log(ctx, middleware.SeverityError, "error","app error",err)
		return nil, err
	}
	
	return &FindUserUsecaseDto{
		ID: user.ID(),
		Email: user.Email(),
		Phonenumber: user.PhoneNumber(),
		Lastname: user.LastName(),
		Firstname: user.FirstName(),
		Address: user.Pref() + user.City() + user.Extra(),
	}, nil
}