package services

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"task-service/models"
	"task-service/repository"

	"github.com/redis/go-redis/v9"
)

type TaskService interface {
	CreateTask(task *models.Task) error
	GetTaskByID(id uint, userID uint) (*models.Task, error)
	GetTasksByUserID(userID uint) ([]models.Task, error)
	UpdateTask(task *models.Task) error
	DeleteTask(id uint, userID uint) error
}

type taskService struct {
	repo        repository.TaskRepository
	redisClient *redis.Client
	ctx         context.Context
}

func NewTaskService(repo repository.TaskRepository, redisClient *redis.Client) TaskService {
	return &taskService{
		repo:        repo,
		redisClient: redisClient,
		ctx:         context.Background(),
	}
}

func (s *taskService) invalidateCache(userID uint) {
	cacheKey := fmt.Sprintf("tasks_user_%d", userID)
	s.redisClient.Del(s.ctx, cacheKey)
}

func (s *taskService) CreateTask(task *models.Task) error {
	err := s.repo.CreateTask(task)
	if err == nil {
		s.invalidateCache(task.UserID)
	}
	return err
}

func (s *taskService) GetTaskByID(id uint, userID uint) (*models.Task, error) {
	return s.repo.GetTaskByID(id, userID)
}

func (s *taskService) GetTasksByUserID(userID uint) ([]models.Task, error) {
	cacheKey := fmt.Sprintf("tasks_user_%d", userID)

	// Try fetching from Redis first
	val, err := s.redisClient.Get(s.ctx, cacheKey).Result()
	if err == nil {
		var tasks []models.Task
		if json.Unmarshal([]byte(val), &tasks) == nil {
			return tasks, nil
		}
	}

	// Fetch from Database
	tasks, err := s.repo.GetTasksByUserID(userID)
	if err != nil {
		return nil, err
	}

	// Store in Redis
	tasksJSON, err := json.Marshal(tasks)
	if err == nil {
		s.redisClient.Set(s.ctx, cacheKey, tasksJSON, time.Minute*10) // Cache for 10 minutes
	}

	return tasks, nil
}

func (s *taskService) UpdateTask(task *models.Task) error {
	err := s.repo.UpdateTask(task)
	if err == nil {
		s.invalidateCache(task.UserID)
	}
	return err
}

func (s *taskService) DeleteTask(id uint, userID uint) error {
	err := s.repo.DeleteTask(id, userID)
	if err == nil {
		s.invalidateCache(userID)
	}
	return err
}
