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

	seenIndex := make(map[int]int)

	for x, xv := range ints {
		seenIndex[xv] = x
	}

	for x, xv := range ints {
		for y, yv := range ints {
			if x == y {
				continue
			}

			diff := 2020 - xv - yv

			z, found := seenIndex[diff]

			if found && x != z && y != z {
				product := xv * yv * diff
				fmt.Printf("%v", product)
				return
			}
		}
	}
}
