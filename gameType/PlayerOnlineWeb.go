package gameType

type PlayerOnlineWeb struct {
	PlayerId int64
	Taps     int64
	Shares   int64
	Energy   int64
	Earned   int64
	Reward   int64
	Spent    int64

	LastTapPlayer int64

	LevelTap     int64
	LevelPlayer  int64
	LevelEnergy  int64
	LevelCharge  int64
	LevelLeagues int64

	LastOnlineWeb int64
}
