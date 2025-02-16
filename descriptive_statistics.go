package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Desciptive Statistic - Count the most words in founded in every cell defined
func countWordFrequencies(records [][]string, columnIndex int) map[string]int {
	wordCounts := make(map[string]int)

	for _, record := range records {
		if len(record) > columnIndex {
			word := strings.ToLower(strings.TrimSpace(record[columnIndex]))
			wordCounts[word]++
		}
	}

	return wordCounts
}
func findMostFrequentWord(wordCounts map[string]int) (string, int) {
	var mostFrequentWord string
	maxCount := 0

	for word, count := range wordCounts {
		if count > maxCount {
			mostFrequentWord = word
			maxCount = count
		}
	}

	return mostFrequentWord, maxCount
}

// Descriptive Statistic - Find Mean, Modus, Median
func findMean(values []float64) float64 {
	var sum float64
	for _, value := range values {
		sum += value
	}
	return sum / float64(len(values))
}

func findMode(values []float64) []float64 {
	frequencyMap := make(map[float64]int)
	var maxCount int

	for _, value := range values {
		frequencyMap[value]++
		if frequencyMap[value] > maxCount {
			maxCount = frequencyMap[value]
		}
	}

	var modes []float64
	for value, count := range frequencyMap {
		if count == maxCount {
			modes = append(modes, value)
		}
	}

	return modes
}
func findMedian(values []float64) float64 {
	sort.Float64s(values)
	n := len(values)
	if n%2 == 0 {
		return (values[n/2-1] + values[n/2]) / 2
	}
	return values[n/2]
}

func main() {
	// Read CSV
	file, err := os.Open("Messy_HR_Dataset_Detailed.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Desciptive Statistic - Count the most words in founded in every cell defined
	headers := records[0]
	target_col := []string{"Title", "Supervisor", "BusinessUnit", "EmployeeStatus", "EmployeeType", "PayZone", "EmployeeClassificationType", "DepartmentType",
		"TerminationType", "Division", "State", "GenderCode", "JobFunctionDescription", "RaceDesc", "MaritalDesc", "Performance Score",
		"Training Program Name", "Training Type", "Training Outcome", "Location", "Trainer"}

	for _, colName := range target_col {
		columnIndex := -1
		for i, header := range headers {
			if header == colName {
				columnIndex = i
				break
			}
		}

		if columnIndex == -1 {
			fmt.Printf("Column '%s' not found in the CSV file\n", colName)
			continue
		}

		wordCounts := countWordFrequencies(records[1:], columnIndex)
		mostFrequentWord, maxCount := findMostFrequentWord(wordCounts)

		fmt.Printf("For column '%s', most frequent word: '%s' with %d occurrences\n", colName, mostFrequentWord, maxCount)
	}

	// Descriptive Statistic - Find Mean, Modus, Median
	target_col_2 := []string{"Current Employee Rating", "Engagement Score", "Satisfaction Score", "Work-Life Balance Score", "Training Duration(Days)"}
	for _, colName := range target_col_2 {
		columnIndex := -1
		for i, header := range headers {
			if header == colName {
				columnIndex = i
				break
			}
		}

		if columnIndex == -1 {
			fmt.Printf("Column '%s' not found in the CSV file\n", colName)
			continue
		}

		var columnValues []float64
		for _, record := range records[1:] {
			if len(record) > columnIndex {
				value := record[columnIndex]
				if v, err := strconv.ParseFloat(value, 64); err == nil {
					columnValues = append(columnValues, v)
				}
			}
		}

		fmt.Printf("\nFor column '%s'\n", colName)
		mean := findMean(columnValues)
		fmt.Printf("Mean: %.2f\n", mean)
		mode := findMode(columnValues)
		fmt.Printf("Mode: %v\n", mode)
		median := findMedian(columnValues)
		fmt.Printf("Median: %.2f\n", median)
	}
}
