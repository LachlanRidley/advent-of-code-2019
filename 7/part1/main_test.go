package main

import "testing"

// func TestRunOpcode(t *testing.T) {
// 	runRunAmplifierControllerSoftware(t, 43210, "3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0", []int{4, 3, 2, 1, 0})
// 	runRunAmplifierControllerSoftware(t, 54321, "3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0", []int{0, 1, 2, 3, 4})
// 	runRunAmplifierControllerSoftware(t, 65210, "3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0", []int{1, 0, 4, 3, 2})
// }

func TestSwap(t *testing.T) {
	arr := []int{1, 2, 3}

	Swap(arr, 0, 2)

	if arr[0] != 3 {
		t.Errorf("got: %d, wanted: %d", arr[0], 3)
	}

	if arr[2] != 1 {
		t.Errorf("got: %d, wanted: %d", arr[2], 1)
	}
}

// func runRunAmplifierControllerSoftware(t *testing.T, expectedResult int, program string, phaseSettings []int) {
// 	var actualResult = RunAmplifierControllerSoftware(program, phaseSettings)

// 	if actualResult != expectedResult {
// 		t.Errorf("got: %d, wanted: %d", actualResult, expectedResult)
// 	}
// }

func TestVariousInstructions(t *testing.T) {
	amp := CreateAmplifier("3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0", 4)
	amp.runNextInstruction()

	if amp.program[15] != 4 {
		t.Errorf("got: %d, wanted: %d", amp.program[15], 4)
	}

	amp.runNextInstruction()

	if !amp.waitingForInput {
		t.Errorf("should be waiting for input")
	}
}

func TestPuzzle(t *testing.T) {
	program := "3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5"
	phaseSettings := []int{9, 8, 7, 6, 5}
	output := RunAmplifierControllerSoftware(program, phaseSettings)

	if output != 139629729 {
		t.Errorf("got: %d, wanted: %d", output, 139629729)
	}
}
