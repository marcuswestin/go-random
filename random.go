package random

import (
	cryptoRand "crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"math"
	"math/rand"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Returns a random int in the range [min, lessThan)
func Between(min, lessThan int) int {
	return rand.Intn(lessThan-min) + min
}

// Returns a random number with numDigits digits
func Digits(numDigits int) string {
	min := int(math.Pow10(numDigits - 1))
	lessThan := int(math.Pow10(numDigits))
	return strconv.Itoa(Between(min, lessThan))
}

// UID returns a base64-encoded random string. numChars must be a multiple of 4.
func UID(numChars int) (uid string, err error) {
	if numChars%4 != 0 {
		err = errors.New("UID length must be a multiple of 4")
		return
	}
	buf := make([]byte, numChars)
	_, err = io.ReadFull(cryptoRand.Reader, buf)
	if err != nil {
		return
	}

	uid = base64.URLEncoding.EncodeToString(buf)
	return
}
