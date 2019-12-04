package main

import "testing"

func TestCalculateFuel(t *testing.T) {
	testCode(t, 111111, true)
	testCode(t, 223450, false)
	testCode(t, 123789, false)
}

func testCode(t *testing.T, code int, isValid bool) {
	validityCheck := IsValidCode(code)
	if validityCheck != isValid {
		t.Errorf("%d was bad, got: %t, wanted: %t", code, validityCheck, isValid)
	}
}
