package domain

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims interface {
	SetExpiryTime()
	GetSecretKey() []byte
	Valid() error
}

type LoginClaims struct {
	Username string `json:"username"`
	ExpiresAt int64  `json:"expires_at"`
	Type    string `json:"type"`
	jwt.StandardClaims
}

type RegisterClaims struct {
	User `json:"user"`
	jwt.StandardClaims
}

type PasswordResetClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}


func (l *LoginClaims) SetExpiryTime() {
	l.ExpiresAt = time.Now().Add(5 * time.Minute).Unix()
}

func (l *LoginClaims) GetSecretKey() []byte {
	return []byte("secret")
}

func (l *LoginClaims) Valid() error {
	return l.StandardClaims.Valid()
}

func (r *RegisterClaims) SetExpiryTime() {
	r.ExpiresAt = time.Now().Add(5 * time.Minute).Unix()
}

func (r *RegisterClaims) GetSecretKey() []byte {
	return []byte("secret")
}

func (r *RegisterClaims) Valid() error {
	return r.StandardClaims.Valid()
}

func (p *PasswordResetClaims) SetExpiryTime() {
	p.ExpiresAt = time.Now().Add(5 * time.Minute).Unix()
}

func (p *PasswordResetClaims) GetSecretKey() []byte {
	return []byte("secret")
}

func (p *PasswordResetClaims) Valid() error {
	return p.StandardClaims.Valid()
}


