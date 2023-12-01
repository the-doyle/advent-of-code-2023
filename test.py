# Open the input.txt file in read mode
with open('pkg/one/trebuchet_calibration.txt', 'r') as f:
    # Get all the lines in the input file as a list
    lines = f.read().splitlines()
# Intialize a dictionary of strings of numbers, where the key value is the string format of a number and the value is the number representation
number_strings = {"one" : "1", "two" : "2", "three" : "3", "four" : "4", "five" : "5", "six" : "6", "seven" : "7", "eight" : "8", "nine" : "9"}
# Initialize the result as 0
result = 0
# Loop through the lines list
for line in lines:
    # For each line, initialize a first and a last digit as None
    first_digit, last_digit = None, None
    # Go through each index in the line
    for i in range(len(line)):
        # If the character at the current index is a numeric
        if line[i].isnumeric():
            # if the first_digit has not been assigned then assign the first and the last digit the current value
            if first_digit == None:
                first_digit, last_digit = line[i], line[i]
            # if the first digit has already been assigned then change the value of the last digit
            else:
                last_digit = line[i]
        # if the current character isn't a numeric, then check if it matches any of the number strings
        else:
            # Loop through the number_strings dictionary
            for key, val in number_strings.items():
                # Check if the value from the current_index to current_index + len(key) matches the key value
                if line[i : i + len(key)] == key:
                    # if the first_digit has not been assigned, then assign the first and the last digit the value of number_strings[key]
                    if first_digit == None:
                        first_digit, last_digit = val, val
                    # if the first_digit has been assigned, then update the last digit to the value of number_strings[key]
                    else:
                        last_digit = val
    # Add the current calibration value to the result variable at the end of each line 
    result += int(first_digit + last_digit)
# Print the final result
print(result)