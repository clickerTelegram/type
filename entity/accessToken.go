package entity

type AccessToken struct {
	PlayerId int64  `gorm:"primaryKey"`
	Access   string `gorm:"unique"`
}
