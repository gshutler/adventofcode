package main

import (
	"strconv"
	"strings"
)

type movement struct {
	direction string
	units     int64
}

func dayTwo() {
	horizontal := int64(0)
	depth := int64(0)

	for movement := range stdinMovements() {
		echo("Movement %v", movement)

		switch movement.direction {
		case "forward":
			horizontal += movement.units
		case "down":
			depth += movement.units
		case "up":
			depth -= movement.units
		default:
			echo("Unrecognized direction %v", movement.direction)
		}
	}

	echo("Horizontal %v, depth %v, product %v", horizontal, depth, horizontal*depth)
}

func stdinMovements() chan movement {
	movements := make(chan movement)

	go func() {
		defer close(movements)

		for line := range stdinLines() {
			fields := strings.Fields(line)

			if len(fields) == 2 {
				parsedInt, err := strconv.ParseInt(fields[1], 10, 64)

				if err == nil {
					movements <- movement{
						direction: fields[0],
						units:     parsedInt,
					}
				}
			} else {
				echo("Could not parse: %v", line)
			}
		}
	}()

	return movements
}
