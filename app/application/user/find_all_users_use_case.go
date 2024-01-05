package user

import (
	"context"

	userDomain "github.com/o-ga09/tutorial-ec-backend/app/domain/user"
)

type FindAllUsersUseCase struct {
	userRepo userDomain.UserRepository
}

func NewFindAllUsersUseCase(userRepo userDomain.UserRepository) *FindAllUsersUseCase {
	return &FindAllUsersUseCase{userRepo: userRepo}
}

func(u *FindAllUsersUseCase) Run(ctx context.Context) ([]*FindUserUsecaseDto, error) {
	users, err := u.userRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	res := []*FindUserUsecaseDto{}
	for _, user := range users {
		u := FindUserUsecaseDto{
			ID: user.ID(),
			Email: user.Email(),
			Phonenumber: user.PhoneNumber(),
			Lastname: user.LastName(),
			Firstname: user.FirstName(),
			Address: user.Pref() + user.City() + user.Extra(),
		}
		res = append(res, &u)
	}

	return res, nil
}