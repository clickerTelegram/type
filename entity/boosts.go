package entity

type BoostDBType string

const (
	Energy BoostDBType = "energy"
	Turbo  BoostDBType = "turbo"
)

type Boosts struct {
	Id        int64 `json:"-" gorm:"primaryKey"`
	Type      BoostDBType
	Duration  int64 `json:"duration"`
	RateMulti int64 `json:"rate_multi"`
	Energy    int64 `json:"energy"`
}
