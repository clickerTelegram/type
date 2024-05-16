package entity

type ClaimsData struct {
	Id                int64
	PlayerId          int64
	Active            bool
	LeaguesDB         int64 `gorm:"null"`
	ReferralRewardsDB int64 `gorm:"null"`
	ChargeLevelsDB    int64 `gorm:"null"`
	TapLevelsDB       int64 `gorm:"null"`
	PlayerLevelsDB    int64 `gorm:"null"`
}
