package loanusecase

import "loan_tracker/domain"



func (l *LoanUsecase) ApplyLoan(loan *domain.Loan) error {
	return l.LoanRepository.ApplyLoan(loan)
}