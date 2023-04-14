package repository

import (
	"github.com/kudaibergenoff/task-app/internal/task/domain"
	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db}
}

func (repo *TaskRepository) GetAll() ([]domain.Task, error) {
	var tasks []domain.Task
	result := repo.db.Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}
