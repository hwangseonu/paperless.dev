package database

import (
	"context"
	"time"

	"github.com/hwangseonu/paperless.dev/internal/schema"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type User struct {
	ID           bson.ObjectID `bson:"_id,omitempty"`
	Nickname     string        `bson:"nickname"`
	Email        string        `bson:"email"`
	ProfileImage string        `bson:"profileImage"`
	Password     string        `bson:"password"`
	CreatedAt    time.Time     `bson:"createdAt"`
	UpdatedAt    time.Time     `bson:"updatedAt"`
}

func (user *User) ResponseSchema() *schema.UserResponseSchema {
	s := new(schema.UserResponseSchema)
	s.ID = user.ID.Hex()
	s.Nickname = user.Nickname
	s.Email = user.Email
	s.CreatedAt = user.CreatedAt
	s.UpdatedAt = user.UpdatedAt
	return s
}

type UserRepository interface {
	Create(schema *schema.UserCreateSchema) (*User, error)
	FindByID(id bson.ObjectID) (*User, error)
	FindByEmail(email string) (*User, error)
	Update(id bson.ObjectID, schema *schema.UserUpdateSchema) (*User, error)
	DeleteByID(id bson.ObjectID) (int64, error)
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
		ID:        bson.NewObjectID(),
		Nickname:  user.Nickname,
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

func (r *MongoUserRepository) FindByID(id bson.ObjectID) (*User, error) {
	var user User

	ctx := context.Background()
	filter := bson.M{"_id": id}
	if err := r.collection.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *MongoUserRepository) FindByEmail(email string) (*User, error) {
	var user User

	ctx := context.Background()
	filter := bson.M{"email": email}
	if err := r.collection.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *MongoUserRepository) Update(id bson.ObjectID, schema *schema.UserUpdateSchema) (*User, error) {
	fields := bson.M{}
	if schema.Nickname != nil {
		fields["nickname"] = *schema.Nickname
	}
	if schema.ProfileImage != nil {
		fields["profileImage"] = *schema.ProfileImage
	}
	fields["updatedAt"] = time.Now()

	ctx := context.Background()
	filter := bson.M{"_id": id}
	update := bson.M{"$set": fields}
	opt := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var user User
	if err := r.collection.FindOneAndUpdate(ctx, filter, update, opt).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *MongoUserRepository) DeleteByID(id bson.ObjectID) (int64, error) {
	ctx := context.Background()
	filter := bson.M{"_id": id}

	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return 0, err

	}
	return result.DeletedCount, nil
}
