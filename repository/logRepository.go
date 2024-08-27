package repository

import (
	"context"
	"loan_tracker/domain"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LogRepository struct {
	logCollection *mongo.Collection
}

func NewLogRepository(db *mongo.Database) *LogRepository {
	logCollection := db.Collection("logs")
	return &LogRepository{
		logCollection: logCollection,
	}
}

func (lr *LogRepository) LogEvent(log *domain.Log) error {
	_, err := lr.logCollection.InsertOne(context.Background(), log)
	if err != nil {
		return err
	}
	return nil
}

func (lr *LogRepository) GetLogs(page, limit string) ([]domain.Log, int, error) {
	// Convert page and limit to integers
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return nil, 0, err
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return nil, 0, err
	}

	// Create pagination options
	skip := (pageInt - 1) * limitInt
	findOptions := options.Find().
		SetSort(bson.D{{Key: "timestamp", Value: -1}}). // Sort by timestamp descending
		SetSkip(int64(skip)).
		SetLimit(int64(limitInt))

	// Query to get the count of all documents in the collection
	count, err := lr.logCollection.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		return nil, 0, err
	}

	// Query the database with pagination and sorting
	cursor, err := lr.logCollection.Find(context.Background(), bson.M{}, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(context.Background())

	var logs []domain.Log
	for cursor.Next(context.Background()) {
		var log domain.Log
		if err := cursor.Decode(&log); err != nil {
			return nil, 0, err
		}
		logs = append(logs, log)
	}

	return logs, int(count), nil
}
