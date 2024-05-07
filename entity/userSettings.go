package entity

type UserSettingsDb struct {
	ID          uint `gorm:"primaryKey"`
	UserID      uint `gorm:"unique"`
	Leagues     uint
	LevelUser   uint
	LevelTap    uint
	LevelCharge uint
	LevelSpeed  uint
	BotClicker  bool
}
