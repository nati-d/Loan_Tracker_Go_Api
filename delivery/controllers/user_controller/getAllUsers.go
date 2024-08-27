package usercontroller

import (
	"loan_tracker/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *UserController) GetAllUsers(ctx *gin.Context) {
	// Retrieve the login claims from the context
	claims, ok := ctx.MustGet("claims").(*domain.LoginClaims)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, domain.NewErrorResponse("Unauthorized", domain.ErrUnauthorized))
		return
	}

	username := claims.Username
	if username == "" {
		ctx.JSON(http.StatusUnauthorized, domain.NewErrorResponse("Unauthorized", domain.ErrUnauthorized))
		return
	}

	// Fetch the user using the username or email
	user, err := uc.UserUsecase.GetUserByUsernameOrEmail(username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.NewErrorResponse("Failed to fetch user", err))
		return
	}

	// Check if the user has admin role
	if user.Role != "admin" {
		ctx.JSON(http.StatusUnauthorized, domain.NewErrorResponse("Unauthorized", domain.ErrUnauthorized))
		return
	}

	// Fetch all users
	users, err := uc.UserUsecase.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.NewErrorResponse("Failed to fetch users", err))
		return
	}

	// Return a success response with the list of users
	ctx.JSON(http.StatusOK, domain.NewSuccessResponse("Users fetched successfully", users))
}
