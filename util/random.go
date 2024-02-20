package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdejklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

/**
Generate random number 
 */
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

/*
Generate random string of length n
*/
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomUsername() string {
	return RandomString(6);
}

// return random Rule from rules 
func RandomRule() string {
	rules := []string{"admin", "user", "ghost"}
	n := len(rules)
	return rules[rand.Intn(n)]
}
