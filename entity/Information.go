package entity

type Information struct {
	Id            uint64 `json:"id"`
	Uuid          string `json:"uuid" gorm:"unique"`
	UserId        string `json:"user_id" gorm:"unique"`
	BalanceId     uint64
	Balance       Balance `gorm:"foreignKey:Id;references:BalanceId"`
	LevelId       uint64
	Level         Level `gorm:"foreignKey:Id;references:LevelId"`
	LevelUser     uint64
	LevelUserData LevelUserData `gorm:"foreignKey:Id;references:LevelUser"`
	TaskUserID    uint64
	TaskUser      TaskUser `gorm:"foreignKey:Id;references:TaskUserID"`
}
