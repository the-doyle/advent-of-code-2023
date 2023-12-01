package one

import (
	"bufio"
	"os"
)

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
