package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

// Still onprpgress
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
}
