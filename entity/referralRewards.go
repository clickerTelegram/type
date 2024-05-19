package entity

type ReferralRewardsDB struct {
	Id     int64 `gorm:"primaryKey"`
	Count  int64 `json:"count"`
	Reward int64 `json:"reward"`
}
