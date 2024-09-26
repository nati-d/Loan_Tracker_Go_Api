package bootstrap

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func ConnectDatabase(uri string) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	err = client.Connect(context.Background())
	if err != nil {
		return nil, err
	}
	return client, nil

}

func DisconnectDatabase(client *mongo.Client) error {
	err := client.Disconnect(context.Background())
	if err != nil {
		return err
	}
	return nil
}