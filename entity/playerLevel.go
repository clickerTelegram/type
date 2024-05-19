package entity

type PlayerLevelsDB struct {
	Id     int64 `gorm:"primaryKey"`
	Score  int64 `json:"score"`
	Reward int64 `json:"reward"`
}
