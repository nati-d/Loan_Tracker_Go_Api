package loancontroller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"loan_tracker/domain"
)

func (lc *LoanController) GetAllLoans(ctx *gin.Context) {
	// Retrieve claims from the context
	claims, ok := ctx.MustGet("claims").(*domain.LoginClaims)
	if !ok || claims == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	username := claims.Username

	// Retrieve the user based on the username or email
	user, err := lc.UserUsecase.GetUserByUsernameOrEmail(username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Check if the user is an admin
	if user.Role != "admin" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Retrieve status and reverse parameters from the query string
	status := ctx.Query("status")
	reverseParam := ctx.Query("reverse")
	pageParam := ctx.Query("page")
	limitParam := ctx.Query("limit")

	// Convert reverse parameter to boolean
	reverse := false
	if reverseParam != "" {
		var err error
		reverse, err = strconv.ParseBool(reverseParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'reverse' parameter"})
			return
		}
	}

	// Fetch all loans with the provided parameters
	loans, count, err := lc.LoanUsecase.GetAllLoans(status, reverse,pageParam,limitParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with the list of loans
	ctx.JSON(http.StatusOK, gin.H{"loans": loans, "count": count})
}
