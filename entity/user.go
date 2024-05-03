package entity

type User struct {
	Id            int64       `json:"id" gorm:"primary_key"`
	Uuid          string      `json:"uuid" gorm:"unique"`
	UserId        int64       `json:"user_id" gorm:"unique"`
	InformationId int64       `json:"information_id" gorm:"unique"`
	WalletAddress string      `json:"wallet_address" gorm:"unique"`
	Referrals     []Referrals `gorm:"foreignKey:CallerId;references:UserId"`
	Information   Information `gorm:"foreignKey:UserId;references:InformationId"`
}

func (u *User) TableName() string {
	return "users"
}
