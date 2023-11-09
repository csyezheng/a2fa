package oath

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base32"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
	"strings"
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

func (t *HOTP) GeneratePassCode(secretKey string) (code string, err error) {
	var secret []byte
	if t.base32 {
		// remove spaces and convert to uppercase
		secretKey = strings.Join(strings.Fields(secretKey), "")
		secretKey = strings.ToUpper(secretKey)
		secret, err = base32.StdEncoding.DecodeString(secretKey)
		if err != nil {
			return "", fmt.Errorf("base32 decoding failed: Base32-encoded secret key is invalid")
		}
	} else {
		// remove spaces, hexadecimal is not case-sensitive
		secretKey = strings.Join(strings.Fields(secretKey), "")
		secret, err = hex.DecodeString(secretKey)
		if err != nil {
			return "", fmt.Errorf("hex decoding failed: hex-encoded secret key is invalid")
		}
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
		return "", fmt.Errorf("invalid hash algorithm. valid hash algorithms include values SHA1, SHA256, or SHA512")
	}

	offset := sum[len(sum)-1] & 0xf
	binaryCode := binary.BigEndian.Uint32(sum[offset:])

	codeNum := int64(binaryCode) & 0x7FFFFFFF

	return strconv.FormatInt(codeNum%(int64(math.Pow10(t.valueLength))), 10), err
}
