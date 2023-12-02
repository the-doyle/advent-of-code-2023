https://adventofcode.com/2023/day/2

# Cubes Part One 
Today's problem felt much easier than yesterday. I spent most of my time getting the input data into a workable format. 

To structure the list of input data
```
Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 2 green, 1 blue; 1 red, 2 green; 3 red, 1 blue; 2 blue, 1 green, 8 red; 1 green, 10 red; 10 red
Game 3: 14 red, 9 green, 5 blue; 2 green, 5 red, 7 blue; 1 blue, 14 green; 6 green, 2 red
...
```

I created a few simple types 
```
type Cube struct {
	count int
	color string
}

type Handful []Cube
type Game []Handful
```

From there, reading the input data was a matter of some string splitting and such 
```
// read_games reads the list of games
func read_games() ([]Game, error) {
	file, err := os.Open("/Users/bendoyle/Documents/projects/advent-of-code-2023/pkg/two/games.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var games []Game
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if _, handfuls, found := strings.Cut(scanner.Text(), ": "); found {
			g := Game{}
			for _, handful := range strings.Split(handfuls, "; ") {
				h := Handful{}
				for _, cube := range strings.Split(handful, ", ") {
					if count, color, found := strings.Cut(cube, " "); found {
						if c, err := strconv.Atoi(string(count)); err == nil {
							h = append(h, Cube{
								count: c,
								color: color,
							})
						}
					}
				}
				g = append(g, h)
			}
			games = append(games, g)
		}
	}

	return games, scanner.Err()
}
```

Once the data was parsed into these structs, it made it very easy to solve part one. 
I iterated over each game, handful of cubes, and finally each cube. 
The threshholds for each cube are hardcoded into my switch statement, which continues the outer loop if a threshhold is exceeded. 

```
func CubesPartOne() interface{} {
	games, err := read_games()
	if err != nil {
		return err
	}

	sum_possible_ids := 0

gamesLoop:
	for id, game := range games {
		for _, handful := range game {
			for _, cube := range handful {
				switch cube.color {
				case "red":
					if cube.count > 12 {
						continue gamesLoop
					}
				case "blue":
					if cube.count > 14 {
						continue gamesLoop
					}
				case "green":
					if cube.count > 13 {
						continue gamesLoop
					}
				}
			}
		}
		sum_possible_ids += id + 1
	}

	return sum_possible_ids

}
```

```
Cubes Part One
Result: 2720
Minimum time taken: 122µs
Maximum time taken: 765.042µs
Average time taken: 131.548µs
```

And that's it for Part One! Pretty simple. 

# Cubes Part Two

Part Two was also very simple. I modified my switch statement to instead compare the `Cube.count` to a running maximum for the current game. 

```
func CubesPartTwo() interface{} {
	games, err := read_games()
	if err != nil {
		return err
	}

	sum_powers := 0

	for _, game := range games {
		min_red := 0
		min_blue := 0
		min_green := 0
		for _, handful := range game {
			for _, cube := range handful {
				switch cube.color {
				case "red":
					if cube.count > min_red {
						min_red = cube.count
					}
				case "blue":
					if cube.count > min_blue {
						min_blue = cube.count
					}
				case "green":
					if cube.count > min_green {
						min_green = cube.count
					}
				}
			}
		}
		sum_powers += min_red * min_blue * min_green
	}

	return sum_powers
}
```

Simple. 

```
Cubes Part Two
Result: 71535
Minimum time taken: 122.75µs
Maximum time taken: 736.292µs
Average time taken: 132.173µs
```

# Other thoughts 

Because I'm using a helper function to read and structure the data, I'm going over each row of data twice! 

While it's nice to have my data organized into structs, this problem is pretty simple. It would be more optimal to process the data as it's being read in, in a single pass. Let's try that for Part One: 

```
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
```

As expected, this runs about 2x faster than my previous approach, because it only requires one pass over the data 

```
Cubes Part One
Result: 2720
Minimum time taken: 57.959µs
Maximum time taken: 581.5µs
Average time taken: 63.152µs
```

Refactoring Part Two yielded similar improvements 
```
Cubes Part Two
Result: 71535
Minimum time taken: 74.208µs
Maximum time taken: 348.333µs
Average time taken: 78.288µs
```