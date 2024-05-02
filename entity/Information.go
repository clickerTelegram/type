package entity

type Information struct {
	Id        string `json:"id"`
	UserId    string `json:"user_id" gorm:"unique"`
	BalanceId int
	LevelId   int
}
