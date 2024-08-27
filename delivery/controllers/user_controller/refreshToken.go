package usercontroller

import (
	"loan_tracker/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RefreshToken handles the refresh token request.
// It retrieves the login claims from the context, validates them,
// and attempts to refresh the user's token.
func (uc *UserController) RefreshToken(ctx *gin.Context) {
	// Retrieve the login claims from the context
	claims, ok := ctx.MustGet("claims").(*domain.LoginClaims)
	if !ok {
		// If the claims are not found or invalid, respond with an unauthorized error
		ctx.JSON(http.StatusUnauthorized, domain.NewErrorResponse("Unauthorized", domain.ErrUnauthorized))
		return
	}

	// Attempt to refresh the token using the provided claims
	token, err := uc.UserUsecase.RefreshToken(*claims)
	if err != nil {
		// If there is an error during token refresh, respond with an internal server error
		ctx.JSON(http.StatusInternalServerError, domain.NewErrorResponse("Failed to refresh token", err))
		return
	}

	// If the token refresh is successful, respond with the new token
	ctx.JSON(http.StatusOK, domain.NewSuccessResponse("Token refreshed successfully", gin.H{"token": token}))
}
