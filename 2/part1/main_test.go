package main

import "testing"

func TestRunOpcode(t *testing.T) {
	runOpcode(t, "1,0,0,0,99", "2,0,0,0,99")
	runOpcode(t, "2,3,0,3,99", "2,3,0,6,99")
	runOpcode(t, "2,4,4,5,99,0", "2,4,4,5,99,9801")
	runOpcode(t, "1,1,1,4,99,5,6,0,99", "30,1,1,4,2,5,6,0,99")
}

func runOpcode(t *testing.T, initialOpcode string, expectedResult string) {
	var actualResult = RunOpcode(initialOpcode)

	if actualResult != expectedResult {
		t.Errorf("opcode was bad, got: %s, wanted: %s", actualResult, expectedResult)
	}
}
