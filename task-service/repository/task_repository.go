package repository

import (
	"task-service/models"

	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(task *models.Task) error
	GetTaskByID(id uint, userID uint) (*models.Task, error)
	GetTasksByUserID(userID uint) ([]models.Task, error)
	UpdateTask(task *models.Task) error
	DeleteTask(id uint, userID uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}


func (r *taskRepository) CreateTask(task *models.Task) error {
	return r.db.Create(task).Error
}

func (r *taskRepository) GetTaskByID(id uint, userID uint) (*models.Task, error) {
	var task models.Task
	err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&task).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *taskRepository) GetTasksByUserID(userID uint) ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Where("user_id = ?", userID).Order("id desc").Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) UpdateTask(task *models.Task) error {
	return r.db.Save(task).Error
}

func (r *taskRepository) DeleteTask(id uint, userID uint) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Task{}).Error
}
