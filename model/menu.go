package model

import (
	"gorm.io/gorm"
	"time"
)

type ApiMenuCreateV2 struct {
	ID        uint `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
	Menu string `json:"menu" binding:"required"`
}