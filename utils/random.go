package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomName() string {
	return RandomString(15)
}

func RandomDescription() string {
	return RandomString(3000)
}
func RandomType() string {
	types := []string{
		"Corporations",
		"Non Profit",
		"Cooperative",
		"Sole Proprietorship",
	}

	n := len(types)
	return types[randomIntRange(0, n-1)]
}

func randomIntRange(min, max int) int {
	return min + rand.Intn(max-min+1)

}

func RandomAmount() int32 {
	return int32(randomIntRange(10, 1000))
}

func RandomID() int64 {
	return int64(randomIntRange(10, 1000))

}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
