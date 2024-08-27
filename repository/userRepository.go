package repository

import (
	"context"
	"errors"
	"loan_tracker/domain"
	"regexp"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	userCollection *mongo.Collection
	tokenCollection *mongo.Collection
}



func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		userCollection: db.Collection("users"),
		tokenCollection: db.Collection("tokens"),
	}
}

func filterUser(input string) (string, error) {
	// Regular expression for validating an email address
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	// Regular expression for validating a username
	usernameRegex := `^[a-zA-Z0-9_]+$`
	
	// Compile regular expressions
	emailRe := regexp.MustCompile(emailRegex)
	usernameRe := regexp.MustCompile(usernameRegex)
	
	// Check if input matches email pattern
	if emailRe.MatchString(input) {
		return "email", nil
	}
	
	// Check if input matches username pattern
	if usernameRe.MatchString(input) && len(input) > 0 {
		return "username", nil
	}
	
	// Return an error if input is neither valid email nor username
	return "", errors.New("invalid input: not a valid email or username")
}


func (ur *UserRepository) CheckUserExists(emailorusername string) (bool, error) {
	var user domain.User

	// Create a filter to check for either email or username
	filter := bson.M{
		"$or": []bson.M{
			{"email": emailorusername},
			{"username": emailorusername},
		},
	}

	// Find one user that matches the filter
	err := ur.userCollection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		// If no document is found, return false without error
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		// Return false with the error if there's a different issue
		return false, err
	}

	// If a document is found, return true
	return true, nil
}

func (ur *UserRepository) RegisterUser(user *domain.User) error {
	_,err := ur.userCollection.InsertOne(context.Background(),user)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) GetUserByUsernameOrEmail(input string) (*domain.User, error) {
	// Determine whether the input is an email or username
	inputType, err := filterUser(input)
	if err != nil {
		return nil, err
	}

	var filter bson.M
	if inputType == "email" {
		filter = bson.M{"email": input}
	} else if inputType == "username" {
		filter = bson.M{"username": input}
	} else {
		return nil, errors.New("invalid input type")
	}

	var user domain.User
	err = ur.userCollection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // No user found
		}
		return nil, err // Other errors
	}

	return &user, nil
}

func (ur *UserRepository) InsertToken(token *domain.Token) error {
	_, err := ur.tokenCollection.InsertOne(context.Background(), token)
	if err != nil {
		return err
	}

	return nil

}


func (ur *UserRepository) GetTokenByUserName(username string) (*domain.Token, error) {
	var token domain.Token
	filter := bson.M{"username": username}
	err := ur.tokenCollection.FindOne(context.Background(), filter).Decode(&token)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &token, nil
}
	
func (ur *UserRepository) UpdatePassword(username, password string) error {
	filter := bson.M{"username": username}
	update := bson.M{"$set": bson.M{"password": password}}
	_, err := ur.userCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) GetAllUsers() (*domain.User,error){
	var users domain.User
	cursor, err := ur.userCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		if err := cursor.Decode(&users); err != nil {
			return nil, err
		}
	}
	return &users, nil
}

func (ur *UserRepository) DeleteUser(username string) error {
	filter := bson.M{"username": username}
	_, err := ur.userCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}


