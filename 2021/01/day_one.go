package main

func dayOne(windowSize int) {
	window := make([]int64, 0)
	increases := 0

	for integer := range stdinInts() {
		previous := window

		window = append(window, integer)

		if len(window) > windowSize {
			window = window[1:]
		}

		if len(previous) == windowSize && sum(window) > sum(previous) {
			increases++
		}
	}

	echo("%v increases", increases)
}
