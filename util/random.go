package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	// commented due to rand.Seed() calls deprecation in go1.20
	// rand.Seed(time.Now().UnixNano())
	seed := time.Now().UnixNano()
	rand.New(rand.NewSource(seed))
}

// RandomInt generates a random integer between min and max
func RandomInt(min int64, max int64) int64 {
	return min + rand.Int63n(max-min+1) // min->max
}

// RandomInt32 generates a random integer between min and max with 32-bit integer type
func RandomInt32(min int32, max int32) int32 {
	return min + rand.Int31n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomName generates random name
func RandomName() string {
	return RandomString(6)
}

// RandomBool returned random boolean (true/false) value
func RandomBool() bool {
	return rand.Intn(2) == 0
}
