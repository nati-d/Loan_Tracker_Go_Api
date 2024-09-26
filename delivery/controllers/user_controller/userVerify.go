package usercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"loan_tracker/domain"
)

// VerifyUser handles the user verification process using a token.
func (uc *UserController) VerifyUser(c *gin.Context) {
	// Extract the token from the query parameters
	token := c.Query("token")

	// Verify the user using the provided token
	if err := uc.UserUsecase.VerifyUser(token); err != nil {
		// Respond with an internal server error if verification fails
		c.JSON(http.StatusInternalServerError, domain.NewErrorResponse("Failed to verify user", err))
		return
	}

	// Respond with a success message upon successful verification
	c.JSON(http.StatusOK, domain.NewSuccessResponse("User verified successfully", nil))
}
