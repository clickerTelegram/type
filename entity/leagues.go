package entity

type LeaguesDB struct {
	Id             int64  `gorm:"primaryKey"`
	Name           string `json:"name"`
	Title          string `json:"title"`
	Score          int64  `json:"score"`
	Reward         int64  `json:"reward"`
	RewardReferral int64  `json:"reward_referral"`
}
