package schema

import (
	"time"
)

type ResumeCreateSchema struct {
	Title       string `json:"title" binding:"required,min=1"`
	Description string `json:"description,omitempty"`
	Public      bool   `json:"public,omitempty"`
	Template    string `json:"template,omitempty"`
	OwnerID     string `json:"-"`
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
	Description string                     `json:"description"`
	Image       string                     `json:"image,omitempty"`
	Email       string                     `json:"email,omitempty"`
	URL         string                     `json:"url,omitempty"`
	Public      bool                       `json:"public"`
	Template    string                     `json:"template,omitempty"`
	Skills      []string                   `json:"skills,omitempty"`
	Experiences []ExperienceResponseSchema `json:"experiences,omitempty"`
	Educations  []EducationResponseSchema  `json:"educations,omitempty"`
	Projects    []ProjectResponseSchema    `json:"projects,omitempty"`
	CreatedAt   time.Time                  `json:"createdAt"`
	UpdatedAt   time.Time                  `json:"updatedAt"`
}

type ExperienceUpdateSchema struct {
	Company     *string    `json:"company,omitempty"`
	Title       *string    `json:"title,omitempty"`
	Location    *string    `json:"location,omitempty"`
	StartDate   *time.Time `json:"startDate,omitempty"`
	EndDate     *time.Time `json:"endDate,omitempty"`
	Description *string    `json:"description,omitempty"`
}

type EducationUpdateSchema struct {
	School     *string    `json:"school,omitempty"`
	Degree     *string    `json:"degree,omitempty"`
	Major      *string    `json:"major,omitempty"`
	StartDate  *time.Time `json:"startDate,omitempty"`
	EndDate    *time.Time `json:"endDate,omitempty"`
	GPA        *string    `json:"gpa,omitempty"`
	Activities *string    `json:"activities,omitempty"`
}

type ProjectUpdateSchema struct {
	Title       *string    `json:"title,omitempty"`
	Description *string    `json:"description,omitempty"`
	URL         *string    `json:"url,omitempty"`
	StartDate   *time.Time `json:"startDate,omitempty"`
	EndDate     *time.Time `json:"endDate,omitempty"`
	Skills      *[]string  `json:"skills,omitempty"`
}

type ResumeUpdateSchema struct {
	Title       *string                   `json:"title,omitempty"`
	Description *string                   `json:"description,omitempty"`
	Email       *string                   `json:"email,omitempty"`
	URL         *string                   `json:"url,omitempty"`
	Image       *string                   `json:"image,omitempty"`
	Public      *bool                     `json:"public,omitempty"`
	Template    *string                   `json:"template,omitempty"`
	Skills      *[]string                 `json:"skills,omitempty"`
	Experiences *[]ExperienceUpdateSchema `json:"experiences,omitempty"`
	Educations  *[]EducationUpdateSchema  `json:"educations,omitempty"`
	Projects    *[]ProjectUpdateSchema    `json:"projects,omitempty"`
}
