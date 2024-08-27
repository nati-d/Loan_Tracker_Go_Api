package loanusecase

import "loan_tracker/domain"


func (l *LoanUsecase) GetAllLoans(status string , reverse bool, page,limit string) ([]domain.Loan,int, error) {
	return l.LoanRepository.GetAllLoans(status,reverse,page,limit)
}