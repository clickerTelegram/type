package dto

type SubmitTap struct {
	Player Player `json:"player"`
	IsBot  bool   `json:"is_bot"`
}

type ResPlayerTask struct {
	Player  Player  `json:"player"`
	Account Account `json:"account"`
	IsBot   bool    `json:"is_bot"`
}
