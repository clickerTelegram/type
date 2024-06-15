package entity

type BoostDataDb struct {
	Id       string `gorm:"primaryKey"`
	PlayerId int64
	Type     BoostDBType
	BoostId  int64
	Count    int64
	Boosts   Boosts `gorm:"foreignKey:Id;references:BoostId"`
	CreateAt int64  `gorm:"autoCreateTime"`
}
