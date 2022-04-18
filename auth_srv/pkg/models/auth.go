package models

import (
	"time"
)

type User struct {
	Id         int64  `json:"id" gorm:"primaryKey"`
	Phone      string `json:"phone"`
	ClientType string `json:"clientType"`
}
type SmsVerify struct {
	ID         int64     `json:"id" gorm:"primaryKey"`
	Phone      string    `json:"phone"`
	Code       string    `json:"code"`
	ExpireTime time.Time `json:"expire_time"`
}
