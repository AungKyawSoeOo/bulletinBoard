package model

import "time"

type User struct {
	Id            int    `gorm:"type:int;primary_key"`
	Username      string `gorm:"type:varchar(255);not null"`
	Email         string `gorm:"uniqueIndex;not null"`
	Password      string `gorm:"not null"`
	Profile_Photo string `gorm:"type:varchar(255);"`
	Type          string `gorm:"type:varchar(1);not null;"`
	Phone         string
	Address       string
	Date_Of_Birth *time.Time `gorm:"autoCreateTime:false"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	CreateUserId  int     `gorm:"not null"`
	UpdateUserId  int     `gorm:"not null"`
	Posts         []Posts `gorm:"foreignkey:CreateUserId"`
}
