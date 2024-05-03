package db

type Settings struct {
	Id    int64  `json:"id"`
	Type  string `json:"type" gorm:"unique"`
	Value string `json:"value"`
}
