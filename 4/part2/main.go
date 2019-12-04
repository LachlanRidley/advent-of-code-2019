package main

import (
	"fmt"
)

func hasOnlyIncreasingDigits(code []int) bool {
	for i := 1; i < 6; i++ {
		if code[i] < code[i - 1] {
			return false
		}
	}

	return true
}

func hasDoubleDigits(code []int) bool {
	var counts []int
	count := 1

	for i := 1; i < 6; i++ {
		if code[i] != code[i - 1] {
			counts = append(counts, count)
			count = 1
		} else {
			count++
		}

		if i == 5 {
			counts = append(counts, count)
		}
	}

	fmt.Printf("%v\n", counts)

	return contains(counts, 2)
}

func IsValidCode(code int) bool {
	digits := convertToSlice(code)

	increasing := hasOnlyIncreasingDigits(digits)
	doubles := hasDoubleDigits(digits)

	fmt.Printf("%d increasing %t / doubles %t\n", code, increasing, doubles)

	return increasing && doubles
}

func reverse(numbers []int) []int {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

func contains(s []int, e int) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func convertToSlice(code int) []int {
	var digits []int

	for code > 0 {
		digit := code % 10
		digits = append(digits, digit)
		code = code / 10
	}

	digits = reverse(digits)

	return digits
}

func main() {
	validCodeCount := 0

	for i := 256310; i <= 732736; i++ {
	    isValid := IsValidCode(i)

	    if isValid {
			validCodeCount++
	    }

	    fmt.Printf("%d -> %t\n", i, isValid)
	}

	fmt.Printf("Valid codes: %d\n", validCodeCount)
}
