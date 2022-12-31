package db

import (
	"time"
)

// Cache is a struct that represents the cache table.
type Cache struct {
	ID    string `gorm:"<-:create,type:uuid;primary_key" json:"id"`
	Value string `gorm:"type:text;not null" json:"value"`

	CreatedAt time.Time `gorm:"type:timestamp;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null" json:"updated_at"`
}
