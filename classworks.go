package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func main() {

	// Question 1 and 2
	reservedString, isPalindrome := ReverseAString("example")

	log.Println(fmt.Sprintf("reversed-word: %s \nisPalindrome: %v", reservedString, isPalindrome))

	// Question 3a. built-in : searches for our available fruits (pass in singular fruits, eg. orange, mango, etc)
	isAvailable, duration := inbuiltSearch("Watermelon")
	log.Println(fmt.Sprintf("in-built search\nisAvailable: %v \nduration: %v", isAvailable, duration))

	// Question 3a. custom
	isAvailable, duration = customSearch([]string{"orange", "apple", "guava", "mango", "watermelon"}, "WatErmelon")
	log.Println(fmt.Sprintf("custom search\nisAvailable: %v \nduration: %v", isAvailable, duration))

}

func MostOccurredNumber(numbs []float64) (mostNumber float64, mostCount int) {
	holder := make(map[float64]int)
	for _, num := range numbs {
		holder[num]++
	}

	for n, count := range holder {
		if count > mostCount {
			mostCount = count
			mostNumber = n
		}
	}

	return
}

//1. Reverse a string using a loop
//2. Check for palindrome using string reversal

func ReverseAString(word string) (reservedString string, isPalindrome bool) {
	for i := len(word) - 1; i >= 0; i-- {
		reservedString = reservedString + string(word[i])
	}

	if word == reservedString {
		isPalindrome = true
	}

	return
}

//3. Write an algorithm that searches for a particular item in an array
//using any built-in method, and a custom method.
//Then measure the time it takes for both implementations to execute.

// inbuiltSearch is an in-built function that searches if we have the fruit in an array. Fruit passed into the search is in singular form.
func inbuiltSearch(fruit string) (available bool, duration time.Duration) {
	start := time.Now()
	if strings.TrimSpace(strings.ToLower(fruit)) == "" {
		duration = time.Since(start)
		return false, duration
	}
	fruits := []string{"orange", "apple", "guava", "mango", "watermelon"}

	for _, f := range fruits {
		if strings.TrimSpace(strings.ToLower(fruit)) == f {
			duration = time.Since(start)
			return true, duration
		}
	}

	duration = time.Since(start)
	return false, duration
}

// customSearch allows you to pass in an array of strings and a word you which to search for in the array
func customSearch(wordsArray []string, searchWord string) (available bool, duration time.Duration) {
	start := time.Now()
	if strings.TrimSpace(strings.ToLower(searchWord)) == "" || len(wordsArray) == 0 {
		duration = time.Since(start)
		return false, duration
	}
	wordStorage := make(map[string]bool)
	for _, word := range wordsArray {
		wordStorage[strings.TrimSpace(strings.ToLower(word))] = true
	}

	return wordStorage[strings.TrimSpace(strings.ToLower(searchWord))], time.Since(start)
}
