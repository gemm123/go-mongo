package services

import (
	"context"
	"errors"

	"github.com/gemm123/go-mongo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService interface {
	CreateUser(user models.User) error
	GetAllUser() ([]models.User, error)
	GetUserByName(name string) (models.User, error)
	UpdateUser(name string, user models.User) error
	DeleteUser(name string) error
}

type userService struct {
	userCollection *mongo.Collection
}

func NewUserService(userCollection *mongo.Collection) *userService {
	return &userService{userCollection}
}

func (us *userService) CreateUser(user models.User) error {
	_, err := us.userCollection.InsertOne(context.TODO(), user)
	return err
}

func (us *userService) GetAllUser() ([]models.User, error) {
	var users []models.User
	cursor, err := us.userCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(context.TODO())

	if len(users) == 0 {
		return nil, errors.New("document not found")
	}

	return users, nil
}

func (us *userService) GetUserByName(name string) (models.User, error) {
	var user models.User
	filter := bson.D{bson.E{Key: "name", Value: name}}
	err := us.userCollection.FindOne(context.TODO(), filter).Decode(&user)
	return user, err
}

func (us *userService) UpdateUser(name string, user models.User) error {
	filter := bson.D{bson.E{Key: "name", Value: name}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "name", Value: user.Name}, bson.E{Key: "age", Value: user.Age}, bson.E{Key: "city", Value: user.City}}}}
	result, _ := us.userCollection.UpdateOne(context.TODO(), filter, update)
	if result.MatchedCount != 1 {
		return errors.New("no matched document for update")
	}

	return nil
}

func (us *userService) DeleteUser(name string) error {
	filter := bson.D{bson.E{Key: "name", Value: name}}
	result, _ := us.userCollection.DeleteOne(context.TODO(), filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched document found for delete")
	}

	return nil
}
