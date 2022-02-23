package models

import "time"

type User struct {
	Id       string    `json:"id" binding:"required" gorm:"primary_key"`
	Password string    `json:"password" binding:"required"`
	Name     string    `json:"name" binding:"required"`
	Created  time.Time `json:"created" gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	Updated  time.Time `json:"updated" gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
