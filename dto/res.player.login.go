package dto

import "time"

type LoginRes struct {
	AccessToken string   `json:"access_token"`
	Player      Player   `json:"player"`
	Account     Account  `json:"account"`
	BotShares   int64    `json:"bot_shares"`
	Conf        Conf     `json:"conf"`
	Settings    Settings `json:"settings"`
}
type Boost struct {
	Type string `json:"type"`
	Cnt  int64  `json:"cnt"`
	End  int64  `json:"end"`
}
type Stat struct {
	Taps   int64 `json:"taps"`
	RefIn  int64 `json:"ref_in"`
	RefOut int64 `json:"ref_out"`
	RefCnt int64 `json:"ref_cnt"`
	Earned int64 `json:"earned"`
	Reward int64 `json:"reward"`
	Spent  int64 `json:"spent"`
}
type Player struct {
	ID            int64         `json:"id"`
	Name          string        `json:"name"`
	FullName      string        `json:"full_name"`
	WalletAddress string        `json:"wallet_address"`
	LoginTs       int64         `json:"login_ts"`
	Time          int64         `json:"time"`
	Energy        int64         `json:"energy"`
	Shares        int64         `json:"shares"`
	Tokens        int64         `json:"tokens"`
	Ligue         int64         `json:"ligue"`
	EnergyLevel   int64         `json:"energy_level"`
	ChargeLevel   int64         `json:"charge_level"`
	PlayerLevel   int64         `json:"player_level"`
	TapLevel      int64         `json:"tap_level"`
	TapBot        bool          `json:"tap_bot"`
	Boost         []Boost       `json:"boost"`
	BoostTime     int64         `json:"boost_time"`
	Claims        map[int64]any `json:"claims"`
	Stat          Stat          `json:"stat"`
	IsBot         bool          `json:"is_bot"`
}
type ItemsActive struct {
	Type       string `json:"type"`
	Verified   bool   `json:"verified"`
	VerifiedAt int64  `json:"verified_at"`
}
type Active struct {
	ID    string                `json:"id"`
	Items map[int64]ItemsActive `json:"items"`
}
type MissionsAccount struct {
	Completed []any    `json:"completed"`
	Active    []Active `json:"active"`
}
type Account struct {
	ID       int64           `json:"id"`
	Missions MissionsAccount `json:"missions"`
}
type EnergyLevels struct {
	Limit int64 `json:"limit"`
	Price int64 `json:"price"`
}
type ChargeLevels struct {
	Rate  int64 `json:"rate"`
	Price int64 `json:"price"`
}

type PlayerLevels struct {
	Score  int64 `json:"score"`
	Reward int64 `json:"reward"`
}

type TapLevels struct {
	Rate   int64 `json:"rate"`
	Energy int64 `json:"energy"`
	Price  int64 `json:"price"`
}
type Ligues struct {
	Name      string `json:"name"`
	Title     string `json:"title"`
	Score     int64  `json:"score"`
	Reward    int64  `json:"reward"`
	RewardRef int64  `json:"reward_ref"`
}
type RefRewards struct {
	Cnt    int64 `json:"cnt"`
	Reward int64 `json:"reward"`
}
type Boots struct {
	Duration int64 `json:"duration"`
	RateMult int64 `json:"rateMult"`
	Energy   int64 `json:"energy"`
}

type TapBot struct {
	Duration int64 `json:"duration"`
	Price    int64 `json:"price"`
}

type ItemsMissions struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Link string `json:"link,omitempty"`
}
type MissionsConf struct {
	ID          string                  `json:"id"`
	Title       string                  `json:"title"`
	Description string                  `json:"description"`
	Reward      int64                   `json:"reward"`
	Items       map[int64]ItemsMissions `json:"items"`
}
type Conf struct {
	EnergyLevels map[int64]EnergyLevels `json:"energy_levels"`
	ChargeLevels map[int64]ChargeLevels `json:"charge_levels"`
	PlayerLevels map[int64]PlayerLevels `json:"player_levels"`
	TapLevels    map[int64]TapLevels    `json:"tap_levels"`
	Ligues       map[int64]Ligues       `json:"ligues"`
	RefRewards   map[int64]RefRewards   `json:"ref_rewards"`
	Boosts       map[string]Boots       `json:"boosts"`
	TapBot       TapBot                 `json:"tap_bot"`
	Missions     map[int64]MissionsConf `json:"missions"`
}
type Settings struct {
	StartDate       time.Time `json:"start_date"`
	SubmitIntervalS int64     `json:"submit_interval_s"`
}
