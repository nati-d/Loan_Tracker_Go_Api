package userusecase

import "loan_tracker/domain"

type UserUsecase struct {
	UserRepository domain.UserRepository
}


func NewUserUsecase(ur domain.UserRepository) *UserUsecase {
	return &UserUsecase{
		UserRepository: ur,
	}
}
