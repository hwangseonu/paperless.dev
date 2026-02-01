package database

import (
	"context"
	"time"

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
	Description string        `bson:"description,omitempty"`
	Email       string        `bson:"mail,omitempty"`
	URL         string        `bson:"url,omitempty"`
	Image       string        `bson:"image,omitempty"`
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
	s.Description = resume.Description
	s.Email = resume.Email
	s.URL = resume.URL
	s.Image = resume.Image
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
	FindByID(id bson.ObjectID) (*Resume, error)
	FindManyByOwnerID(ownerID bson.ObjectID) ([]Resume, error)
	Update(id bson.ObjectID, schema *schema.ResumeUpdateSchema) (*Resume, error)
	DeleteByID(id bson.ObjectID) (int64, error)
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
		return nil, err
	}

	doc := Resume{
		OwnerID:     userID,
		Title:       schema.Title,
		Description: schema.Description,
		Public:      schema.Public,
		Template:    schema.Template,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	ctx := context.Background()
	result, err := r.collection.InsertOne(ctx, doc)
	if err != nil {
		return nil, err
	}
	doc.ID = result.InsertedID.(bson.ObjectID)
	return &doc, nil
}

func (r *MongoResumeRepository) FindByID(id bson.ObjectID) (*Resume, error) {
	ctx := context.Background()
	filter := bson.M{"_id": id}

	var doc Resume

	if err := r.collection.FindOne(ctx, filter).Decode(&doc); err != nil {
		return nil, err
	}

	return &doc, nil
}

func (r *MongoResumeRepository) FindManyByOwnerID(ownerID bson.ObjectID) ([]Resume, error) {
	ctx := context.Background()
	filter := bson.M{"ownerID": ownerID}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var result []Resume
	if err = cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *MongoResumeRepository) Update(id bson.ObjectID, updateSchema *schema.ResumeUpdateSchema) (*Resume, error) {
	fields := bson.M{}

	if updateSchema.Title != nil {
		fields["title"] = *updateSchema.Title
	}
	if updateSchema.Description != nil {
		fields["description"] = *updateSchema.Description
	}
	if updateSchema.Image != nil {
		fields["image"] = *updateSchema.Image
	}
	if updateSchema.Email != nil {
		fields["email"] = *updateSchema.Email
	}
	if updateSchema.URL != nil {
		fields["url"] = *updateSchema.URL
	}
	if updateSchema.Public != nil {
		fields["public"] = *updateSchema.Public
	}
	if updateSchema.Template != nil {
		fields["template"] = *updateSchema.Template
	}
	if updateSchema.Skills != nil {
		fields["skills"] = *updateSchema.Skills
	}
	if updateSchema.Experiences != nil {
		fields["experiences"] = *updateSchema.Experiences
	}
	if updateSchema.Educations != nil {
		fields["educations"] = *updateSchema.Educations
	}
	if updateSchema.Projects != nil {
		fields["projects"] = *updateSchema.Projects
	}

	fields["updatedAt"] = time.Now()

	ctx := context.Background()
	filter := bson.M{"_id": id}
	update := bson.M{"$set": fields}
	opt := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var resume Resume

	if err := r.collection.FindOneAndUpdate(ctx, filter, update, opt).Decode(&resume); err != nil {
		return nil, err
	}

	return &resume, nil
}

func (r *MongoResumeRepository) DeleteByID(id bson.ObjectID) (int64, error) {
	ctx := context.Background()
	filter := bson.M{"_id": id}
	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}
