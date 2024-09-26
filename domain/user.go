package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `json:"username"`
	Password string             `json:"password"`
	Email    string             `json:"email"`
	Role     string             `json:"role"`
	Avatar   string             `json:"avatar"`
	JoinedAt time.Time            `json:"joined_at"`
	Address  string             `json:"address"`
}

type UserUsecase interface {
	RegisterUser(user *User) error
	VerifyUser(token string) error
	LoginUser(usernameoremail, password string) (string, string, error)
	RefreshToken(claims LoginClaims) (string, error)
	PasswordResetRequest(email string) error
	PasswordReset(token string, password string) error
	GetAllUsers() (*User, error)
	DeleteUser(username string) error
	GetUserByUsernameOrEmail(usernameoremail string) (*User, error)
}

type UserRepository interface {
	RegisterUser(user *User) error
	CheckUserExists(usernameoremail string) (bool, error)
	GetUserByUsernameOrEmail(usernameoremail string) (*User, error)
	InsertToken(token *Token) error
	GetTokenByUserName(username string) (*Token, error)
	UpdatePassword(username, password string) error
	GetAllUsers() (*User, error)
	DeleteUser(username string) error
}
