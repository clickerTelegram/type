package entity

type ReferralRewardsDB struct {
	Id     int64 `json:"-"`
	Count  int64 `json:"count"`
	Reward int64 `json:"reward"`
}
