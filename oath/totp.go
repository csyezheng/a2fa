package oath

import (
	"encoding/binary"
	"time"
)

type TOTP struct {
	base32      bool
	hashMethod  string
	valueLength int
	epoch       int64
	interval    int64
}

func NewTOTP(base32 bool, hash string, length int, epoch int64, interval int64) *TOTP {
	return &TOTP{
		base32:      base32,
		hashMethod:  hash,
		valueLength: length,
		epoch:       epoch,
		interval:    interval,
	}
}

// counter return the count of the number of durations interval between epoch and currentTime
func (t *TOTP) counter() int64 {
	// currentTime is the current time in seconds since a particular epoch,
	currentTime := time.Now().UTC().Unix()
	delta := currentTime - t.epoch
	return delta / t.interval
}

func (t *TOTP) GeneratePassCode(secretKey string) string {
	hotp := NewHOTP(t.base32, t.hashMethod, t.counter(), t.valueLength)
	return hotp.GeneratePassCode(secretKey)
}

func counterToBytes(counter int64) []byte {
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, uint64(counter))
	return bytes
}
