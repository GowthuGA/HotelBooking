package models

import (
    "time"
)

type Admin struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    Username  string    `json:"username" gorm:"unique;not null"`
    Password  string    `json:"password" gorm:"not null"`  
    Email     string    `json:"email" gorm:"unique;not null"`
    CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
    LastLogin *time.Time `json:"last_login" gorm:"default:NULL"`
    Status    string    `json:"status" gorm:"default:'active'"`
}
