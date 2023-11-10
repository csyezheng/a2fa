package models

import (
	"fmt"
	"github.com/csyezheng/a2fa/oath"
)

type Account struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	AccountName string `json:"accountName" binding:"required"`
	Username    string `json:"username"`
	SecretKey   string `json:"secretKey" binding:"required"`
	Mode        string `json:"mode"`
	Base32      bool   `json:"base32"`
	Hash        string `json:"hash"`
	ValueLength int    `json:"length"`
	Counter     int64  `json:"counter"`
	Epoch       int64  `json:"epoch"`
	Interval    int64  `json:"interval"`
}

// otp generate one-time password code
func (a Account) OTP() (code string, err error) {
	if a.Mode == "hotp" {
		hotp := oath.NewHOTP(a.Base32, a.Hash, a.Counter, a.ValueLength)
		code, err = hotp.GeneratePassCode(a.SecretKey)
	} else if a.Mode == "totp" {
		totp := oath.NewTOTP(a.Base32, a.Hash, a.ValueLength, a.Epoch, a.Interval)
		code, err = totp.GeneratePassCode(a.SecretKey)
	} else {
		return code, fmt.Errorf("mode should be hotp or totp")
	}
	return
}
