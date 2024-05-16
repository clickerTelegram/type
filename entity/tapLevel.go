package entity

type TapLevelsDB struct {
	Id     int64 `json:"-"`
	Rate   int64 `json:"rate"`
	Energy int64 `json:"energy"`
	Price  int64 `json:"price"`
}
