package entity

import "time"

type BoostType uint64

const (
	Tapping BoostType = iota + 1
	FullTank
)

type UserBoostUsageDb struct {
	Id         uint64 `gorm:"primaryKey"`
	UserID     uint
	BoostType  BoostType
	UsageDate  time.Time
	UsageCount uint
}
