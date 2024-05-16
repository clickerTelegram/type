package entity

type EnergyLevelsDB struct {
	Id    int64 `json:"-"`
	Limit int64 `json:"limit"`
	Price int64 `json:"price"`
}
