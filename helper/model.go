package helper

import (
	"muharsadika/golang-belajar-restful-api/model/domain"
	"muharsadika/golang-belajar-restful-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}
}

func ToCategoriesResponse(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse

	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}

	return categoryResponses
}
