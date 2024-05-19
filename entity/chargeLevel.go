package entity

type ChargeLevelsDB struct {
	Id    int64 `json:"-" gorm:"primaryKey"`
	Rate  int64 `json:"rate"`
	Price int64 `json:"price"`
}
