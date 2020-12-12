package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err == nil {
		return
	}

	panic(err)
}

type match func(value string) bool

type passwordRecord struct {
	character   string
	firstIndex  int
	secondIndex int
	password    string
}

func matchingLines(path string, test match) int {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	matches := 0

	for scanner.Scan() {
		if test(scanner.Text()) {
			matches += 1
		}
	}

	return matches
}

func (pr *passwordRecord) isValid() bool {
	charStrings := strings.Split(pr.password, "")

	firstMatch := charStrings[pr.firstIndex] == pr.character
	secondMatch := charStrings[pr.secondIndex] == pr.character

	return (firstMatch || secondMatch) && !(firstMatch && secondMatch)
}

func passwordPolicyCheck(value string) bool {
	record := passwordRecordParser(value)
	return record.isValid()
}

func passwordRecordParser(value string) passwordRecord {
	parts := strings.Split(value, " ")

	countRange, suffixedChar, password := parts[0], parts[1], parts[2]

	rangeParts := strings.Split(countRange, "-")

	first, err := strconv.Atoi(rangeParts[0])
	check(err)
	second, err := strconv.Atoi(rangeParts[1])
	check(err)

	return passwordRecord{
		character:   strings.Split(suffixedChar, ":")[0],
		firstIndex:  first - 1,
		secondIndex: second - 1,
		password:    password,
	}
}

func main() {
	matches := matchingLines("input.txt", passwordPolicyCheck)
	fmt.Printf("%v\n", matches)
}
