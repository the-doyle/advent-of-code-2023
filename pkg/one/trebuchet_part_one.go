package one

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type Digit struct {
	first      int
	last       int
	first_last int

	firstSet bool
	lastSet  bool
}

// TrebuchetPartOne calculates the sum of the values
func TrebuchetPartOne() {
	sum := 0

	values, err := read_calibration_doc("/Users/bendoyle/Documents/projects/advent-of-code-2023/pkg/one/trebuchet_calibration.txt")
	if err != nil {
		log.Println(err)
	}

	for _, value := range values {
		d := Digit{
			first: 0,
			last:  len(value) - 1,
		}

		if err := d.first_last_digits(value); err != nil {
			log.Println(err)
		}

		sum += d.first_last
	}

	log.Println(sum)
}

// read_calibration_doc reads the list of trebuchet calibration values
func read_calibration_doc(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// first_last_digits returns a 2-digit integer from a string value
func (d *Digit) first_last_digits(value string) error {
	// Find the first digit in value
	for !d.firstSet {
		if first, err := strconv.Atoi(string(value[d.first])); err == nil {
			d.firstSet = true
			d.first = first
		} else {
			d.first++
		}
	}

	// Find the last digit in value
	for !d.lastSet {
		if last, err := strconv.Atoi(string(value[d.last])); err == nil {
			d.lastSet = true
			d.last = last
		} else {
			d.last--
		}
	}

	// Convert the first and last digits to int and return
	d.first_last = d.first*10 + d.last
	return nil
}
