package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(err error) {
	if err == nil {
		return
	}

	panic(err)
}

func fileInts(path string) (ints []int) {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		check(err)
		ints = append(ints, i)
	}

	return ints
}

func main() {
	ints := fileInts("input.txt")

	seen := make(map[int]bool)

	for _, i := range ints {
		diff := 2020 - i
		fmt.Printf("%v, %v, %v\n", i, diff, seen[diff])

		if seen[diff] {
			product := i * diff
			fmt.Printf("%v", product)
			break
		}

		seen[i] = true
	}
}
