package entity

type BoostDataDb struct {
	Id       int64 `gorm:"primaryKey"`
	PlayerId int64
	Type     BoostDBType
	BoostId  int64
	Count    int64
	Boosts   Boosts `gorm:"foreignKey:BoostId;references:ID"`
	CreateAt int64  `gorm:"autoCreateTime"`
}
