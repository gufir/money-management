package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(length int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < length; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomPeriod() string {
	rand.Seed(time.Now().UnixNano())

	year := rand.Intn(31) + 2000

	month := rand.Intn(12) + 1

	return fmt.Sprintf("%d-%02d", year, month)
}

func RandomName() string {
	return RandomString(6)
}

func RandomEmail() string {
	return RandomString(6) + "@email.com"
}

func RandomUUID() string {
	id, err := uuid.NewRandom()
	if err != nil {
		fmt.Print(err)
	}

	return id.String()
}
