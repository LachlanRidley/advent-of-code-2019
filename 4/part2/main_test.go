package main

import "testing"

func TestCalculateFuel(t *testing.T) {
	testCode(t, 112233, true)
	testCode(t, 123444, false)
	testCode(t, 111122, true)
}

func testCode(t *testing.T, code int, isValid bool) {
	validityCheck := IsValidCode(code)
	if validityCheck != isValid {
		t.Errorf("%d was bad, got: %t, wanted: %t", code, validityCheck, isValid)
	}
}
