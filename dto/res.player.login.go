package dto

import "time"

type LoginRes struct {
	AccessToken string   `json:"access_token"`
	Player      Player   `json:"player"`
	Account     Account  `json:"account"`
	BotShares   int      `json:"bot_shares"`
	Conf        Conf     `json:"conf"`
	Settings    Settings `json:"settings"`
}
type Boost struct {
	Type string `json:"type"`
	Cnt  int    `json:"cnt"`
	End  int    `json:"end"`
}
type Stat struct {
	Taps   int `json:"taps"`
	RefIn  int `json:"ref_in"`
	RefOut int `json:"ref_out"`
	RefCnt int `json:"ref_cnt"`
	Earned int `json:"earned"`
	Reward int `json:"reward"`
	Spent  int `json:"spent"`
}
type Player struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	FullName    string  `json:"full_name"`
	LoginTs     int64   `json:"login_ts"`
	Time        int64   `json:"time"`
	Energy      int     `json:"energy"`
	Shares      int     `json:"shares"`
	Tokens      int     `json:"tokens"`
	Ligue       int     `json:"ligue"`
	EnergyLevel int     `json:"energy_level"`
	ChargeLevel int     `json:"charge_level"`
	TapLevel    int     `json:"tap_level"`
	TapBot      bool    `json:"tap_bot"`
	Boost       []Boost `json:"boost"`
	BoostTime   int64   `json:"boost_time"`
	Claims      []any   `json:"claims"`
	Stat        Stat    `json:"stat"`
}
type ItemsActive struct {
	Type       string `json:"type"`
	Verified   bool   `json:"verified"`
	VerifiedAt int64  `json:"verified_at"`
}
type Active struct {
	ID    string        `json:"id"`
	Items []ItemsActive `json:"items"`
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
	Limit int `json:"limit"`
	Price int `json:"price"`
}
type ChargeLevels struct {
	Rate  int `json:"rate"`
	Price int `json:"price"`
}
type TapLevels struct {
	Rate   int `json:"rate"`
	Energy int `json:"energy"`
	Price  int `json:"price"`
}
type Ligues struct {
	Name      string `json:"name"`
	Title     string `json:"title"`
	Score     int    `json:"score"`
	Reward    int    `json:"reward"`
	RewardRef int    `json:"reward_ref"`
}
type RefRewards struct {
	Cnt    int `json:"cnt"`
	Reward int `json:"reward"`
}
type Energy struct {
	Duration int `json:"duration"`
}
type Turbo struct {
	Duration int `json:"duration"`
	RateMult int `json:"rateMult"`
}
type Double struct {
	Duration int `json:"duration"`
	RateMult int `json:"rateMult"`
	Energy   int `json:"energy"`
}
type Boosts struct {
	Energy Energy `json:"energy"`
	Turbo  Turbo  `json:"turbo"`
	Double Double `json:"double"`
}
type TapBot struct {
	Duration int `json:"duration"`
	Price    int `json:"price"`
}

type ItemsMissions struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Link string `json:"link,omitempty"`
}
type MissionsConf struct {
	ID          string          `json:"id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Reward      int             `json:"reward"`
	Items       []ItemsMissions `json:"items"`
}
type Conf struct {
	EnergyLevels []EnergyLevels `json:"energy_levels"`
	ChargeLevels []ChargeLevels `json:"charge_levels"`
	TapLevels    []TapLevels    `json:"tap_levels"`
	Ligues       []Ligues       `json:"ligues"`
	RefRewards   []RefRewards   `json:"ref_rewards"`
	Boosts       Boosts         `json:"boosts"`
	TapBot       TapBot         `json:"tap_bot"`
	Missions     []MissionsConf `json:"missions"`
}
type Settings struct {
	StartDate       time.Time `json:"start_date"`
	SubmitIntervalS int       `json:"submit_interval_s"`
}
