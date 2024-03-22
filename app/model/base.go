package model

import "time"

type Base struct {
	ID        uint64    `gorm:"column:id;primaryKey" json:"id,omitempty"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt,omitempty"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt,omitempty"`
}
