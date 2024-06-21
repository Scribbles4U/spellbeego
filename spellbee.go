package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	// Open our words file
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	// Initiate variables
	words := make([]string, 0)
	var uniq string
	var let string
	s := bufio.NewScanner(f)

	// Scan and add words to our slice of strings
	for s.Scan() {
		words = append(words, s.Text())
	}
	if err := s.Err(); err != nil {
		log.Println(err)
		return
	}

	// Obtain our letters from user
	fmt.Print("Enter the center/required letter: ")
	fmt.Scanln(&uniq)
	fmt.Print("Enter the other letters: ")
	fmt.Scanln(&let)

	// Ensure the letters are lowercase
	uniq = strings.ToLower(uniq)
	let = strings.ToLower(let)

	// Remove our letters from alphabet for easier deduction
	alphabet := replaceChars(let, uniq)

	// Search and compile list of anagrams
	anagrams := wordSearch(words, uniq, alphabet)

	fmt.Println("Number of anagrams found: ", len(anagrams))
	fmt.Println(anagrams)
}

func wordSearch(words []string, uniq string, alphabet string) (anagrams []string) {
	pattern := "[" + alphabet + "]+"
	rNot, _ := regexp.Compile(pattern)

	for _, w := range words {
		if strings.Contains(w, uniq) && len(w) >= 4 {
			if !(rNot.MatchString(w)) {
				anagrams = append(anagrams, w)
			}
		}
	}
	return anagrams
}

func replaceChars(letters string, uniq string) string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	substring := uniq + letters
	replaceChar := rune(-1) // Can use negative value to denote blank rune instead of ''

	// Map for quick lookup of characters
	charMap := make(map[rune]bool)
	for _, c := range substring {
		charMap[c] = true
	}

	// Replace our chars using strings.Map
	return strings.Map(func(r rune) rune {
		if charMap[r] {
			return replaceChar
		}
		return r
	}, alphabet)
}
