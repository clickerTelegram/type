package entity

import "time"

type LastUpdateOnlineDb struct {
	ID         uint   `gorm:"primaryKey"`
	UserID     uint64 `gorm:"unique"`
	TimeOnline int64
}

type LastTapUserDb struct {
	ID       uint   `gorm:"primaryKey"`
	UserID   uint64 `gorm:"unique"`
	CountTap uint64
	TimeTap  time.Time
}
