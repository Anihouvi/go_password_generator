package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
	"regexp"
)


func generatePassword(length int) string {
	var password strings.Builder
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < length; i++ {
		randomChar := rune(rand.Intn(126-33) + 33) // Generate a random character between 33 and 126
		password.WriteRune(randomChar)
	}

	return password.String()
}


func hashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

func isValidPassword(password string) bool {
	if len(password) < 10 {
		return false
	}

	hasSpecial := false
	for _, r := range password {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			hasSpecial = true
			break
		}
	}

	return hasSpecial
}

func main() {
	var length int
	var err error

	for {
		color.Set(color.FgHiCyan)
		fmt.Print("➡️ Enter the desired password length: ")
		color.Unset()

		lengthInput := ""
		fmt.Scanln(&lengthInput)

		// Check if the input contains only digits
		isDigits := regexp.MustCompile(`^\d+$`).MatchString(lengthInput)
		if !isDigits {
			color.Set(color.FgHiRed)
			fmt.Println("❌ Please enter a valid number.")
			color.Unset()
			continue
		}

		length, err = strconv.Atoi(lengthInput)

		if err != nil || length < 10 {
			color.Set(color.FgHiRed)
			fmt.Println("❌ Password length should be at least 10.")
			color.Unset()
			continue
		} else {
			break
		}
	}

	color.Set(color.FgHiYellow)
	fmt.Println("🚏 Using default algorithm: SHA-256")
	color.Unset()

	password := generatePassword(length)

	if !isValidPassword(password) {
		color.Set(color.FgHiRed)
		fmt.Println("❌ Password does not meet the requirements.")
		color.Unset()
		os.Exit(1)
	}

	hashedPassword := hashPassword(password)

	color.Set(color.FgHiGreen)
	fmt.Println("------------------------------------------------------------")
	fmt.Printf("%-20s|%-30s|%-30s\n", "Length", "Algorithm", "Hash")
	fmt.Println("------------------------------------------------------------")
	fmt.Printf("%-20d|%-30s|", length, "SHA-256")
	for i := 0; i < len(hashedPassword); i += 16 {
		hashPart := hashedPassword[i:Min(i+16, len(hashedPassword))]
		if i == 0 {
			fmt.Printf("%-30s\n", hashPart)
		} else {
			fmt.Printf("%-20s|%-30s|%-30s\n", "", "", hashPart)
		}
	}
	fmt.Println("------------------------------------------------------------")
	fmt.Printf("%-20s\n",                "✅🔐Password Generated")
	fmt.Println("------------------------------------------------------------")
	fmt.Printf("%-20s\n", password)
	fmt.Println("------------------------------------------------------------")
	color.Unset()
}


func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
