package random

import (
	"fmt"
	"math/rand"
)


func NewRandomString(length int) string {
	randStr := ""

	for i := 0; i < length; i++ {
		code := rand.Intn(26) + 97 // generate random ASCII code of the letter
		letter := string(rune(code))
		randStr = fmt.Sprintf("%s%s", randStr, letter)
	}

	return randStr
}