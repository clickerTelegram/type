package entity

type ChargeLevelsDB struct {
	Id    int64 `json:"-"`
	Rate  int64 `json:"rate"`
	Price int64 `json:"price"`
}
