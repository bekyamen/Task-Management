package models

import "time"

type User struct {
	ID                   uint       `gorm:"primaryKey" json:"id"`
	Email                string     `gorm:"uniqueIndex;not null" json:"email"`
	Password             string     `gorm:"not null" json:"-"`
	TokenVersion         int        `gorm:"default:0" json:"-"`
	ResetTokenHash       *string    `json:"-"`
	ResetTokenExpiresAt *time.Time `json:"-"`
	CreatedAt            time.Time  `json:"createdAt"`
}
