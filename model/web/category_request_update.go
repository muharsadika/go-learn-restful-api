package web

type CategoryRequestUpdate struct {
	ID   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required,min=1,max=100"`
}
