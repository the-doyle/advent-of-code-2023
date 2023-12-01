package one

import (
	"log"
	"strconv"
	"strings"
)

var StringInts = []string{
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
	}

	d := CalibrationValue{}

	for _, rawValue := range values {
		d.first, d.last = 0, len(rawValue)-1
		d.rawValue = rawValue
		d.find_value_enhanced()
		sum += d.value
	}

	return sum
}

// find_value_enhanced also considers the possibility of a string integer, e.g. "four"
func (d *CalibrationValue) find_value_enhanced() {
	// Find the first digit
outerFirst:
	for {
		// Check for a regular digit, e.g. "4"
		if first, err := strconv.Atoi(string(d.rawValue[d.first])); err == nil {
			d.first = first
			break
		}

		// Iterate over each possible stringInt and check if it exists from the current index
		for i, stringInt := range StringInts {

			// Make sure the stringInt fits before checking
			if d.first+len(stringInt) <= len(d.rawValue) {

				// Now check
				potentialStringInt := d.rawValue[d.first : d.first+len(stringInt)]

				// If there's a match, save the stringInt (using it's index + 1) to d.first and break out of the outer loop
				if strings.EqualFold(potentialStringInt, stringInt) {
					d.first = i + 1
					break outerFirst
				}
			}
		}

		d.first++
	}

	// Find the last digit
outerLast:
	for {
		if last, err := strconv.Atoi(string(d.rawValue[d.last])); err == nil {
			d.last = last
			break
		}

		for i, stringInt := range StringInts {
			if d.last-len(stringInt) >= 0 {
				potentialStringInt := d.rawValue[d.last-len(stringInt)+1 : d.last+1]
				if strings.EqualFold(potentialStringInt, stringInt) {
					d.last = i + 1
					break outerLast
				}
			}
		}

		d.last--
	}

	// Convert the first and last digits to int and return
	d.value = d.first*10 + d.last
}
