// data akses layer ke domain category (repository pattern)
// best practice buat kontrak seperti file ini repository domain menggunakan interface
// kemudian membuat implementasi dalam bentuk struct
// jadi disini adalah repository function yang bisa support semua api yang dibutuhkan

package repository

import (
	"context"
	"database/sql"
	"muharsadika/golang-belajar-restful-api/model/domain"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	FindByID(ctx context.Context, tx *sql.Tx, categoryID int) (domain.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
}
