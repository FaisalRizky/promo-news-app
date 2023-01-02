package util

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

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
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

func RandomOwner() string {
	return RandomString(6)
}

func RandomBoolean() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) == 1
}

func RandomUnixTime() int64 {
	randomTime := rand.Int63n(time.Now().Unix()-94608000) + 94608000
	randomNow := time.Unix(randomTime, 0)
	return randomNow.Unix()
}

func RandomTime() time.Time {
	randomTime := rand.Int63n(time.Now().Unix()-94608000) + 94608000
	return time.Unix(randomTime, 0)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}
	n := len(currencies)
	return currencies[(rand.Intn(n))]
}

func RandomDay() string {
	day := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	n := len(day)
	return day[(rand.Intn(n))]
}

func RandomAmHours() string {
	hours := []string{"00:00", "01:00", "02:00", "03:00", "05:00", "06:00", "07:00", "08:00", "09:00", "10:00", "11:00"}
	n := len(hours)
	return hours[(rand.Intn(n))]
}

func RandomPmHours() string {
	hours := []string{"12:00", "13:00", "14:00", "15:00", "16:00", "17:00", "18:00", "19:00", "20:00", "21:00", "22:00", "23:00"}
	n := len(hours)
	return hours[(rand.Intn(n))]
}

// RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
