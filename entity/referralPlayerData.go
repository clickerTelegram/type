package entity

type ReferralPlayerData struct {
	Id           int64 `json:"id" gorm:"primaryKey"`
	PlayerId     int64 `json:"player_id" gorm:"primaryKey"`
	PlayerDataId int64 `json:"player_data_id"`
}
