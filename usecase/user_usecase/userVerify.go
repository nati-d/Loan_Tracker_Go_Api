package userusecase

import (
	"loan_tracker/config"
	"loan_tracker/domain"
)


func (u *UserUsecase) VerifyUser(token string) error {
	
	claims := domain.RegisterClaims{}
	err := config.ValidateToken(token, &claims)

	if err != nil {
		return err
	}

	err = u.UserRepository.RegisterUser(&claims.User)

	if err != nil {
		return err
	}



	return nil
	
	
	
}