package entity

type Balance struct {
	Id      int64
	UserId  int64 `json:"user_id" gorm:"unique"`
	Tap     int64
	Coin    int64 `json:"balance"`
	CoinAll int64 `json:"balance_all"`
}
