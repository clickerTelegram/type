package entity

import "time"

type ReferralDb struct {
	ID         uint `gorm:"primaryKey"`
	UserID     uint `gorm:"unique"`
	Caller     uint
	TimeCreate time.Time
}
