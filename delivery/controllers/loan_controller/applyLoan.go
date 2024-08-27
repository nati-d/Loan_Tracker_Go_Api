package loancontroller

import (
	"loan_tracker/domain"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (lc *LoanController) ApplyLoan(ctx *gin.Context) {
	// Retrieve the login claims from the context
	claims, ok := ctx.MustGet("claims").(*domain.LoginClaims)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, domain.NewErrorResponse("Unauthorized", domain.ErrUnauthorized))
		return
	}

	//check if the user loaned before and the loan is approved
	loans, err := lc.LoanUsecase.GetMyLoans(claims.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.NewErrorResponse("Failed to get loans", err))
		return
	}

	for _, loan := range loans {
		if loan.Status == "approved" {
			ctx.JSON(http.StatusBadRequest, domain.NewErrorResponse("You have an approved loan you canot get other before returning it", nil))
			return
		}
	}

	// Retrieve the loan details from the request body
	var loan domain.Loan
	if err := ctx.ShouldBindJSON(&loan); err != nil {
		ctx.JSON(http.StatusBadRequest, domain.NewErrorResponse("Invalid loan details", err))
		return
	}

	if loan.Amount <= 0 {
		ctx.JSON(http.StatusBadRequest, domain.NewErrorResponse("Invalid loan amount", nil))
		return
	}

	if loan.Duration <= 0 {
		ctx.JSON(http.StatusBadRequest, domain.NewErrorResponse("Invalid loan duration", nil))
		return
	}

	loan.Interest = loan.Amount * 0.1 * float64(loan.Duration)

	loan.Status = "pending"

	// Set the user ID of the loan to the user ID from the claims
	loan.Username = claims.Username

	loan.LoanDate = time.Now()

	// Apply the loan
	if err := lc.LoanUsecase.ApplyLoan(&loan); err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.NewErrorResponse("Failed to apply loan", err))
		return
	}

	// Respond with a success message
	ctx.JSON(http.StatusOK, domain.NewSuccessResponse("Loan applied successfully", nil))
}
