package entity

type MissionsDB struct {
	Id          int64             `gorm:"primaryKey"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Reward      int64             `json:"reward"`
	Items       []MissionsItemsDB `gorm:"foreignKey:MissionsId;references:Id"`
}

type MissionsItemsDB struct {
	Id         int64  `gorm:"primaryKey"`
	MissionsId int64  `json:"missionsId"`
	Type       string `json:"type"`
	Name       string `json:"name"`
	Link       string `json:"link"`
}
