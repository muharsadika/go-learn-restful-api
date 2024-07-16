package repository

import (
	"context"
	"database/sql"
	"errors"
	"muharsadika/golang-belajar-restful-api/helper"
	"muharsadika/golang-belajar-restful-api/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	categoryRepository := &CategoryRepositoryImpl{}

	return categoryRepository
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "INSERT INTO category(name) VALUES (?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.ID = int(id)

	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "UPDATE category set name = ? where id = ?"
	result, err := tx.ExecContext(ctx, SQL, category.Name, category.ID)
	helper.PanicIfError(err)

	rowsAffected, err := result.RowsAffected()
	helper.PanicIfError(err)

	if rowsAffected == 0 {
		panic("category not found")
	}

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "DELETE FROM category where id = ?"
	result, err := tx.ExecContext(ctx, SQL, category.ID)
	helper.PanicIfError(err)

	rowsAffected, err := result.RowsAffected()
	helper.PanicIfError(err)

	if rowsAffected == 0 {
		panic("category not found")
	}

	return category
}

func (repository *CategoryRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, categoryID int) (domain.Category, error) {
	SQL := "SELECT id, name FROM category where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, categoryID)
	helper.PanicIfError(err)
	defer rows.Close()

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.ID, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("category not found")
	}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "SELECT id, name FROM category"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.ID, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}

	return categories
}
