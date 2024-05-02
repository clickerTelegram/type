package entity

import "time"

type LevelType int

const (
	LeveLTap LevelType = iota + 1
	LevelEnergy
	LevelRecharging
)

type Level struct {
	Id              int64 `json:"id" gorm:"primary_key"`
	UserId          int64 `json:"user_id" gorm:"unique"`
	LeveLTap        LevelData
	LevelEnergy     LevelData
	LevelRecharging LevelData
	BotActive       bool
	TapingGuru      int
	FullTank        int
}

type LevelData struct {
	Id        int64
	Name      string `json:"name" gorm:"unique"`
	Type      LevelType
	Cont      int64
	CreatedAt time.Time `json:"created_at" gorm:"autoUpdateTime:nano"`
}
