package model

import "time"

type Posts struct {
	Id           int    `gorm:"type:int;primary_key"`
	Title        string `gorm:"type:varchar(255)"`
	Description  string `gorm:"type:varchar"`
	Status       int    `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	CreateUserId int `gorm:"not null"`
	UpdateUserId int `gorm:"not null"`
}
