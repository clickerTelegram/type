package entity

type LeagueDb struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"unique"`
	Bones uint64
}
