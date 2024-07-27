package dto

type SubmitTap struct {
	Player Player `json:"player"`
}

type ResPlayerTask struct {
	Player  Player  `json:"player"`
	Account Account `json:"account"`
}
