package entity

type LastTapPlayer struct {
	PlayerId int64 `gorm:"primaryKey"`
	LastTime int64
}
