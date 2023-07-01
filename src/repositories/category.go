package repositories

import (
	"fmt"
	"strings"

	"github.com/SergeyCherepiuk/todo-app/src/models"
	"github.com/jmoiron/sqlx"
)

type CategoryRepository interface {
	GetById(uint64) (models.Category, error)
	GetAll() ([]models.Category, error)
	Create(models.Category) (models.Category, error)
	Update(uint64, map[string]any) (models.Category, error)
	Delete(uint64) error
	DeleteAll() (uint64, error)
}

type CategoryRepositoryImpl struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{db: db}
}

func (repository CategoryRepositoryImpl) GetById(id uint64) (models.Category, error) {
	category := models.Category{}
	sql := "SELECT * FROM category WHERE id = $1"
	err := repository.db.Get(&category, sql, id)
	return category, err
}

func (repository CategoryRepositoryImpl) GetAll() ([]models.Category, error) {
	categories := []models.Category{}
	sql := "SELECT * FROM category"
	err := repository.db.Select(&categories, sql)
	return categories, err
}

func (repository CategoryRepositoryImpl) Create(category models.Category) (models.Category, error) {
	insertedCategory := models.Category{}
	sql := "INSERT INTO category (name) VALUE ($1) RETURNING *"
	row := repository.db.QueryRowx(sql, category.Name)
	if row.Err() != nil {
		return insertedCategory, row.Err()
	}
	err := row.StructScan(&insertedCategory)
	return insertedCategory, err
}

func (repository CategoryRepositoryImpl) Update(id uint64, fieldsWithNewValues map[string]any) (models.Category, error) {
	updatedCategory := models.Category{}
	sql := []byte("UPDATE category SET ")

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

	row := repository.db.QueryRowx(string(sql), id)
	if row.Err() != nil {
		return updatedCategory, row.Err()
	}
	err := row.StructScan(&updatedCategory)
	return updatedCategory, err
}

func (repository CategoryRepositoryImpl) Delete(id uint64) error {
	sql := "DELETE FROM category WHERE id = $1"
	res, err := repository.db.Exec(sql, id)
	if err != nil {
		return err
	}
	nRows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if nRows < 1 {
		return fmt.Errorf("category with id = %d not found", id)
	}
	return nil
}

func (repository CategoryRepositoryImpl) DeleteAll() (uint64, error) {
	sql := "DELETE FROM category"
	res, err := repository.db.Exec(sql)
	if err != nil {
		return 0, err
	}
	nRows, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return uint64(nRows), nil
}
