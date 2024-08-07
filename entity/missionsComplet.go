package entity

type MissionsComplete struct {
	Id             string `gorm:"primaryKey"`
	PlayerId       int64
	Verified       bool
	VerifiedAt     int64
	MissionsId     int64
	MissionsItemId int64

	MissionsDB     MissionsDB      `gorm:"foreignKey:Id;references:MissionsId"`
	MissionsItemDB MissionsItemsDB `gorm:"foreignKey:Id;references:MissionsItemId"`
}

type MissionCompleteData struct {
	Id         string `gorm:"primaryKey"`
	PlayerId   int64
	MissionsId int64
	Verified   bool
}
