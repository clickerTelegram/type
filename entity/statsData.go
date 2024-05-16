package entity

type StatsData struct {
	PlayerId int64 `gorm:"primaryKey"`
	Taps     int64
	RefIn    int64
	RefOut   int64
	RefCount int64
	Earned   int64
	Reward   int64
	Spent    int64
}
