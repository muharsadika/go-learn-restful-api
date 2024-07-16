// bussiness logic

package service

import (
	"context"
	"muharsadika/golang-belajar-restful-api/model/web"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryRequestCreate) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryRequestUpdate) web.CategoryResponse
	Delete(ctx context.Context, categoryID int) web.CategoryResponse
	FindByID(ctx context.Context, categoryID int) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}
