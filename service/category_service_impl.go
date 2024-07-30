package service

import (
	"context"
	"database/sql"
	"muharsadika/golang-belajar-restful-api/exception"
	"muharsadika/golang-belajar-restful-api/helper"
	"muharsadika/golang-belajar-restful-api/model/domain"
	"muharsadika/golang-belajar-restful-api/model/web"
	"muharsadika/golang-belajar-restful-api/repository"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	categoryService := &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validate:           validate,
	}

	return categoryService
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryRequestCreate) web.CategoryResponse {
	errValidate := service.Validate.Struct(request)
	helper.PanicIfError(errValidate)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Save(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

// func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryRequestUpdate) web.CategoryResponse {
// 	errValidate := service.Validate.Struct(request)
// 	helper.PanicIfError(errValidate)

// 	tx, err := service.DB.Begin()
// 	helper.PanicIfError(err)
// 	defer helper.CommitOrRollback(tx)

// 	category, err := service.CategoryRepository.FindByID(ctx, tx, request.ID)
// 	if err != nil {
// 		panic(exception.NewErrorNotFound(err.Error()))
// 	}

// 	category.Name = request.Name

// 	category = service.CategoryRepository.Update(ctx, tx, category)

// 	return helper.ToCategoryResponse(category)
// }

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryRequestUpdate) web.CategoryResponse {
	errValidate := service.Validate.Struct(request)
	helper.PanicIfError(errValidate)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Find category by ID
	category, err := service.CategoryRepository.FindByID(ctx, tx, request.ID)
	if err != nil {
		// Handle category not found error
		panic(exception.NewErrorNotFound(err.Error()))
	}

	// Check if the new name is different
	if category.Name != request.Name {
		// Update category name
		category.Name = request.Name

		// Save updated category
		category = service.CategoryRepository.Update(ctx, tx, category)
	}

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryID int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindByID(ctx, tx, categoryID)
	if err != nil {
		panic(exception.NewErrorNotFound(err.Error()))
	}

	category = service.CategoryRepository.Delete(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindByID(ctx context.Context, categoryID int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindByID(ctx, tx, categoryID)
	if err != nil {
		panic(exception.NewErrorNotFound(err.Error()))
	}

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	return helper.ToCategoriesResponse(categories)
}
