package user

import (
	"context"

	userDomain "github.com/o-ga09/tutorial-ec-backend/app/domain/user"
)

type DeleteUserUsecase struct {
	userRepo userDomain.UserRepository
}

func NewDeleteUserUsecase(userRepo userDomain.UserRepository) *DeleteUserUsecase {
	return &DeleteUserUsecase{
		userRepo: userRepo,
	}
}

func(u DeleteUserUsecase) Run(ctx context.Context, id string) error {
	err := u.userRepo.Delete(ctx,id)
	if err != nil {
		return err
	}

	return nil
}