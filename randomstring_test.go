package randomstring

import (
	"fmt"
	"testing"
)

func TestGenerate(t *testing.T) {
	length := 16
	s, err := Generate(2, 2, 2, 2, length)
	if err != nil {
		t.Logf("Failed to generate. Error: %s", err)
		t.Fail()
	}
	if len(s) != length {
		t.Logf("Returned string is not long enough. Got: %d, Want: %d", len(s), length)
		t.Fail()
	}
	t.Logf("Got random string: %s", s)
}

func TestGenerateTooLarge(t *testing.T) {
	_, err := Generate(10, 10, 10, 10, 10)
	if err == nil {
		t.Logf("Expected Generate to error because length is too small.")
		t.Fail()
	}
}

func ExampleGenerate() {
	length := 32
	// Skip the error because we have a perfect world in the example.

	// Here we will generate a strong password
	password, _ := Generate(2, 2, 2, 2, length)

	// By using all the characters on the digits you will leave no more room to add anything else
	// the result will be a string of only numbers.
	numbers, _ := Generate(32, 0, 0, 0, length)

	// Again we can do the same thing but only generate special characters
	specials, _ := Generate(0, 32, 0, 0, length)

	fmt.Println(password)
	fmt.Println(numbers)
	fmt.Println(specials)
}
