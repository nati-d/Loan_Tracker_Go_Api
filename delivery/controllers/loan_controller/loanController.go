package loancontroller

import "loan_tracker/domain"


type LoanController struct {
	LoanUsecase domain.LoanUsecase
	UserUsecase domain.UserUsecase
}


func NewLoanController(ls domain.LoanUsecase, us domain.UserUsecase) *LoanController {
	return &LoanController{
		LoanUsecase: ls,
		UserUsecase: us,
	}
}