package entity

type TokenClaim struct {
	Id       int64  `gorm:"primaryKey"`
	PlayerId int64  `json:"player_id" gorm:"primaryKey;unique"`
	Uuid     string `json:"uuid" gorm:"primaryKey;unique"`
	Coin     int64
	Leagues  int64
	Level    int64
	Ref      int64
	Total    int64
	IsFee    bool
	Address  string
	CreateAt int64 `gorm:"autoCreateTime"`
}
