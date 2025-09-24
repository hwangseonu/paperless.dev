package database

import (
	"context"
	"errors"
	"time"

	"github.com/hwangseonu/paperless.dev/schema"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
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

func (user *User) ResponseSchema() *schema.UserResponseSchema {
	s := new(schema.UserResponseSchema)
	s.ID = user.ID.Hex()
	s.Username = user.Username
	s.Email = user.Email
	s.Name = user.Name
	s.Bio = user.Bio
	s.ProfileImageURL = user.ProfileImageURL
	s.CreatedAt = user.CreatedAt
	s.UpdatedAt = user.UpdatedAt
	return s
}

type UserRepository interface {
	Create(schema *schema.UserCreateSchema) (*User, error)
	FindByID(id string) (*User, error)
	FindByUsername(username string) (*User, error)
	FindByUsernameOrEmail(username, email string) (*User, error)
	Update(id string, schema *schema.UserUpdateSchema) (*User, error)
	DeleteByID(id string) error
}

func NewUserRepository() UserRepository {
	return &MongoUserRepository{
		collection: mongoDatabase.Collection("users"),
	}
}

type MongoUserRepository struct {
	collection *mongo.Collection
}

func (r *MongoUserRepository) Create(user *schema.UserCreateSchema) (*User, error) {
	doc := &User{
		Username:  user.Username,
		Password:  user.Password,
		Email:     user.Email,
		CreatedAt: time.Now(),
	}

	result, err := r.collection.InsertOne(context.Background(), doc)
	if err != nil {
		return nil, err
	}

	doc.ID = result.InsertedID.(bson.ObjectID)
	return doc, nil
}

func (r *MongoUserRepository) FindByID(id string) (*User, error) {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid user ID format")
	}

	var user User
	err = r.collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *MongoUserRepository) FindByUsername(username string) (*User, error) {
	var user User
	err := r.collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *MongoUserRepository) FindByUsernameOrEmail(username, email string) (*User, error) {
	filter := bson.M{
		"$or": []bson.M{
			{"username": username},
			{"email": email},
		},
	}

	var user User
	err := r.collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *MongoUserRepository) Update(id string, schema *schema.UserUpdateSchema) (*User, error) {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid schema ID format")
	}
	updateFields := bson.M{}

	if schema.Username != nil {
		updateFields["Username"] = *schema.Username
	}
	if schema.Email != nil {
		updateFields["Email"] = *schema.Email
	}
	if schema.Name != nil {
		updateFields["Name"] = *schema.Name
	}
	if schema.Bio != nil {
		updateFields["Bio"] = *schema.Bio
	}

	updateFields["updatedAt"] = time.Now()

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": updateFields}
	opt := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var updatedUser User
	err = r.collection.FindOneAndUpdate(context.TODO(), filter, update, opt).Decode(&updatedUser)
	if err != nil {
		return nil, err
	}

	return &updatedUser, nil
}

func (r *MongoUserRepository) DeleteByID(id string) error {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid user ID format")
	}

	result, err := r.collection.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments // 또는 ErrUserNotFound
	}

	return nil
}
