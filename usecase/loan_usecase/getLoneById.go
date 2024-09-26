package loanusecase

import "loan_tracker/domain"


func(lu *LoanUsecase) GetLoanByID(id string) (*domain.Loan, error) {
	loan, err := lu.LoanRepository.GetLoanByID(id)
	if err != nil {
		return nil, err
	}
	return loan, nil
}