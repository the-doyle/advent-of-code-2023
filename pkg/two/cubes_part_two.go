package two

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func CubesPartTwo() interface{} {
	file, err := os.Open("/Users/bendoyle/Documents/projects/advent-of-code-2023/pkg/two/games.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	sum_powers := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		min_red := 0
		min_blue := 0
		min_green := 0
		if _, handfuls, found := strings.Cut(scanner.Text(), ": "); found {
			for _, handful := range strings.Split(handfuls, "; ") {
				for _, cube := range strings.Split(handful, ", ") {
					if count, color, found := strings.Cut(cube, " "); found {
						if c, err := strconv.Atoi(string(count)); err == nil {
							switch color {
							case "red":
								if c > min_red {
									min_red = c
								}
							case "blue":
								if c > min_blue {
									min_blue = c
								}
							case "green":
								if c > min_green {
									min_green = c
								}
							}
						}
					}
				}
			}

			sum_powers += min_red * min_blue * min_green
		}
	}

	return sum_powers
}
