package entity

type TaskType uint64

const (
	Referral TaskType = iota + 1
	LevelUser
	Special
	Leagues
)

type TaskListDB struct {
	ID             uint   `gorm:"primaryKey"`
	Name           string `gorm:"unique"`
	Bonus          uint64
	RequiredPoints uint64
	Type           TaskType
}
