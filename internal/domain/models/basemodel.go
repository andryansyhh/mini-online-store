package models

import (
	"time"
)

type BaseModelUuid struct {
	Uuid      string    `gorm:"primaryKey;size:36;" json:"uuid,omitempty" uri:"uuid"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at,omitempty" uri:"created_at"`
	CreatedBy string    `gorm:"column:created_by;size:50" json:"created_by,omitempty" uri:"created_by"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime:milli" json:"updated_at,omitempty" uri:"updated_at"`
	UpdatedBy string    `gorm:"column:updated_by;size:50" json:"updated_by,omitempty" uri:"updated_by"`
	DeletedAt time.Time `gorm:"column:deleted_at" json:"deleted_at,omitempty" uri:"deleted_at"`
	DeletedBy string    `gorm:"column:deleted_by;size:50" json:"deleted_by,omitempty" uri:"deleted_by"`
}