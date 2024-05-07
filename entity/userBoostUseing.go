package entity

import "time"

type UserBoostUsage struct {
	UserID     int
	BoostType  string
	UsageDate  time.Time
	UsageCount int
}
