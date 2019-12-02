package main

import "testing"

func TestCalculateFuel(t *testing.T) {
	testFuel(t, 12, 2)
	testFuel(t, 14, 2)
	testFuel(t, 1969, 654)
	testFuel(t, 100756, 33583)
}

func testFuel(t *testing.T, weight int, expectedFuel int) {
	var actualFuel = CalculateFuel(weight)

	if actualFuel != expectedFuel {
		t.Errorf("fuel was bad, got: %d, wanted: %d", actualFuel, expectedFuel)
	}
}
