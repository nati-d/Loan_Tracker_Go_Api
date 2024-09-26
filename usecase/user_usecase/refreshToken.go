package userusecase

import (
	"loan_tracker/config"
	"loan_tracker/domain"
)


func (u *UserUsecase) RefreshToken(claims domain.LoginClaims) (string, error) {
	// get user by id
	user, err := u.UserRepository.GetTokenByUserName(claims.Username)
	if err != nil {
		return "", err
	}

	newToken := &domain.LoginClaims{
		Username:  user.Username,
		ExpiresAt: claims.ExpiresAt,
		Type:	  "access",	
	}

	// generate new token
	token, err := config.GenerateToken(newToken)
	if err != nil {
		return "", err
	}

	return token, nil
}
