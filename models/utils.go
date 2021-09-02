package models

import (
	"math/rand"
	"time"
)

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandomString(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RandomFloat64(min, max float64) float64 {
	return seededRand.Float64() * (max - min)
}

func RandomInt(min, max int) int {
	return seededRand.Intn(max) + min
}

func RandomDate(min, max int64) time.Time {
	delta := max - min

	sec := rand.Int63n(delta) + min
	t := time.Unix(sec, 0)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}
