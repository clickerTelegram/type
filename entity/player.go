package entity

type Player struct {
	Id            int64  `json:"id"`
	PlayerId      int64  `json:"player_id" gorm:"primaryKey"`
	UserName      string `json:"user_name"`
	Name          string `json:"name"`
	WalletAddress string `json:"wallet_address"`
	CreateAt      int64  `gorm:""`
}
