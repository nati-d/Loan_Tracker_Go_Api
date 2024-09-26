package loanusecase

import "loan_tracker/domain"


func(l *LoanUsecase) GetMyLoans(usernameOrEmail string) ([]domain.Loan, error) {
	return l.LoanRepository.GetMyLoans(usernameOrEmail)
}