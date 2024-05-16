package entity

type OnlineWebPlayer struct {
	PlayerId      int64 `gorm:"primaryKey"`
	TimeOnlineWeb int64
}
