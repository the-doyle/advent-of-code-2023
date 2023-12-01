# Trebuchet Part One
For the first problem, I created a `Digit` struct to store information about each string. 
```
type Digit struct {
	first      int  // Index of first digit; after digit is found, this becomes the value of the first digit
	last       int  // Index of last digit; after digit is found, this becomes the value of the last digit
	firstSet   bool // Whether the first digit has been found
	lastSet    bool // Whether the last digit has been found
	first_last int  // Combined first and last digits
}
``` 

I decided to use the same attributes to store the pointer indices and the final digit values (`Digit.first` and `Digit.last`). 

Next, I created a function to find the 2-digit integer for each string value in the input. 
I decided to take a 2-pointer approach to find the first and last digits in each string. 
```
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
``` 

Putting it all together
```
// TrebuchetPartOne calculates the sum of the values
func TrebuchetPartOne() interface{} {
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

	return sum
}
```

Yielded the following results:

```
Trebuchet Part One
Result: 54597
Minimum time taken: 395.208µs
Maximum time taken: 833.583µs
Average time taken: 435.528µs
```

Let's see if I can optimize this any further. 