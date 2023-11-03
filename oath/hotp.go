package oath

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base32"
	"encoding/binary"
	"encoding/hex"
	"math"
	"strconv"
)

type HOTP struct {
	base32      bool
	hashMethod  string
	counter     int64
	valueLength int
}

func NewHOTP(base32 bool, hash string, counter int64, length int) *HOTP {
	return &HOTP{
		base32:      base32,
		hashMethod:  hash,
		counter:     counter,
		valueLength: length,
	}
}

func (t *HOTP) GeneratePassCode(secretKey string) string {
	var secret []byte
	if t.base32 {
		secret, _ = base32.StdEncoding.DecodeString(secretKey)
	} else {
		secret, _ = hex.DecodeString(secretKey)
	}

	var sum []byte

	switch t.hashMethod {
	case "SHA1":
		mac := hmac.New(sha1.New, secret)
		mac.Write(counterToBytes(t.counter))
		sum = mac.Sum(nil)

	case "SHA256":
		mac := hmac.New(sha256.New, secret)
		mac.Write(counterToBytes(t.counter))
		sum = mac.Sum(nil)

	case "SHA512":
		mac := hmac.New(sha512.New, secret)
		mac.Write(counterToBytes(t.counter))
		sum = mac.Sum(nil)

	default:
		panic("invalid hash algorithm")
	}

	offset := sum[len(sum)-1] & 0xf
	binaryCode := binary.BigEndian.Uint32(sum[offset:])

	codeNum := int64(binaryCode) & 0x7FFFFFFF

	return strconv.FormatInt(codeNum%(int64(math.Pow10(t.valueLength))), 10)
}
