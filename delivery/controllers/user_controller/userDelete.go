package usercontroller

import (
	"loan_tracker/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteUser handles the request to delete a user.
// It validates the user's claims, checks permissions,
// and calls the use case to perform the user deletion.
func (uc *UserController) DeleteUser(ctx *gin.Context) {
	// Retrieve and validate the user's claims
	claims, ok := ctx.MustGet("claims").(*domain.LoginClaims)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, domain.NewErrorResponse("Unauthorized", nil))
		return
	}

	username := claims.Username

	// Ensure the username is valid
	if username == "" {
		ctx.JSON(http.StatusUnauthorized, domain.NewErrorResponse("Unauthorized", nil))
		return
	}

	// Retrieve user details to check role permissions
	user, err := uc.UserUsecase.GetUserByUsernameOrEmail(username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.NewErrorResponse("Failed to retrieve user", err))
		return
	}

	// Check if the user has admin privileges
	if user.Role != "admin" {
		ctx.JSON(http.StatusUnauthorized, domain.NewErrorResponse("Unauthorized", nil))
		return
	}

	// Perform user deletion
	err = uc.UserUsecase.DeleteUser(username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.NewErrorResponse("Failed to delete user", err))
		return
	}

	// Return success response
	ctx.JSON(http.StatusOK, domain.NewSuccessResponse("User deleted successfully", nil))
}
