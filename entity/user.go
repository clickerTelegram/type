package entity

type User struct {
	Id            int64  `json:"id" gorm:"primary_key"`
	Uuid          string `json:"uuid" gorm:"unique"`
	UserId        int64  `json:"user_id" gorm:"unique"`
	WalletAddress string `json:"wallet_address"`
	InformationId
	ReferralsIds []Referrals
}

func (u *User) TableName() string {
	return "users"
}
