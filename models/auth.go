package models

import (
	"github.com/jinzhu/gorm"
)

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserMessage struct {
	ID                int    `gorm:"primary_key" json:"id"`
	Username          string `json:"username"`
	Nickname          string `json:"nickname"`
	Email             string `json:"email"`
	Organization      string `json:"organization"`
	Activetime        string `json:"active_time"`
	ChatStatus        string `json:"chat_status"`
	CanUsePaid        int    `json:"can_use_paid"`
	MaxConvCount      int    `json:"max_conv_count"`
	AvailableAskCount int    `json:"available_ask_count"`
	IsSuperuser       int    `json:"is_superuser"`
	IsActive          int    `json:"is_active"`
	IsVerified        int    `json:"is_verified"`
	HashedPassword    string `json:"hashed_password"`
	FunctionList      string `json:"function_list"`
}

// CheckAuth checks if authentication information exists
func CheckAuth(username, password string) []UserMessage {
	// var auth Auth
	var info UserMessage
	var userArrays []UserMessage
	err := db.Table("user").Model(&info).Where("username = ?", username).First(&userArrays).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return userArrays
	}
	return userArrays
}
