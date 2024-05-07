package entity

type UserDb struct {
	ID            uint   `gorm:"primaryKey"`
	UserId        uint64 `gorm:"unique"`
	ReferralCode  string `gorm:"unique"`
	AddressWallet string `gorm:"unique"`
}
