package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uint           `gorm:"type:int; primaryKey; autoIncrement"`
	CreatedAt time.Time      `gorm:"type:int; autoCreateTime; not null`
	UpdatedAt time.Time      `gorm:"type:int; autoCreateTime; autoUpdateTime; not null`
	DeletedAt gorm.DeletedAt `gorm:"type:int`
}
