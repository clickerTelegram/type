package entity

type TransformDb struct {
	ID         uint   `gorm:"primaryKey"`
	UserID     uint64 `gorm:"unique"`
	TapAll     uint64
	Balance    uint64
	BalanceAll uint64
}
