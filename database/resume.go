package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Experience struct {
	ID          bson.ObjectID `bson:"_id,omitempty"`
	Company     string        `bson:"company"`
	Title       string        `bson:"title"`
	Location    string        `bson:"location,omitempty"`
	StartDate   time.Time     `bson:"startDate"`
	EndDate     *time.Time    `bson:"endDate,omitempty"`
	Description string        `bson:"description,omitempty"`
}

type Education struct {
	ID         bson.ObjectID `bson:"_id,omitempty"`
	School     string        `bson:"school"`
	Degree     string        `bson:"degree"`
	Major      string        `bson:"major"`
	StartDate  time.Time     `bson:"startDate"`
	EndDate    *time.Time    `bson:"endDate,omitempty"`
	GPA        string        `bson:"gpa,omitempty"`
	Activities string        `bson:"activities,omitempty"`
}

type Project struct {
	ID          bson.ObjectID `bson:"_id,omitempty"`
	Title       string        `bson:"title"`
	Description string        `bson:"description"`
	URL         string        `bson:"url,omitempty"`
	StartDate   time.Time     `bson:"startDate"`
	EndDate     *time.Time    `bson:"endDate,omitempty"`
	Skills      []string      `bson:"skills,omitempty"`
}

type Resume struct {
	ID          bson.ObjectID `bson:"_id,omitempty"`
	UserID      bson.ObjectID `bson:"userId"`
	Title       string        `bson:"title"`
	Bio         string        `bson:"bio,omitempty"`
	Public      bool          `bson:"public"`
	Template    string        `bson:"template,omitempty"`
	Skills      []string      `bson:"skills,omitempty"`
	Experiences []Experience  `bson:"experiences,omitempty"`
	Educations  []Education   `bson:"educations,omitempty"`
	Projects    []Project     `bson:"projects,omitempty"`
	CreatedAt   time.Time     `bson:"createdAt"`
	UpdatedAt   time.Time     `bson:"updatedAt"`
}

type ResumeRepository struct {
	collection *mongo.Collection
}

func NewResumeRepository() *ResumeRepository {
	return &ResumeRepository{
		collection: mongoDatabase.Collection("resumes"),
	}
}

func (r *ResumeRepository) InsertOne(doc Resume) (*Resume, error) {
	result, err := r.collection.InsertOne(context.Background(), doc)
	if err != nil {
		return nil, err
	}
	doc.ID = result.InsertedID.(bson.ObjectID)
	return &doc, nil
}
