package one

import (
	"log"
	"strconv"
)

type CalibrationValue struct {
	rawValue string // The raw calibration value
	first    int    // Index of first digit; after digit is found, this becomes the value of the first digit
	last     int    // Index of last digit; after digit is found, this becomes the value of the last digit
	value    int    // Combined first and last digits of the rawValue
}

// TrebuchetPartOne calculates the sum of the values
func TrebuchetPartOne() interface{} {
	sum := 0

	values, err := read_calibration_doc("/Users/bendoyle/Documents/projects/advent-of-code-2023/pkg/one/trebuchet_calibration.txt")
	if err != nil {
		log.Println(err)
	}

	d := CalibrationValue{}

	for _, rawValue := range values {
		d.first, d.last = 0, len(rawValue)-1
		d.rawValue = rawValue
		d.find_value()
		sum += d.value
	}

	return sum
}

// find_value finds the first and last digits in the raw string value
func (d *CalibrationValue) find_value() {
	// Find the first digit
	for {
		if first, err := strconv.Atoi(string(d.rawValue[d.first])); err == nil {
			d.first = first
			break
		}
		d.first++
	}

	// Find the last digit
	for {
		if last, err := strconv.Atoi(string(d.rawValue[d.last])); err == nil {
			d.last = last
			break
		}
		d.last--
	}

	// Convert the first and last digits to int and return
	d.value = d.first*10 + d.last
}
