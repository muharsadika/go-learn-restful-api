package web

type CategoryRequestCreate struct {
	Name string `json:"name" validate:"required" `
}
