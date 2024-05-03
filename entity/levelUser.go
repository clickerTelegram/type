package entity

type LevelUser struct {
	Id                int64 `json:"id"`
	UserId            int64 `json:"user_id" gorm:"unique"`
	LevelUserTelegram int64
	LevelUserData     LevelUserData `gorm:"foreignKey:Id;references:LevelUserTelegram"`
	NextLevelExp      int64
}

type LevelUserData struct {
	Id      int64 `json:"id" gorm:"unique"`
	ExpNext int64 `json:"exp_next" gorm:"unique"`
}
