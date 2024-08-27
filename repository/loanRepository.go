package repository

import (
	"context"
	"loan_tracker/domain"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LoanRepository struct {
	loanCollection *mongo.Collection
	userCollection *mongo.Collection
}

func NewLoanRepository(db *mongo.Database) *LoanRepository {
	loanCollection := db.Collection("loans")
	userCollection := db.Collection("users")
	return &LoanRepository{
		loanCollection: loanCollection,
		userCollection: userCollection,
	}
}

func (lr *LoanRepository) ApplyLoan(loan *domain.Loan) error {
	_, err := lr.loanCollection.InsertOne(context.Background(), loan)
	if err != nil {
		return err
	}
	return nil
}

func (lr *LoanRepository) GetMyLoans(usernameOrEmail string) ([]domain.Loan, error) {
	username := usernameOrEmail

	var loans []domain.Loan
	cursor, err := lr.loanCollection.Find(context.Background(), bson.M{"username": username})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var loan domain.Loan
		if err := cursor.Decode(&loan); err != nil {
			return nil, err
		}
		loans = append(loans, loan)
	}

	return loans, nil

}

func (lr *LoanRepository) GetAllLoans(status string, reverse bool, page, limit string) ([]domain.Loan, int, error) {
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return nil, 0, err
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return nil, 0, err
	}

	var loans []domain.Loan

	// Create a filter for the status
	filter := bson.M{}
	if status != "" {
		filter["status"] = status
	}

	// Create sorting options
	sort := bson.D{}
	if status == "pending" {
		sort = createSortOptions(reverse, true)
	} else {
		sort = createSortOptions(reverse, false)
	}

	// Create pagination options
	skip := (pageInt - 1) * limitInt
	findOptions := options.Find().SetSort(sort).SetSkip(int64(skip)).SetLimit(int64(limitInt))

	// Query the database with pagination and sorting
	cursor, err := lr.loanCollection.Find(context.Background(), filter, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(context.Background())

	// Count the number of fetched documents
	var fetchedCount int
	for cursor.Next(context.Background()) {
		var loan domain.Loan
		if err := cursor.Decode(&loan); err != nil {
			return nil, 0, err
		}
		loans = append(loans, loan)
		fetchedCount++
	}

	return loans, fetchedCount, nil
}

// createSortOptions returns the sorting options based on the status and reverse flag
func createSortOptions(reverse bool, isPending bool) bson.D {
	if isPending {
		if reverse {
			return bson.D{{Key: "createdAt", Value: -1}} // Sort by creation date descending
		}
		return bson.D{{Key: "createdAt", Value: 1}} // Sort by creation date ascending
	}
	if reverse {
		return bson.D{{Key: "createdAt", Value: 1}} // Sort by creation date ascending
	}
	return bson.D{{Key: "createdAt", Value: -1}} // Sort by creation date descending
}



func (lr *LoanRepository) ApproveLoan(id, newStatus string) error {
	// Convert the ID string to an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	// Create a filter for the ID
	filter := bson.M{"_id": objectID}

	// Create an update to set the status to "approved"
	update := bson.M{"$set": bson.M{"status": newStatus}}

	// Update the loan with the provided ID
	_, err = lr.loanCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (lr *LoanRepository) GetLoanByID(id string) (*domain.Loan, error) {
	// Convert the ID string to an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Create a filter for the ID
	filter := bson.M{"_id": objectID}

	// Find the loan with the provided ID
	var loan domain.Loan
	err = lr.loanCollection.FindOne(context.Background(), filter).Decode(&loan)
	if err != nil {
		return nil, err
	}

	return &loan, nil
}
