package databaseBase

import (
	"time"
)

type (
	BaseEntity struct {
		ID        uint      `gorm:"primaryKey" json:"id"`
		CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
		UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	}
)
