package main

import (
	"fmt"
	"github.com/mvellucci100/MSDS431Week9_Repo1/trimmedmean"
)

func main() {
	// Example data
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
	// Trimmed mean with 10% trimming on both sides
	result, err := trimmedmean.TrimmedMean(data, 0.1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Trimmed Mean: %v\n", result)
	}
}
