package database

import (
	"context"
	"log"
	"time"

	"github.com/hwangseonu/paperless.dev"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type User struct {
	ID              bson.ObjectID `bson:"_id,omitempty"`
	Username        string        `bson:"username"`
	Email           string        `bson:"email"`
	Password        string        `bson:"password,omitempty"`
	Provider        string        `bson:"provider"`
	IsEmailVerified bool          `bson:"isEmailVerified,omitempty"`
	Name            string        `bson:"name,omitempty"`
	Bio             string        `bson:"bio,omitempty"`
	ProfileImageURL string        `bson:"profileImageURL,omitempty"`

	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
	LastLogin time.Time `bson:"lastLogin,omitempty"`
}

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		collection: mongoDatabase.Collection("users"),
	}
}

func (repository *UserRepository) FindByID(id bson.ObjectID) (*User, error) {
	var user User
	err := repository.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repository *UserRepository) FindByUsername(username string) (*User, error) {
	filter := bson.M{"username": username}

	var user User
	err := repository.collection.FindOne(context.Background(), filter).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repository *UserRepository) FindByUsernameOrEmail(username, email string) (*User, error) {
	filter := bson.M{
		"$or": []bson.M{
			{"username": username},
			{"email": email},
		},
	}

	var user User
	err := repository.collection.FindOne(context.Background(), filter).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repository *UserRepository) InsertOne(user User) (*User, error) {
	result, err := repository.collection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}
	user.ID = result.InsertedID.(bson.ObjectID)
	return &user, nil
}

func (repository *UserRepository) UpdateOne(id bson.ObjectID, updateFields bson.M) (*mongo.UpdateResult, error) {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": updateFields}

	result, err := repository.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
		return result, paperless.ErrDatabase
	}

	return result, nil
}

func (repository *UserRepository) DeleteByID(id bson.ObjectID) (*mongo.DeleteResult, error) {
	filter := bson.M{"_id": id}
	result, err := repository.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	return result, nil
}
