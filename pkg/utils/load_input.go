package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func LoadData(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scannerResults := make([]string, 0)

	for scanner.Scan() {
		scannerResults = append(scannerResults, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return scannerResults
}

func ParseInt(s string) (parsedInteger int) {
	parsedInteger, _ = strconv.Atoi(s)
	return
}
