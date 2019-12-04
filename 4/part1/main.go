package main

import (
	"fmt"
)

func hasDoubleDigits(code []int) bool {
	for i := 1; i < 6; i++ {
		if code[i] == code[i - 1] {
			return true
		}
	}

	return false
}

func hasOnlyIncreasingDigits(code []int) bool {
	for i := 1; i < 6; i++ {
		if code[i] < code[i - 1] {
			return false
		}
	}

	return true
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
