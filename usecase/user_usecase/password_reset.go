package userusecase

import (
	"errors"
	"loan_tracker/config"
	"loan_tracker/domain"
)


func (u *UserUsecase) PasswordReset(token string, password string) error {
	claims := &domain.PasswordResetClaims{}
	err := config.ValidateToken(token, claims)
	if err != nil {
		return err
	}

	if claims.Username == "" {
		return errors.New("invalid token")
	}

	err = config.ValidatePassword(password)
	if err != nil {
		return err
	}

	hashedPassword, err := config.HashPassword(password)
	if err != nil {
		return err
	}

	err = u.UserRepository.UpdatePassword(claims.Username, hashedPassword)
	if err != nil {
		return err
	}

	return nil
}