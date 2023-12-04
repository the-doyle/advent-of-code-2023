package three

import (
	"strconv"
)

func EnginePartOne() interface{} {

	input, hasAdjacentSymbol, _ := read_input()

	numHasAdjacent := false
	num := 0
	sum := 0

	for y, yRow := range input {
		for x := range yRow {
			checkAdjacent(x, y, x+1, y, &input, &hasAdjacentSymbol)
			checkAdjacent(x, y, x-1, y, &input, &hasAdjacentSymbol)
			checkAdjacent(x, y, x+1, y+1, &input, &hasAdjacentSymbol)
			checkAdjacent(x, y, x+1, y-1, &input, &hasAdjacentSymbol)
			checkAdjacent(x, y, x, y+1, &input, &hasAdjacentSymbol)
			checkAdjacent(x, y, x, y-1, &input, &hasAdjacentSymbol)
			checkAdjacent(x, y, x-1, y+1, &input, &hasAdjacentSymbol)
			checkAdjacent(x, y, x-1, y-1, &input, &hasAdjacentSymbol)

			// If number is found, continue getting the number
			if n, err := strconv.Atoi(string((input)[y][x])); err == nil {

				// brand new number
				if num != 0 {
					num = num*10 + n
				} else {
					num = n
				}

				numHasAdjacent = numHasAdjacent || hasAdjacentSymbol[y][x]
				continue
			}

			// If we reach the end of a number, add to sum if it has an adjacent
			if num != 0 {
				if numHasAdjacent {
					sum += num
				}
			}

			// Reset num, numHasAdjacent
			num = 0
			numHasAdjacent = false
		}

		// At the end of each row, check again
		if num != 0 {
			if numHasAdjacent {
				sum += num
			}
		}

		// Reset num, numHasAdjacent
		num = 0
		numHasAdjacent = false
	}

	return sum
}

func checkAdjacent(x int, y int, checkX int, checkY int, input *[][]string, hasAdjacentSymbol *[][]bool) {
	// Out of bounds check
	if checkX < 0 || checkX > len((*input)[0])-1 || checkY < 0 || checkY > len((*input))-1 {
		return
	}

	if _, err := strconv.Atoi(string((*input)[checkY][checkX])); err != nil && string((*input)[checkY][checkX]) != "." {
		(*hasAdjacentSymbol)[y][x] = true
	}
}
