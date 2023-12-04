package main

import (
	"fmt"
	"time"

	"github.com/the-doyle/advent-of-code-2023/pkg/three"
)

func main() {
	// measureExecutionTime(one.TrebuchetPartOne, "Trebuchet Part One")
	// measureExecutionTime(one.TrebuchetPartTwo, "Trebuchet Part Two")

	// measureExecutionTime(two.CubesPartOne, "Cubes Part One")
	// measureExecutionTime(two.CubesPartTwo, "Cubes Part Two")

	three.EnginePartTwo()

}

// measureExecutionTime takes another function as a parameter, runs it 10k times,
// and returns the result as well as the min/max/avg time taken
func measureExecutionTime(f func() interface{}, name string) {
	runs := 10000

	var totalTime time.Duration

	min := time.Duration(int(^uint(0) >> 1))
	var max time.Duration
	var result interface{}

	for i := 0; i < runs; i++ {
		start := time.Now()
		result = f()
		elapsed := time.Since(start)

		totalTime += elapsed

		if elapsed < min {
			min = elapsed
		}
		if elapsed > max {
			max = elapsed
		}
	}

	avg := totalTime / time.Duration(runs)

	fmt.Println("————————————————————————————-")
	fmt.Printf(name + "\n")
	fmt.Printf("Result: %v\n", result)
	fmt.Printf("Minimum time taken: %v\n", min.String())
	fmt.Printf("Maximum time taken: %v\n", max.String())
	fmt.Printf("Average time taken: %v\n", avg.String())
	fmt.Println()
}
