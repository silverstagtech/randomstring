// Package randomstring is used to generate random strings that could be used in different applications.
// Strings generated can contain digits, special character, upper and lowercase letters.
// This package only focuses on English letters.
package randomstring

import (
	"errors"
	"math/rand"
	"time"
)

const (
	digits      = "0123456789"
	specials    = "~=+%^*/()[]{}/!@#$?|"
	letterUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterLower = "abcdefghijklmnopqrstuvwxyz"
	all         = digits + specials + letterUpper + letterLower
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Generate creates a random string and returns it to the caller.
// The string will contain at least the number of character types defined on the function call.
// If the combined length of the required characters is larger than the length required an error is returned.
func Generate(digit, special, upper, lower, length int) (string, error) {
	if (digit + special + upper + lower) > length {
		return "", errors.New("requirement's are longer than length")
	}

	buf := make([]byte, length)
	index := 0

	// Furfill the required characters
	for i := 0; i < digit; i++ {
		buf[index] = digits[rand.Intn(len(digits))]
		index++
	}
	for i := 0; i < special; i++ {
		buf[index] = specials[rand.Intn(len(specials))]
		index++
	}
	for i := 0; i < upper; i++ {
		buf[index] = letterUpper[rand.Intn(len(letterUpper))]
		index++
	}
	for i := 0; i < lower; i++ {
		buf[index] = letterLower[rand.Intn(len(letterLower))]
		index++
	}

	// Do the rest
	for i := index; i < length; i++ {
		buf[i] = all[rand.Intn(len(all))]
	}

	// shuffle characters
	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})

	return string(buf), nil
}
