package user

import (
	"context"

	userDomain "github.com/o-ga09/tutorial-ec-backend/app/domain/user"
)

type SaveUserUsecase struct {
	userRepo userDomain.UserRepository
}

func NewSaveUserUsecase(userRepo userDomain.UserRepository) *SaveUserUsecase {
	return &SaveUserUsecase{userRepo: userRepo}
}

type SaveUserUsecaseDto struct {
	Email string
	Phonenumber string
	Lastname string
	Firstname string
	Pref string
	City string
	Extra string
}

func(u *SaveUserUsecase) Run(ctx context.Context, dto SaveUserUsecaseDto) error {
	// userからdtoへ変換
	user, err := userDomain.NewUser(dto.Email,dto.Phonenumber,dto.Lastname,dto.Firstname,dto.Pref,dto.City,dto.Extra)
	if err != nil {
		return err
	}

	return u.userRepo.Save(ctx,user)
}