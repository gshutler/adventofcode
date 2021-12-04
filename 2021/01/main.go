package main

import "os"

func main() {
	switch os.Getenv("DAY") {
	case "1.1":
		dayOne(1)
	case "1.2":
		dayOne(3)
	case "2.1":
		dayTwo()
	case "":
		echo("Must set DAY environment variable")
	default:
		echo("DAY %v not recognized", os.Getenv("DAY"))
	}
}
