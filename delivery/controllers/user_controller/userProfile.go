package usercontroller

import (
	"loan_tracker/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUserProfile retrieves the profile of the currently authenticated user.
// It uses the user's claims to fetch and return the user profile.
func (uc *UserController) GetUserProfile(ctx *gin.Context) {
	// Retrieve and validate user claims from the context
	claims, ok := ctx.MustGet("claims").(*domain.LoginClaims)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, domain.NewErrorResponse("Unauthorized access", nil))
		return
	}

	// Fetch the user profile using the username from claims
	user, err := uc.UserUsecase.GetUserByUsernameOrEmail(claims.Username)
	if err != nil {
		// Handle errors specific to user retrieval
		if err == domain.ErrUserNotFound {
			ctx.JSON(http.StatusNotFound, domain.ErrUserNotFound)
			return
		}
		ctx.JSON(http.StatusInternalServerError, domain.NewErrorResponse("Failed to retrieve user profile", err))
		return
	}

	// Return the user profile in the response
	ctx.JSON(http.StatusOK, domain.NewSuccessResponse("User profile retrieved successfully", user))
}
