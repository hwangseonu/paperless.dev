package database

import (
	"context"
	"log"
	"time"

	"github.com/hwangseonu/paperless.dev"
	"github.com/hwangseonu/paperless.dev/schema"
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

func (resume *Resume) FromModel() *schema.ResumeResponseSchema {
	s := &schema.ResumeResponseSchema{}
	s.ID = resume.ID.Hex()
	s.Title = resume.Title
	s.Bio = resume.Bio
	s.Public = resume.Public
	s.Template = resume.Template
	s.Skills = resume.Skills
	s.CreatedAt = resume.CreatedAt
	s.UpdatedAt = resume.UpdatedAt

	s.Experiences = make([]schema.ExperienceResponseSchema, len(resume.Experiences))
	for i, exp := range resume.Experiences {
		s.Experiences[i] = schema.ExperienceResponseSchema{
			ID:          exp.ID.Hex(),
			Company:     exp.Company,
			Title:       exp.Title,
			Location:    exp.Location,
			StartDate:   exp.StartDate,
			EndDate:     exp.EndDate,
			Description: exp.Description,
		}
	}

	s.Educations = make([]schema.EducationResponseSchema, len(resume.Educations))
	for i, edu := range resume.Educations {
		s.Educations[i] = schema.EducationResponseSchema{
			ID:         edu.ID.Hex(),
			School:     edu.School,
			Degree:     edu.Degree,
			Major:      edu.Major,
			StartDate:  edu.StartDate,
			EndDate:    edu.EndDate,
			GPA:        edu.GPA,
			Activities: edu.Activities,
		}
	}

	s.Projects = make([]schema.ProjectResponseSchema, len(resume.Projects))
	for i, proj := range resume.Projects {
		s.Projects[i] = schema.ProjectResponseSchema{
			ID:          proj.ID.Hex(),
			Title:       proj.Title,
			Description: proj.Description,
			URL:         proj.URL,
			StartDate:   proj.StartDate,
			EndDate:     proj.EndDate,
			Skills:      proj.Skills,
		}
	}

	return s
}

type ResumeRepository struct {
	collection *mongo.Collection
}

func NewResumeRepository() *ResumeRepository {
	return &ResumeRepository{
		collection: mongoDatabase.Collection("resumes"),
	}
}

func (r *ResumeRepository) FindByID(id bson.ObjectID) (*Resume, error) {
	doc := new(Resume)
	err := r.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(doc)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func (r *ResumeRepository) InsertOne(doc Resume) (*Resume, error) {
	result, err := r.collection.InsertOne(context.Background(), doc)
	if err != nil {
		return nil, err
	}
	doc.ID = result.InsertedID.(bson.ObjectID)
	return &doc, nil
}

func (r *ResumeRepository) UpdateOne(resumeID bson.ObjectID, updateFields bson.M) (*mongo.UpdateResult, error) {
	filter := bson.M{"_id": resumeID}
	update := bson.M{"$set": updateFields}

	result, err := r.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
		return result, paperless.ErrDatabase
	}

	return result, nil
}
