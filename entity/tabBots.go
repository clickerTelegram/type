package entity

type TabBotsDB struct {
	Id       int64 `json:"-"`
	Duration int64 `json:"duration"`
	Price    int64 `json:"price"`
}
