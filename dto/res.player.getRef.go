package dto

type GetRef struct {
	Refs RefList `json:"refs"`
}

type RefList struct {
	LeaguesName string `json:"leagues_name"`
	ID          int64
	Name        string
	FullName    string
	Earned      int64
	Rewards     int64
}
