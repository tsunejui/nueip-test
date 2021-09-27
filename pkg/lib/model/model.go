package model

import "time"

type DefaultTimeFields struct {
	CreatedAt time.Time `json:"created_at" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updated_at" swaggerignore:"true"`
	DeletedAt time.Time `sql:"index" json:"deleted_at" swaggerignore:"true"`
}

type Header struct {
	ID      uint32 `gorm:"primary_key" json:"id" swaggerignore:"true"`
	Version int32  `json:"version" swaggerignore:"true"`
	DefaultTimeFields
}
