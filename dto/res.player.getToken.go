package dto

type ResGetToken struct {
	Total      int64  `json:"total"`
	TokenCoins int64  `json:"token_coins"`
	TokenRefs  int64  `json:"token_refs"`
	TokenLig   int64  `json:"token_Ligue"`
	TokenLevel int64  `json:"token_level"`
	IsFee      bool   `json:"is_fee"`
	Command    string `json:"command"`
	TokenUSD   int64  `json:"token_usd"`
}
