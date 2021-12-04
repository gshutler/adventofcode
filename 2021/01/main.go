package main

func main() {
	window := make([]int64, 0)
	increases := 0

	for integer := range stdinInts() {
		previous := window

		window = append(window, integer)

		if len(window) > 3 {
			window = window[1:]
		}

		if len(previous) == 3 && sum(window) > sum(previous) {
			increases++
		}
	}

	echo("%v increases", increases)
}
