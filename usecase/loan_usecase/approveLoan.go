package loanusecase

func (l *LoanUsecase) ApproveLoan(id, newStatus string) error {
	return l.LoanRepository.ApproveLoan(id, newStatus)
}
