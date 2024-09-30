package dto

type ResGetToken struct {
	Total      float64 `json:"total"`
	TokenCoins float64 `json:"token_coins"`
	TokenRefs  float64 `json:"token_refs"`
	TokenLig   float64 `json:"token_Ligue"`
	TokenLevel float64 `json:"token_level"`
	IsFee      bool    `json:"is_fee"`
	Command    string  `json:"command"`
	TokenUSD   float64 `json:"token_usd"`
}
