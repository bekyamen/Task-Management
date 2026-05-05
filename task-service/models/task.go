package models

import "time"

type Task struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	Description string    `json:"description"`
	Status      string    `gorm:"not null;default:'pending'" json:"status"` // pending, in-progress, completed
	UserID      uint      `gorm:"not null;index" json:"userId"`
	CreatedAt   time.Time `json:"createdAt"`
}
