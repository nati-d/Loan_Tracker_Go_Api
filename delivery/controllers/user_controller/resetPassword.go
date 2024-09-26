package usercontroller

import (
	"loan_tracker/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PasswordReset handles the password reset request.
// It validates the input, extracts the reset token and new password,
// and calls the use case to perform the password reset.
func (uc *UserController) PasswordReset(c *gin.Context) {
	// Define a struct to bind the incoming JSON request
	var req struct {
		Token       string `json:"token"`
		NewPassword string `json:"newPassword"`
	}

	// Extract the token from the query parameter
	req.Token = c.Query("token")

	// Bind the JSON request to the struct
	if err := c.ShouldBindJSON(&req); err != nil {
		// Respond with an error if the input is invalid
		c.JSON(http.StatusBadRequest, domain.NewErrorResponse("Invalid input", err))
		return
	}

	// Call the use case to reset the password
	err := uc.UserUsecase.PasswordReset(req.Token, req.NewPassword)
	if err != nil {
		// Respond with an error if the password reset fails
		c.JSON(http.StatusInternalServerError, domain.NewErrorResponse("Failed to reset password", err))
		return
	}

	// Respond with a success message if the password reset is successful
	c.JSON(http.StatusOK, domain.NewSuccessResponse("Password reset successfully", nil))
}
