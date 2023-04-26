package util

import (
	"math/rand"
	"time"
)

const (
	charset        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	referralLength = 8
)

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateReferral() string {
	referral := make([]byte, referralLength)
	for r := range referral {
		referral[r] = charset[seededRand.Intn(len(charset))]
	}

	return string(referral)
}
