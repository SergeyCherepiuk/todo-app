package repositories

import (
	"fmt"
	"strings"

	"github.com/SergeyCherepiuk/todo-app/src/models"
	"github.com/jmoiron/sqlx"
)

type CategoryRepository interface {
	HasCategory(categoryId uint64, userId uint64) (bool, error)
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

func (repository CategoryRepositoryImpl) HasCategory(categoryId, userId uint64) (bool, error) {
	query := `SELECT COUNT(*) FROM categories WHERE id = :categoryId AND user_id = :userId`
	namedParams := map[string]any{
		"categoryId": categoryId,
		"userId":     userId,
	}

	stmt, err := repository.db.PrepareNamed(query)
	if err != nil {
		return false, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	var count uint64
	if err := stmt.Get(&count, namedParams); err != nil {
		return false, fmt.Errorf("failed to execute query: %w", err)
	}

	return count > 0, nil
}

func (repository CategoryRepositoryImpl) GetById(categoryId uint64) (models.Category, error) {
	query := `SELECT * FROM categories WHERE id = :categoryId`
	namedParams := map[string]any{
		"categoryId": categoryId,
	}

	stmt, err := repository.db.PrepareNamed(query)
	if err != nil {
		return models.Category{}, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	category := models.Category{}
	if err := stmt.Get(&category, namedParams); err != nil {
		return models.Category{}, fmt.Errorf("failed to execute query: %w", err)
	}

	return category, nil
}

func (repository CategoryRepositoryImpl) GetAll(userId uint64) ([]models.Category, error) {
	query := `SELECT * FROM categories WHERE user_id = :userId`
	namedParams := map[string]any{
		"userId": userId,
	}

	stmt, err := repository.db.PrepareNamed(query)
	if err != nil {
		return []models.Category{}, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	categories := []models.Category{}
	if err := stmt.Select(&categories, namedParams); err != nil {
		return []models.Category{}, fmt.Errorf("failed to execute query: %w", err)
	}

	return categories, nil
}

func (repository CategoryRepositoryImpl) Create(category models.Category) (models.Category, error) {
	query := `INSERT INTO categories (name, user_id) VALUES (:name, :userId) RETURNING *`
	namedParams := map[string]any{
		"name":   category.Name,
		"userId": category.UserID,
	}

	stmt, err := repository.db.PrepareNamed(query)
	if err != nil {
		return models.Category{}, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	insertedCategory := models.Category{}
	if err := stmt.Get(&insertedCategory, namedParams); err != nil {
		return models.Category{}, fmt.Errorf("failed to execute query: %w", err)
	}

	return insertedCategory, nil
}

func (repository CategoryRepositoryImpl) Update(categoryId uint64, fieldsWithNewValues map[string]any) (models.Category, error) {
	query := []byte("UPDATE categories SET ")
	namedParams := map[string]any{
		"categoryId": categoryId,
	}

	updates := []string{}
	for field, newValue := range fieldsWithNewValues {
		switch newValue.(type) {
		case string, byte, rune:
			updates = append(updates, fmt.Sprintf("%s = '%s'", field, newValue))
		default:
			updates = append(updates, fmt.Sprintf("%s = %s", field, newValue))
		}
	}
	query = append(query, strings.Join(updates, ", ")...)
	query = append(query, "WHERE id = :categoryId RETURNING *"...)

	stmt, err := repository.db.PrepareNamed(string(query))
	if err != nil {
		return models.Category{}, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	updatedCategory := models.Category{}
	if err := stmt.Get(&updatedCategory, namedParams); err != nil {
		return models.Category{}, fmt.Errorf("failed to execute query: %w", err)
	}

	return updatedCategory, nil
}

func (repository CategoryRepositoryImpl) Delete(categoryId uint64) error {
	query := `DELETE FROM categories WHERE id = :categoryId`
	namedParams := map[string]any{
		"categoryId": categoryId,
	}

	stmt, err := repository.db.PrepareNamed(query)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(namedParams)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}
	nRows, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get number of rows: %w", err)
	}
	if nRows < 1 {
		return fmt.Errorf("failed to find category with id %d", categoryId)
	}
	return nil
}

func (repository CategoryRepositoryImpl) DeleteAll(userId uint64) (uint64, error) {
	query := `DELETE FROM categories WHERE user_id = :userId`
	namedParams := map[string]any{
		"userId": userId,
	}

	stmt, err := repository.db.PrepareNamed(query)
	if err != nil {
		return 0, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(namedParams)
	if err != nil {
		return 0, fmt.Errorf("failed to execute query: %w", err)
	}
	nRows, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("failed to get number of rows: %w", err)
	}
	return uint64(nRows), nil
}
