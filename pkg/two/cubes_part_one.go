package two

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func CubesPartOne() interface{} {
	file, err := os.Open("/Users/bendoyle/Documents/projects/advent-of-code-2023/pkg/two/games.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	sum_possible_ids := 0

	scanner := bufio.NewScanner(file)

gamesLoop:
	for scanner.Scan() {
		if i, handfuls, found := strings.Cut(scanner.Text(), ": "); found {
			for _, handful := range strings.Split(handfuls, "; ") {
				for _, cube := range strings.Split(handful, ", ") {
					if count, color, found := strings.Cut(cube, " "); found {
						if c, err := strconv.Atoi(string(count)); err == nil {
							switch color {
							case "red":
								if c > 12 {
									continue gamesLoop
								}
							case "blue":
								if c > 14 {
									continue gamesLoop
								}
							case "green":
								if c > 13 {
									continue gamesLoop
								}
							}
						}
					}
				}
			}

			if _, idString, found := strings.Cut(i, " "); found {
				if id, err := strconv.Atoi(string(idString)); err == nil {
					sum_possible_ids += id
				}
			}
		}
	}

	return sum_possible_ids
}
