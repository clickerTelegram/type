package entity

import (
	"time"
)

type Referrals struct {
	ID        int64     `json:"id" gorm:"primary_key"`
	CallerId  int64     `json:"caller_id" gorm:"unique"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at" gorm:"autoUpdateTime:nano"`
}

func (r Referrals) TableName() string {
	return "referrals"
}
