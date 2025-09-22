package schema

type ResumeCreateSchema struct {
	Title    string `json:"title" binding:"required,min=1"`
	Bio      string `json:"bio,omitempty"`
	Public   *bool  `json:"public,omitempty"`
	Template string `json:"template,omitempty"`
}
