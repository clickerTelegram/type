package gameType

import "math/big"

type GameStatus struct {
	TotalShareBalance *big.Int
	TotalTouches      *big.Int
	TotalPlayers      int64
	DailyUsers        int64
	OnlinePlayers     int64
}
