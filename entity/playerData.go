package entity

type PlayerData struct {
	Id          int64 `json:"id"`
	PlayerId    int64 `json:"player_id" gorm:"primaryKey"`
	TapLevel    int64 `json:"tap_level"`
	EnergyLevel int64 `json:"energy_level"`
	ChargeLevel int64 `json:"charge_level"`
	Leagues     int64 `json:"leagues"`
	PlayerLevel int64 `json:"player_level"`
	TabBot      bool  `json:"tab_bot"`

	TapLevelDB      TapLevelsDB     `gorm:"foreignKey:Id;references:TapLevel"`
	AccessToken     AccessToken     `gorm:"foreignKey:PlayerId;references:PlayerId"`
	LastTapPlayer   LastTapPlayer   `gorm:"foreignKey:PlayerId;references:PlayerId"`
	OnlineWebPlayer OnlineWebPlayer `gorm:"foreignKey:PlayerId;references:PlayerId"`

	StatsData StatsData `gorm:"foreignKey:PlayerId;references:PlayerId"`

	EnergyLevelDB  EnergyLevelsDB `gorm:"foreignKey:Id;references:EnergyLevel"`
	ChargeLevelsDB ChargeLevelsDB `gorm:"foreignKey:Id;references:ChargeLevel"`
	LeaguesDB      LeaguesDB      `gorm:"foreignKey:Id;references:Leagues"`
	PlayerLevelDB  PlayerLevelsDB `gorm:"foreignKey:Id;references:PlayerLevel"`
	Player         Player         `gorm:"foreignKey:PlayerId;references:PlayerId"`

	ReferralPlayer   []ReferralPlayerData `gorm:"foreignKey:PlayerDataId;references:PlayerId"`
	MissionsComplete []MissionsComplete   `gorm:"foreignKey:PlayerId;references:PlayerId"`
}
