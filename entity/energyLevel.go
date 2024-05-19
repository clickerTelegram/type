package entity

type EnergyLevelsDB struct {
	Id    int64 `gorm:"primaryKey"`
	Limit int64 `json:"limit"`
	Price int64 `json:"price"`
}
