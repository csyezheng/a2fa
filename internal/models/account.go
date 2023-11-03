package models

import (
	"github.com/csyezheng/a2fa/oath"
	"log/slog"
)

type Account struct {
	Name        string `json:"name" gorm:"primaryKey"`
	SecretKey   string `json:"secretkey" binding:"required"`
	Mode        string `json:"mode"`
	Base32      bool   `json:"base32"`
	Hash        string `json:"hash"`
	ValueLength int    `json:"length"`
	Counter     int64  `json:"counter"`
	Epoch       int64  `json:"epoch"`
	Interval    int64  `json:"interval"`
}

// otp generate one-time password code
func (a Account) OTP() string {
	code := ""
	if a.Mode == "hotp" {
		hotp := oath.NewHOTP(a.Base32, a.Hash, a.Counter, a.ValueLength)
		code = hotp.GeneratePassCode(a.SecretKey)
	} else if a.Mode == "totp" {
		totp := oath.NewTOTP(a.Base32, a.Hash, a.ValueLength, a.Epoch, a.Interval)
		code = totp.GeneratePassCode(a.SecretKey)
	} else {
		slog.Error("mode should be hotp or totp")
	}
	return code
}
