package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID         uint
	Name       string
	CategoryID uint
	Category   Category `gorm:"foreignKey:CategoryID"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}
