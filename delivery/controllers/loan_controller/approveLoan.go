package loancontroller

import (
	"loan_tracker/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)



// ApproveLoan handles the approval or rejection of a loan
func (lc *LoanController) ApproveLoan(c *gin.Context) {
	// Retrieve claims from context
	claim, ok := c.MustGet("claims").(*domain.LoginClaims)
	if !ok || claim == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Check if the user is an admin
	user, err := lc.UserUsecase.GetUserByUsernameOrEmail(claim.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if user.Role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var input struct {
		NewStatus string `json:"newStatus"`
	}

	

	// Bind JSON body to the input struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Validate the newStatus value
	if input.NewStatus != domain.StatusApproved && input.NewStatus != domain.StatusRejected && input.NewStatus != domain.StatusReturned {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status. Only 'approved', 'rejected', or 'returned' are allowed."})
		return
	}

	// Retrieve the loan ID from the URL parameters
	id := c.Param("id")

	// Check if the loan exists and is pending
	loanData, err := lc.LoanUsecase.GetLoanByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Loan not found"})
		return
	}

	if loanData.Status != domain.StatusPending {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Loan status is not pending"})
		return
	}

	// Call the use case to approve or reject the loan
	err = lc.LoanUsecase.ApproveLoan(id, input.NewStatus)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Respond with success message
	c.JSON(http.StatusOK, gin.H{"message": "Loan status updated successfully"})
}
