package repositories

import (
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
	return models.Category{}, nil
}

func (repository CategoryRepositoryImpl) GetAll() ([]models.Category, error) {
	return []models.Category{}, nil
}

func (repository CategoryRepositoryImpl) Create(models.Category) (models.Category, error) {
	return models.Category{}, nil
}

func (repository CategoryRepositoryImpl) Update(id uint64, fieldsWithNewValues map[string]any) (models.Category, error) {
	return models.Category{}, nil
}

func (repository CategoryRepositoryImpl) Delete(id uint64) error {
	return nil
}

func (repository CategoryRepositoryImpl) DeleteAll() (uint64, error) {
	return 0, nil
}