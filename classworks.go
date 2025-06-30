package main

import "log"

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

func main() {
	example := []float64{1, 2, 2, 4, 5, 2}

	mostNumber, mostCount := MostOccurredNumber(example)

	log.Println("number.......................... :", mostNumber)
	log.Println("mostCount.......................... :", mostCount)

	// answer:
	//2025/06/30 10:51:41 number.......................... : 2
	//2025/06/30 10:51:41 mostCount.......................... : 3

}
