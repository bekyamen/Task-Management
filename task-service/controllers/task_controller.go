package controllers

import (
	"net/http"
	"strconv"

	"task-service/models"
	"task-service/services"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	service services.TaskService
}

func NewTaskController(service services.TaskService) *TaskController {
	return &TaskController{service}
}

func (c *TaskController) CreateTask(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(uint)

	var task models.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.UserID = userID
	if err := c.service.CreateTask(&task); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	ctx.JSON(http.StatusCreated, task)
}

func (c *TaskController) GetTasks(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(uint)

	tasks, err := c.service.GetTasksByUserID(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}

	// Always return array even if empty instead of null
	if tasks == nil {
		tasks = []models.Task{}
	}
	ctx.JSON(http.StatusOK, tasks)
}

func (c *TaskController) UpdateTask(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(uint)
	taskID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	task, err := c.service.GetTaskByID(uint(taskID), userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	var updateData models.Task
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.Title = updateData.Title
	task.Description = updateData.Description
	task.Status = updateData.Status

	if err := c.service.UpdateTask(task); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	ctx.JSON(http.StatusOK, task)
}

func (c *TaskController) DeleteTask(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(uint)
	taskID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	if err := c.service.DeleteTask(uint(taskID), userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task or task not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
