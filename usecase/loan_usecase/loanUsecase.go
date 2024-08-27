package loanusecase

import "loan_tracker/domain"


type LoanUsecase struct {
	LoanRepository domain.LoanRepository
}

func NewLoanUsecase(lr domain.LoanRepository) *LoanUsecase {
	return &LoanUsecase{
		LoanRepository: lr,
	}
}