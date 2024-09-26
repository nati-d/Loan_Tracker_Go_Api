package loancontroller

import (
	"loan_tracker/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (lc *LoanController) GetMyLoans(c *gin.Context) {
	claims := c.MustGet("claims").(*domain.LoginClaims)
	if claims == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	username := claims.Username
	loans, err := lc.LoanUsecase.GetMyLoans(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, loans)
}