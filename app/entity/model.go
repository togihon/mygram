package entity

import "time"

type MyGramID struct {
	ID uint `gorm:"primaryKey" json:"id"`
}

type MyGramDate struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
