package database

import (
	"context"
	"errors"
	"time"

	"github.com/hwangseonu/paperless.dev/internal/common"
	"github.com/hwangseonu/paperless.dev/internal/schema"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
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
	OwnerID     bson.ObjectID `bson:"ownerID"`
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

func (resume *Resume) ResponseSchema() *schema.ResumeResponseSchema {
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

type ResumeRepository interface {
	Create(schema *schema.ResumeCreateSchema) (*Resume, error)
	FindByID(id string) (*Resume, error)
	FindManyByOwnerID(ownerID string) ([]Resume, error)
	Update(id string, schema *schema.ResumeUpdateSchema) (*Resume, error)
	DeleteByID(id string) error
}

type MongoResumeRepository struct {
	collection *mongo.Collection
}

func NewResumeRepository() ResumeRepository {
	return &MongoResumeRepository{
		collection: mongoDatabase.Collection("resumes"),
	}
}

func (r *MongoResumeRepository) Create(schema *schema.ResumeCreateSchema) (*Resume, error) {
	userID, err := bson.ObjectIDFromHex(schema.OwnerID)

	if err != nil {
		return nil, common.ErrInvalidUserID
	}

	doc := Resume{
		OwnerID:   userID,
		Title:     schema.Title,
		Bio:       schema.Bio,
		Public:    schema.Public,
		Template:  schema.Template,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	result, err := r.collection.InsertOne(context.Background(), doc)
	if err != nil {
		return nil, common.ErrDatabase
	}
	doc.ID = result.InsertedID.(bson.ObjectID)
	return &doc, nil
}

func (r *MongoResumeRepository) FindByID(id string) (*Resume, error) {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, common.ErrInvalidResumeID
	}

	doc := new(Resume)
	err = r.collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(doc)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, common.ErrResumeNotFound
		}
		return nil, common.ErrDatabase
	}
	return doc, nil
}

func (r *MongoResumeRepository) FindManyByOwnerID(ownerID string) ([]Resume, error) {
	ownerObjID, err := bson.ObjectIDFromHex(ownerID)
	if err != nil {
		return nil, common.ErrInvalidUserID
	}
	filter := bson.M{"ownerID": ownerObjID}

	cursor, err := r.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, common.ErrDatabase
	}

	var result []Resume
	if err = cursor.All(context.Background(), &result); err != nil {
		return nil, common.ErrDatabase
	}

	return result, nil
}

func (r *MongoResumeRepository) Update(id string, updateSchema *schema.ResumeUpdateSchema) (*Resume, error) {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, common.ErrInvalidResumeID
	}

	updateFields := bson.M{}

	if updateSchema.Title != nil {
		updateFields["title"] = *updateSchema.Title
	}
	if updateSchema.Bio != nil {
		updateFields["bio"] = *updateSchema.Bio
	}
	if updateSchema.Public != nil {
		updateFields["public"] = *updateSchema.Public
	}
	if updateSchema.Template != nil {
		updateFields["template"] = *updateSchema.Template
	}
	if updateSchema.Skills != nil {
		updateFields["skills"] = *updateSchema.Skills
	}
	if updateSchema.Experiences != nil {
		updateFields["experiences"] = *updateSchema.Experiences
	}
	if updateSchema.Educations != nil {
		updateFields["educations"] = *updateSchema.Educations
	}
	if updateSchema.Projects != nil {
		updateFields["projects"] = *updateSchema.Projects
	}

	updateFields["updatedAt"] = time.Now()

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": updateFields}
	opt := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var updatedResume Resume

	err = r.collection.FindOneAndUpdate(context.Background(), filter, update, opt).Decode(&updatedResume)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, common.ErrResumeNotFound
		}
		return nil, common.ErrDatabase
	}

	return &updatedResume, nil
}

func (r *MongoResumeRepository) DeleteByID(id string) error {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return common.ErrInvalidResumeID
	}

	result, err := r.collection.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return common.ErrDatabase
	}

	if result.DeletedCount == 0 {
		return common.ErrResumeNotFound
	}

	return nil
}
