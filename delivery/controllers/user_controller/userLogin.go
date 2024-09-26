package usercontroller

import (
	"loan_tracker/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login handles user login requests.
// It validates the provided JSON, checks user credentials,
// and returns access and refresh tokens upon successful authentication.
func (uc *UserController) Login(ctx *gin.Context) {
	// Bind JSON body to the user struct
	var user domain.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, domain.NewErrorResponse("Invalid JSON provided", err))
		return
	}

	// Validate user input
	if user.Email == "" && user.Username == "" {
		ctx.JSON(http.StatusUnprocessableEntity, domain.NewErrorResponse("Email or Username is required", nil))
		return
	}

	if user.Password == "" {
		ctx.JSON(http.StatusUnprocessableEntity, domain.NewErrorResponse("Password is required", nil))
		return
	}

	// Attempt to log in based on email or username
	var accessToken, refreshToken string
	var err error

	if user.Email != "" {
		accessToken, refreshToken, err = uc.UserUsecase.LoginUser(user.Email, user.Password)
	} else {
		accessToken, refreshToken, err = uc.UserUsecase.LoginUser(user.Username, user.Password)
	}

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, domain.NewErrorResponse("Invalid email/username or password", err))
		return
	}

	// Return the tokens to the client
	ctx.JSON(http.StatusOK, domain.NewSuccessResponse("Login successful", gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}))
}
