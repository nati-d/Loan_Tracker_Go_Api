package usercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func (uc *UserController) PasswordResetRequest(ctx *gin.Context) {
	var input struct {
		Email string `json:"email" binding:"required,email"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := uc.UserUsecase.PasswordResetRequest(input.Email)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Password reset request sent to your email", "user": err})

}