package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

const (
	letters        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers        = "0123456789"
	specialChars   = "!@#$%^&*()-_=+[]{};:,.<>/?"
	lettersNumbers = letters + numbers
	allChars       = letters + numbers + specialChars
)

// Function to generate a random password
func generatePassword(length int, useNumbers bool, useSpecialChars bool) string {
	charset := letters
	if useNumbers {
		charset += numbers
	}
	if useSpecialChars {
		charset += specialChars
	}

	var password strings.Builder
	for i := 0; i < length; i++ {
		charIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		password.WriteByte(charset[charIndex.Int64()])
	}

	return password.String()
}

func main() {
	var length int
	var useNumbers, useSpecialChars bool

	fmt.Println("Password Generator")
	fmt.Print("Enter desired password length: ")
	fmt.Scanf("%d", &length)

	fmt.Print("Include numbers? (true/false): ")
	fmt.Scanf("%t", &useNumbers)

	fmt.Print("Include special characters? (true/false): ")
	fmt.Scanf("%t", &useSpecialChars)

	password := generatePassword(length, useNumbers, useSpecialChars)
	fmt.Printf("Generated password: %s\n", password)
}
