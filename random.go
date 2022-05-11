package iridium

import (
	"errors"
	"math"
	"math/rand"
	"time"
	"unsafe"
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

// RandomString return a random string with given length that only contains lower/upper letters
func RandomString(length int) string {
	var src = rand.NewSource(time.Now().UnixNano())
	b := make([]byte, length)
	for i, cache, remain := length-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}

// RandomInt return a pseudorandom int between min and max
func RandomInt(min, max int) (n int, err error) {
	if min > max {
		return 0, errors.New("wrong min/max sequence")
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Int()%(max-min+1) + min, nil
}

// RandomFloat return a random float with given precision, the integer of float will be 0.
func RandomFloat(precision int) (f float64, err error) {
	n, err := RandomInt(int(math.Pow10(precision-1)), int(math.Pow10(precision)-1))
	if err != nil {
		return 0, err
	}
	return float64(n) / math.Pow10(precision), nil
}
