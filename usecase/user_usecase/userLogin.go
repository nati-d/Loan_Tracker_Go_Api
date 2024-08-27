package userusecase

import (
	"loan_tracker/config"
	"loan_tracker/domain"
	"time"
)

func (u *UserUsecase) LoginUser(email, password string) (string, string, error) {
	// Fetch the user by username or email
	user, err := u.UserRepository.GetUserByUsernameOrEmail(email)
	if err != nil {
		return "", "", err
	}

	// Compare the provided password with the stored password hash
	err = config.ComparePassword(user.Password, password)
	if err != nil {
		return "", "", err
	}

	// Generate Access Token
	accessClaims := domain.LoginClaims{
		Username:  user.Username,
		ExpiresAt: time.Now().Add(time.Minute * 15).Unix(), 
		Type : "access",
	}
	accessToken, err := config.GenerateToken(&accessClaims)
	if err != nil {
		return "", "", err
	}

	// Generate Refresh Token
	refreshClaims := domain.LoginClaims{
		Username:  user.Username,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(), 
		Type : "refresh",
	}
	refreshToken, err := config.GenerateToken(&refreshClaims)
	if err != nil {
		return "", "", err
	}

	// Save the refresh token in the repository

	token := &domain.Token{
		Username: user.Username,
		ExpiresAt: refreshClaims.ExpiresAt,
	}

	err = u.UserRepository.InsertToken(token)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
