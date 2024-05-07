package entity

type LevelType int64

const (
	Tap LevelType = iota + 1
	LeveUser
	LevelEnergy
	Recharging
)

type LevelGame struct {
	ID                uint   `gorm:"primaryKey"`
	Name              string `gorm:"unique"`
	Number            uint
	RequiredCoins     uint
	Type              LevelType
	RewardDescription string
}
