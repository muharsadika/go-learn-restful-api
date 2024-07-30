package web

type CategoryRequestCreate struct {
	Name string `json:"name" validate:"required,min=1,max=100" `
}
