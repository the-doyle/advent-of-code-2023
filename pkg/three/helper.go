package three

import (
	"bufio"
	"os"
	"strings"
)

// read_calibration_doc reads the list of trebuchet calibration values
func read_input() (input [][]string, hasAdjacentSymbol [][]bool, err error) {
	file, err := os.Open("/Users/bendoyle/Documents/projects/advent-of-code-2023/pkg/three/input.txt")
	if err != nil {
		return input, hasAdjacentSymbol, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, strings.Split(scanner.Text(), ""))
		hasAdjacentSymbol = append(hasAdjacentSymbol, make([]bool, len(scanner.Text())))
	}
	return input, hasAdjacentSymbol, scanner.Err()
}
