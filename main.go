package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/abdullahrehmat/go-diceware/dice"
	"github.com/abdullahrehmat/go-diceware/wordlist"
)

//func die(sides int64) int {
//
//	x, err := rand.Int(rand.Reader, big.NewInt(sides))
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	return int(x.Int64()) + 1
//}

func rollDice(x int) []int {

	codes := make([]int, x)

	for i := 0; i < x; i++ {

		var digits string
		var err error

		for z := 0; z < 5; z++ {

			//digits = fmt.Sprintf("%s%d", digits, die(6))
			digits = fmt.Sprintf("%s%d", digits, dice.Die(6))

		}

		codes[i], err = strconv.Atoi(digits)

		if err != nil {
			log.Fatal(err)
		}
	}

	return codes

}

func createPhrase(digits []int) string {

	// Import EFF Diceware Wordlist
	w := wordlist.EffLong()

	// Initialise Passphrase
	words := make([]string, len(digits))

	for i := 0; i < len(digits); i++ {

		// Check Digits Concists Of 5 Numbers
		if digits[i] < 11111 {

			log.Fatal("Incorrect Amount Of Digits")

		}

		words[i] = w[digits[i]]

	}

	phrase := strings.Join(words, "-")

	return phrase

}

func main() {

	// Program Information
	fmt.Println("Diceware | Go Implementation v1.0.0")
	fmt.Println("Wordlist: EFF Diceware (Large)")
	fmt.Println("")

	// Get Number Of Words Required
	var phraseLength int
	fmt.Println("Enter Amount of Phrase Words Required [Min 5]...")
	fmt.Scanln(&phraseLength)
	fmt.Println("")

	// Check Phase Length Is Not Too Short
	if phraseLength < 5 {

		fmt.Println("Too Few Words...")
		os.Exit(1)

	}

	// Roll Dice
	codes := rollDice(phraseLength)

	// Create Phrase
	phrase := createPhrase(codes)
	fmt.Println("Passphrase:")
	fmt.Println("> ", phrase)

}
