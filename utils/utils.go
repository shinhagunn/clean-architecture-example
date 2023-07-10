package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateUID() string {
	rand.Seed(time.Now().UnixNano())

	result := "UID"

	for i := 0; i < 10; i++ {
		result += fmt.Sprint(rand.Intn(9))
	}

	return result
}
