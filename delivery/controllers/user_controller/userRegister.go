package usercontroller

import (
	"loan_tracker/domain"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// RegisterUser handles the registration of a new user.
// It validates input, creates the user, and sends a verification email.
func (uc *UserController) RegisterUser(ctx *gin.Context) {
	var user domain.User

	// Bind JSON body to the user struct
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, domain.NewErrorResponse("Invalid JSON provided", err))
		return
	}

	// Validate user input
	if user.Email == "" {
		ctx.JSON(http.StatusUnprocessableEntity, domain.NewErrorResponse("Email is required", nil))
		return
	}

	if user.Password == "" {
		ctx.JSON(http.StatusUnprocessableEntity, domain.NewErrorResponse("Password is required", nil))
		return
	}

	if user.Username == "" {
		ctx.JSON(http.StatusUnprocessableEntity, domain.NewErrorResponse("Username is required", nil))
		return
	}

	// Set default role for the user
	user.Role = "user"

	user.JoinedAt = time.Now()

	// Create user through usecase
	if err := uc.UserUsecase.RegisterUser(&user); err != nil {
		// Handle specific registration errors if needed
		ctx.JSON(http.StatusInternalServerError, domain.NewErrorResponse("Failed to register user", err))
		return
	}

	// Respond with success message
	ctx.JSON(http.StatusCreated, domain.NewSuccessResponse("Verification email has been sent to your email address", nil))
}
