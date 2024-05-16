package entity

type OnlineWebPlayer struct {
	PlayerId      int64 `gorm:"primaryKey"`
	TimeOnlineWeb int64
}

type OnlineBotPlayer struct {
	PlayerId      int64 `gorm:"primaryKey"`
	TimeOnlineBot int64
}
