package entity

type MissionsComplete struct {
	Id             int64 `gorm:"primaryKey"`
	PlayerId       int64
	Verified       bool
	VerifiedAt     int64
	MissionsId     int64
	MissionsItemId int64

	MissionsDB     MissionsDB      `gorm:"foreignKey:Id;references:MissionsId"`
	MissionsItemDB MissionsItemsDB `gorm:"foreignKey:Id;references:MissionsItemId"`
}
