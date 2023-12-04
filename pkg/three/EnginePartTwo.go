package three

import (
	"strconv"
	"strings"
)

// For part two, I should keep a map of gears
// Key = location in the grid
// Value = []Nums that the gears is attached to

func EnginePartTwo() interface{} {

	input, _, _ := read_input()

	gears := make(map[string][]int, 0)
	var numAdjacentGears []string

	num := 0
	sum := 0

	for y, yRow := range input {
		for x := range yRow {

			// If number is found, check for adjacent gears.
			// Then, continue getting the number
			if n, err := strconv.Atoi(string((input)[y][x])); err == nil {
				checkAdjacentGears(x, y, x+1, y, &input, &numAdjacentGears)
				checkAdjacentGears(x, y, x-1, y, &input, &numAdjacentGears)
				checkAdjacentGears(x, y, x+1, y+1, &input, &numAdjacentGears)
				checkAdjacentGears(x, y, x+1, y-1, &input, &numAdjacentGears)
				checkAdjacentGears(x, y, x, y+1, &input, &numAdjacentGears)
				checkAdjacentGears(x, y, x, y-1, &input, &numAdjacentGears)
				checkAdjacentGears(x, y, x-1, y+1, &input, &numAdjacentGears)
				checkAdjacentGears(x, y, x-1, y-1, &input, &numAdjacentGears)

				// brand new number
				if num != 0 {
					num = num*10 + n
				} else {
					num = n
				}

				continue
			}

			// Add number to map of gears
			if num != 0 {
				for _, gear := range numAdjacentGears {
					gears[gear] = append(gears[gear], num)
				}
			}

			// Reset num, numAdjacentGears
			num = 0
			numAdjacentGears = nil
		}

		// At the end of each row, add number and reset again
		if num != 0 {
			for _, gear := range numAdjacentGears {
				gears[gear] = append(gears[gear], num)
			}
		}
		num = 0
		numAdjacentGears = nil
	}

	for y, yRow := range input {
		for x := range yRow {
			if nums, found := gears[strconv.Itoa(y)+"-"+strconv.Itoa(x)]; found && len(nums) == 2 {
				sum += nums[0] * nums[1]
			}
		}
	}

	return sum
}

func checkAdjacentGears(x int, y int, checkX int, checkY int, input *[][]string, adjacentGears *[]string) {
	// Out of bounds check
	if checkX < 0 || checkX > len((*input)[0])-1 || checkY < 0 || checkY > len((*input))-1 {
		return
	}

	if string((*input)[checkY][checkX]) == "*" {
		for _, ag := range *adjacentGears {
			if strings.EqualFold(ag, strconv.Itoa(checkY)+"-"+strconv.Itoa(checkX)) {
				return
			}
		}
		*adjacentGears = append(*adjacentGears, strconv.Itoa(checkY)+"-"+strconv.Itoa(checkX))
	}
}
