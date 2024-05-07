package entity

type UserSettingsDb struct {
	ID          uint64 `gorm:"primaryKey"`
	UserID      uint64 `gorm:"unique"`
	Leagues     uint
	LevelUser   uint
	LevelTap    uint
	LevelCharge uint
	LevelSpeed  uint
	BotClicker  bool
}
