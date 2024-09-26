package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Loan struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `json:"username"`
	Amount   float64            `json:"amount"`
	Interest float64            `json:"interest"`
	Duration int                `json:"duration"`
	LoanDate time.Time          `json:"loan_date"`
	Status   string             `json:"status"`
}

type LoanRepository interface {
	ApplyLoan(loan *Loan) error
	GetMyLoans(username string) ([]Loan, error)
	GetAllLoans(status string, reverse bool, page, limit string) ([]Loan, int, error)
	ApproveLoan(id, newStatus string) error
	GetLoanByID(id string) (*Loan, error)
}

type LoanUsecase interface {
	ApplyLoan(loan *Loan) error
	GetMyLoans(username string) ([]Loan, error)
	GetAllLoans(status string, reverse bool, page, limit string) ([]Loan,int, error)
	ApproveLoan(id, newStatus string) error
	GetLoanByID(id string) (*Loan, error)
}
