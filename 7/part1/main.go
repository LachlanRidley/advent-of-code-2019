package main

import (
	"fmt"
	"strconv"
	"strings"
)

func stringToIntSlice(opcode string) []int {
	var opcodeStrings = strings.Split(opcode, ",")

	var opcodeSlice []int

	for _, value := range opcodeStrings {
		number, _ := strconv.Atoi(value)
		opcodeSlice = append(opcodeSlice, number)
	}

	return opcodeSlice
}

func intSliceToString(opcode []int) string {
	var opcodeStrings []string
	for _, value := range opcode {
		opcodeStrings = append(opcodeStrings, strconv.Itoa(value))
	}

	return strings.Join(opcodeStrings, ",")
}

type command struct {
	opcode     int
	paramModes [3]int
}

func (c command) String() string {
	return fmt.Sprintf("%d -> %d %d %d", c.opcode, c.paramModes[0], c.paramModes[1], c.paramModes[2])
}

func parseCommand(param int) command {
	command := command{-1, [3]int{0, 0, 0}}

	command.opcode = param % 100

	remaining := param / 100

	for i := 0; remaining > 0; i++ {
		command.paramModes[i] = remaining % 10
		remaining = remaining / 10
	}

	return command
}

func getValue(command command, paramIndex int, memory []int, pointer int) int {
	if command.paramModes[paramIndex] == 0 {
		return memory[memory[pointer+(paramIndex+1)]]
	}

	return memory[pointer+(paramIndex+1)]
}

type amplifier struct {
	program         []int
	inputs          []int
	pointer         int
	inputPointer    int
	output          int
	waitingForInput bool
	ended           bool
}

func CreateAmplifier(program string, phaseSetting int) amplifier {
	return amplifier{stringToIntSlice(program), []int{phaseSetting}, 0, 0, 0, false, false}
}

func (amp *amplifier) sendInput(input int) {
	amp.inputs = append(amp.inputs, input)
	amp.waitingForInput = false
}

func (amp *amplifier) advancePointer(steps int) {
	amp.pointer = amp.pointer + steps
}

func (amp *amplifier) runNextInstruction() {
	if amp.program[amp.pointer] == 99 {
		amp.ended = true
		return
	}

	var command = parseCommand(amp.program[amp.pointer])

	if command.opcode == 1 {
		// addition
		value1 := getValue(command, 0, amp.program, amp.pointer)
		value2 := getValue(command, 1, amp.program, amp.pointer)

		amp.program[amp.program[amp.pointer+3]] = value1 + value2
		amp.advancePointer(4)
	} else if command.opcode == 2 {
		// multiplication
		value1 := getValue(command, 0, amp.program, amp.pointer)
		value2 := getValue(command, 1, amp.program, amp.pointer)

		amp.program[amp.program[amp.pointer+3]] = value1 * value2
		amp.advancePointer(4)
	} else if command.opcode == 3 {
		// receive input

		if amp.inputPointer >= len(amp.inputs) {
			amp.waitingForInput = true
			return
		}

		amp.program[amp.program[amp.pointer+1]] = amp.inputs[amp.inputPointer]
		amp.inputPointer++

		amp.advancePointer(2)
	} else if command.opcode == 4 {
		// print output
		output := getValue(command, 0, amp.program, amp.pointer)

		amp.advancePointer(2)

		amp.output = output
	} else if command.opcode == 5 {
		// jump-if-true
		if getValue(command, 0, amp.program, amp.pointer) != 0 {
			amp.pointer = getValue(command, 1, amp.program, amp.pointer)
		} else {
			amp.advancePointer(3)
		}
	} else if command.opcode == 6 {
		// jump-if-false
		if getValue(command, 0, amp.program, amp.pointer) == 0 {
			amp.pointer = getValue(command, 1, amp.program, amp.pointer)
		} else {
			amp.advancePointer(3)
		}
	} else if command.opcode == 7 {
		// less than
		if getValue(command, 0, amp.program, amp.pointer) < getValue(command, 1, amp.program, amp.pointer) {
			amp.program[amp.program[amp.pointer+3]] = 1
		} else {
			amp.program[amp.program[amp.pointer+3]] = 0
		}

		amp.advancePointer(4)
	} else if command.opcode == 8 {
		// equals
		if getValue(command, 0, amp.program, amp.pointer) == getValue(command, 1, amp.program, amp.pointer) {
			amp.program[amp.program[amp.pointer+3]] = 1
		} else {
			amp.program[amp.program[amp.pointer+3]] = 0
		}

		amp.advancePointer(4)
	}
}

func RunAmplifierControllerSoftware(program string, phaseSettings []int) int {
	var amps []amplifier

	for _, phaseSetting := range phaseSettings {
		amps = append(amps, CreateAmplifier(program, phaseSetting))
	}

	amps[0].sendInput(0)

	currentActiveAmp := 0

	for !amps[4].ended {
		nextAmp := currentActiveAmp + 1

		if nextAmp > 4 {
			nextAmp = 0
		}

		for !amps[currentActiveAmp].waitingForInput && !amps[currentActiveAmp].ended {
			amps[currentActiveAmp].runNextInstruction()
		}

		// fmt.Printf("%d -> %d (%d)\n", currentActiveAmp, nextAmp, amps[currentActiveAmp].output)

		if amps[4].ended {
			return amps[currentActiveAmp].output
		}

		amps[nextAmp].sendInput(amps[currentActiveAmp].output)

		currentActiveAmp++

		if currentActiveAmp > 4 {
			currentActiveAmp = 0
		}
	}

	return -1
}

func generate(n int, a []int) [][]int {
	var perms [][]int

	c := make([]int, n)
	for i := 0; i < n; i++ {
		c[i] = 0
	}

	perms = append(perms, append([]int(nil), a...))

	for i := 0; i < n; {
		if c[i] < i {
			if i%2 == 0 {
				Swap(a, 0, i)
			} else {
				Swap(a, c[i], i)
			}

			perms = append(perms, append([]int(nil), a...))

			c[i] = c[i] + 1

			i = 0
		} else {
			c[i] = 0
			i++
		}
	}

	return perms
}

func Swap(arr []int, i1 int, i2 int) {
	i1Value := arr[i1]
	i2Value := arr[i2]

	arr[i2] = i1Value
	arr[i1] = i2Value
}

func main() {
	perms := generate(5, []int{9, 8, 7, 6, 5})

	program := "3,8,1001,8,10,8,105,1,0,0,21,38,55,72,93,118,199,280,361,442,99999,3,9,1001,9,2,9,1002,9,5,9,101,4,9,9,4,9,99,3,9,1002,9,3,9,1001,9,5,9,1002,9,4,9,4,9,99,3,9,101,4,9,9,1002,9,3,9,1001,9,4,9,4,9,99,3,9,1002,9,4,9,1001,9,4,9,102,5,9,9,1001,9,4,9,4,9,99,3,9,101,3,9,9,1002,9,3,9,1001,9,3,9,102,5,9,9,101,4,9,9,4,9,99,3,9,101,1,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,101,1,9,9,4,9,99,3,9,101,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,99,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,99,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,101,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,99"

	highestOutput := 0
	for _, phaseSettings := range perms {

		// program := "3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5"
		// phaseSettings := []int{9, 8, 7, 6, 5}

		// program := "3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10"
		// phaseSettings := []int{9, 7, 8, 5, 6}

		output := RunAmplifierControllerSoftware(program, phaseSettings)

		if output > highestOutput {
			highestOutput = output
		}

	}

	fmt.Printf("%d\n", highestOutput)
}
