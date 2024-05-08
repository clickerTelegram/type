package entity

import "time"

type ReferralDb struct {
	ID         uint   `gorm:"primaryKey"`
	UserID     uint64 `gorm:"unique"`
	Caller     uint64
	TimeCreate time.Time
}
