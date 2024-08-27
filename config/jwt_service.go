package config

import (
	"loan_tracker/domain"

	"github.com/golang-jwt/jwt"
)


func GenerateToken(claim domain.Claims) (string,error) {
	claim.SetExpiryTime()
	secretKey := claim.GetSecretKey()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claim)
	signedToken, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return signedToken, nil

}

func ValidateToken(tokenString string, claim domain.Claims) error {
	secretKey := claim.GetSecretKey()
	token, err := jwt.ParseWithClaims(tokenString,claim,func(token *jwt.Token)(interface{},error){
		return secretKey,nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return err
	}

	return nil
}
