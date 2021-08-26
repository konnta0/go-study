package main

import (
	"math/rand"
	"time"
)

func randomString(n int) string {

	if n <= 0 {
		return ""
	}

	letters := []string{
		"12345",
		"67890",
		"abcde",
		"ABCDE",
	}

	rand.Seed(time.Now().UnixNano())

	letter := letters[rand.Intn(len(letters))]
	if n >= len(letter) {
		return letter
	}

	start := rand.Intn(len(letter) - n + 1)
	return letter[start : start+n]
}
