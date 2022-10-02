package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	Id int32 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}