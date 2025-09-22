package schema

import (
	"time"

	"github.com/hwangseonu/paperless.dev/database"
)

type ResumeCreateSchema struct {
	Title    string `json:"title" binding:"required,min=1"`
	Bio      string `json:"bio,omitempty"`
	Public   bool   `json:"public,omitempty"`
	Template string `json:"template,omitempty"`
}

type ExperienceResponseSchema struct {
	ID          string     `json:"id"`
	Company     string     `json:"company"`
	Title       string     `json:"title"`
	Location    string     `json:"location,omitempty"`
	StartDate   time.Time  `json:"startDate"`
	EndDate     *time.Time `json:"endDate,omitempty"`
	Description string     `json:"description,omitempty"`
}

type EducationResponseSchema struct {
	ID         string     `json:"id"`
	School     string     `json:"school"`
	Degree     string     `json:"degree"`
	Major      string     `json:"major,omitempty"`
	StartDate  time.Time  `json:"startDate"`
	EndDate    *time.Time `json:"endDate,omitempty"`
	GPA        string     `json:"gpa,omitempty"`
	Activities string     `json:"activities,omitempty"`
}

type ProjectResponseSchema struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description,omitempty"`
	URL         string     `json:"url,omitempty"`
	StartDate   time.Time  `json:"startDate"`
	EndDate     *time.Time `json:"endDate,omitempty"`
	Skills      []string   `json:"skills,omitempty"`
}

type ResumeResponseSchema struct {
	ID          string                     `json:"id"`
	Title       string                     `json:"title"`
	Bio         string                     `json:"bio,omitempty"`
	Public      bool                       `json:"public"`
	Template    string                     `json:"template,omitempty"`
	Skills      []string                   `json:"skills,omitempty"`
	Experiences []ExperienceResponseSchema `json:"experiences,omitempty"`
	Educations  []EducationResponseSchema  `json:"educations,omitempty"`
	Projects    []ProjectResponseSchema    `json:"projects,omitempty"`
	CreatedAt   time.Time                  `json:"createdAt"`
	UpdatedAt   time.Time                  `json:"updatedAt"`
}

func (s *ResumeResponseSchema) FromModel(resume database.Resume) *ResumeResponseSchema {
	s.ID = resume.ID.Hex()
	s.Title = resume.Title
	s.Bio = resume.Bio
	s.Public = resume.Public
	s.Template = resume.Template
	s.Skills = resume.Skills
	s.CreatedAt = resume.CreatedAt
	s.UpdatedAt = resume.UpdatedAt

	s.Experiences = make([]ExperienceResponseSchema, len(resume.Experiences))
	for i, exp := range resume.Experiences {
		s.Experiences[i] = ExperienceResponseSchema{
			ID:          exp.ID.Hex(),
			Company:     exp.Company,
			Title:       exp.Title,
			Location:    exp.Location,
			StartDate:   exp.StartDate,
			EndDate:     exp.EndDate,
			Description: exp.Description,
		}
	}

	s.Educations = make([]EducationResponseSchema, len(resume.Educations))
	for i, edu := range resume.Educations {
		s.Educations[i] = EducationResponseSchema{
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

	s.Projects = make([]ProjectResponseSchema, len(resume.Projects))
	for i, proj := range resume.Projects {
		s.Projects[i] = ProjectResponseSchema{
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
