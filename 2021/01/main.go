package main

func main() {
	current := int64(-1)
	increases := -1

	for integer := range stdinInts() {
		if integer > current {
			echo("%v >  %v", integer, current)
			increases++
		} else {
			echo("%v <= %v", integer, current)
		}

		current = integer
	}

	echo("%v increases", increases)
}
