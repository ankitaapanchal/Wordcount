package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {

	var fileName string
	// Taking input from user
	fmt.Println("Enter the file name: ")
	fmt.Scanln(&fileName)

	//open and read file
	bytearray, err := ioutil.ReadFile(fileName)
	if err != nil {
		// If found error, so it will exit the program.
		log.Fatal("File not found...error", err)
	}

	// convert byte array into string
	str := string(bytearray)
	// splits string into lines
	lines := strings.Split(str, "\n")
	wordCount := countWords(lines)

	reportResults(wordCount)
}

func countWords(lines []string) map[string]int {
	// Initialize an empty map
	wordCount := make(map[string]int)
	// Split into words **
	for _, line := range lines {
		words := strings.Fields(line)

		//Increments each word count in the map **
		for _, word := range words {
			wordCount[word]++
		}
	}

	return wordCount
}

func reportResults(count map[string]int) {
	// Keys are words and values are counts
	for key, value := range count {
		fmt.Printf("%s : %d\n", key, value)
	}
}

// Referred ChatGPT to implement the code marked with ** //
