package repositories

import (
	"fmt"
	"strings"

	"github.com/SergeyCherepiuk/todo-app/src/models"
	"github.com/jmoiron/sqlx"
)

type CategoryRepository interface {
	GetById(categoryId uint64) (models.Category, error)
	GetAll(userId uint64) ([]models.Category, error)
	Create(category models.Category) (models.Category, error)
	Update(categoryId uint64, fieldsWithNewValues map[string]any) (models.Category, error)
	Delete(categoryId uint64) error
	DeleteAll(userId uint64) (uint64, error)
}

type CategoryRepositoryImpl struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{db: db}
}

func (repository CategoryRepositoryImpl) GetById(categoryId uint64) (models.Category, error) {
	category := models.Category{}
	sql := "SELECT * FROM categories WHERE id = $1"
	err := repository.db.Get(&category, sql, categoryId)
	return category, err
}

func (repository CategoryRepositoryImpl) GetAll(userId uint64) ([]models.Category, error) {
	categories := []models.Category{}
	sql := "SELECT * FROM categories WHERE user_id = $1"
	err := repository.db.Select(&categories, sql, userId)
	return categories, err
}

func (repository CategoryRepositoryImpl) Create(category models.Category) (models.Category, error) {
	insertedCategory := models.Category{}
	sql := "INSERT INTO categories (name, user_id) VALUES ($1, $2) RETURNING *"
	row := repository.db.QueryRowx(sql, category.Name, category.UserID)
	if row.Err() != nil {
		return insertedCategory, row.Err()
	}
	err := row.StructScan(&insertedCategory)
	return insertedCategory, err
}

func (repository CategoryRepositoryImpl) Update(categoryId uint64, fieldsWithNewValues map[string]any) (models.Category, error) {
	updatedCategory := models.Category{}
	sql := []byte("UPDATE categories SET ")

	updates := []string{}
	for field, newValue := range fieldsWithNewValues {
		switch newValue.(type) {
		case string, byte, rune:
			updates = append(updates, fmt.Sprintf("%s = '%s'", field, newValue))
		default:
			updates = append(updates, fmt.Sprintf("%s = %s", field, newValue))
		}
	}
	sql = append(sql, strings.Join(updates, ", ")...)
	sql = append(sql, "WHERE id = $1 RETURNING *"...)

	row := repository.db.QueryRowx(string(sql), categoryId)
	if row.Err() != nil {
		return updatedCategory, row.Err()
	}
	err := row.StructScan(&updatedCategory)
	return updatedCategory, err
}

func (repository CategoryRepositoryImpl) Delete(categoryId uint64) error {
	sql := "DELETE FROM categories WHERE id = $1"
	res, err := repository.db.Exec(sql, categoryId)
	if err != nil {
		return err
	}
	nRows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if nRows < 1 {
		return fmt.Errorf("category with id = %d not found", categoryId)
	}
	return nil
}

func (repository CategoryRepositoryImpl) DeleteAll(userId uint64) (uint64, error) {
	sql := "DELETE FROM categories WHERE user_id = $1"
	res, err := repository.db.Exec(sql, userId)
	if err != nil {
		return 0, err
	}
	nRows, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return uint64(nRows), nil
}
