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
		log.Fatal("File not found...error", err)
	}

	str := string(bytearray)
	lines := strings.Split(str, "")
	wordCount := countWords(lines)

	reportResults(wordCount)
}

func reportResults(count map[string]int) {

}

func countWords(lines []string) map[string]int {
	wordCount := make(map[string]int)

}
