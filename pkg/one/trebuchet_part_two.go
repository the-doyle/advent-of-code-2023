package one

import (
	"log"
	"strconv"
	"strings"
)

var StringDigits = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

// TrebuchetPartTwo calculates the sum of the values, considering string values as well as integers
func TrebuchetPartTwo() interface{} {
	sum := 0

	values, err := read_calibration_doc("/Users/bendoyle/Documents/projects/advent-of-code-2023/pkg/one/trebuchet_calibration.txt")
	if err != nil {
		log.Println(err)
		return sum
	}

	for _, value := range values {
		d := Digit{
			first: 0,
			last:  len(value) - 1,
		}

		if err := d.first_last_digits_part_two(value); err != nil {
			log.Println(err)
		}

		sum += d.first_last
	}

	return sum
}

// first_last_digits_part_two returns a 2-digit integer from a string value.
// This function also considers the possibility of a string integer, e.g. "four"
//
// I'll still check for a regular integer first.
// Then, I'll use a sliding window of five characters to check for the possiblity of a string integer
func (d *Digit) first_last_digits_part_two(value string) error {
	// Find the first digit in value
outerFirst:
	for !d.firstSet {
		if first, err := strconv.Atoi(string(value[d.first])); err == nil {
			d.firstSet = true
			d.first = first
			break
		}

		for i, stringDigit := range StringDigits {
			if d.first+len(stringDigit) <= len(value) {
				potentialStringDigit := value[d.first : d.first+len(stringDigit)]
				if strings.EqualFold(potentialStringDigit, stringDigit) {
					d.firstSet = true
					d.first = i + 1
					break outerFirst
				}
			}
		}

		d.first++
	}

	// Find the last digit in value
outerLast:
	for !d.lastSet {
		if last, err := strconv.Atoi(string(value[d.last])); err == nil {
			d.lastSet = true
			d.last = last
			break
		}

		for i, stringDigit := range StringDigits {
			if d.last-len(stringDigit) >= 0 {
				potentialStringDigit := value[d.last-len(stringDigit)+1 : d.last+1]
				if strings.EqualFold(potentialStringDigit, stringDigit) {
					d.lastSet = true
					d.last = i + 1
					break outerLast
				}
			}
		}

		d.last--
	}

	// Convert the first and last digits to int and return
	d.first_last = d.first*10 + d.last
	return nil
}
