package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Log struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Timestamp   time.Time          `bson:"timestamp" json:"timestamp"`
	EventType   string             `bson:"eventType" json:"eventType"`
	UserID      string             `bson:"userId,omitempty" json:"userId,omitempty"`
	Username    string             `bson:"username" json:"username"`
	Details     string             `bson:"details" json:"details"`
	IP          string             `bson:"ip,omitempty" json:"ip,omitempty"`
	Status      string             `bson:"status,omitempty" json:"status,omitempty"`
}


type LogRepository interface {
	LogEvent(log *Log) error
	// GetLogs(page, limit string) ([]Log, int, error)
}

type LogUsecase interface {
	LogEvent(log *Log) error
	// GetLogs(page, limit string) ([]Log, int, error)
}