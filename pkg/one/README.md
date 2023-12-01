
https://adventofcode.com/2023/day/1

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

Yielded the following results (10,000 test runs):

```
Trebuchet Part One
Result: 54597
Minimum time taken: 393.083µs
Maximum time taken: 1.47675ms
Average time taken: 423.346µs
```

I feel my first approach is pretty decent, but perhaps overly wordy/complicated. Let's see if I can make things simpler. 

First, I'll rename my struct to `CalibrationValue` to reflect the information I'm storing. 
I'll also add the raw input value so I don't have to pass it to my function each time. 
I decided to restructure my loops to remove the need for the `firstSet` and `lastSet` attributes, so I'll remove those. 
```
type CalibrationValue struct {
	rawValue string // The raw calibration value
	first    int    // Index of first digit; after digit is found, this becomes the value of the first digit
	last     int    // Index of last digit; after digit is found, this becomes the value of the last digit
	value    int    // Combined first and last digits of the rawValue
}
```

Instead, I'll break out of the loop once the value is found. I also renamed `first_last_digits()` to something more descriptive. 

And finally, it isn't necessary to return the error I'm checking for, because it simply means a character in the string is not a digit. 

```
// find_value finds the first and last digits in a raw CalibrationValue
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
```

These changes make my main function a bit simpler, too
```
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
```

These changes didn't make my code run faster in a significant way, but I do feel things are much more readable. On to part two! 

# Trebuchet Part Two 

Part two introduces the need to check for a string integer inside each `rawValue` — e.g. "four" — in addition to regular integers. 

The first thing I decided to try was creating a constant slice of string integers
```
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
```

After first checking for a regular integer, I added logic to check for string integers
```
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

    etc... 
}
```

The logic is similar (reversed) for finding the last digit. 

Interestingly, my code for Part Two runs slightly faster than Part One. 

```
Trebuchet Part Two
Result: 54504
Minimum time taken: 356.209µs
Maximum time taken: 1.46825ms
Average time taken: 374.367µs
```

I can't think of any speed optimizations for Part Two, either. The main factor for speed in Part Two is the need to check for each `StringInt` at each index in the string. 

Implementing `StringInts` as a slice (instead of a `map`) was the correct choice because it uses less memory and while allowing me to calculate the actual integer value based on the index of each `StringInt`. 

# Other thoughts 

To make PartOne and PartTwo more readable, I suppose I could do everything in a single pass (instead of going from the start and end of each string). Here's what that would look like for Part One: 

```
// find_value finds the first and last digits in a raw CalibrationValue
func (d *CalibrationValue) find_value() {

	firstSet := false

	for i := 0; i < len(d.rawValue); i++ {
		if digit, err := strconv.Atoi(string(d.rawValue[i])); err == nil {

			// Set first if unset
			if !firstSet {
				d.first = digit
				firstSet = true
			}

			// Set last every time
			d.last = digit
		}
	}

	// Convert the first and last digits to int and return
	d.value = d.first*10 + d.last
}
```

The code is simpler because only loop is required, but...

```
Trebuchet Part One
Result: 54597
Minimum time taken: 550µs
Maximum time taken: 2.20325ms
Average time taken: 596.358µs
```

It takes ~40% longer to run. Why? 

* Iterating over the string start-to-finish means I will always process each character. 
* Using two loops to iterate from the start, then the back, means I might not always process each character. 

For a string like `ad1asdfefsfrOnefour54two`, the single-loop approach has to process all characters, while the two-loop approach only process the first 3 and the last 4 characters. Much faster. 

So I'll stick with my first approach for speed. 
