package main

import (
	"bufio"
	"fmt"
	"os"
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
	} else {
		return memory[pointer+(paramIndex+1)]
	}
}

func RunOpcode(opcode string) string {
	var intSlice = stringToIntSlice(opcode)

	var pointer = 0

	for intSlice[pointer] != 99 {
		var command = parseCommand(intSlice[pointer])

		// addition
		if command.opcode == 1 {
			value1 := getValue(command, 0, intSlice, pointer)
			value2 := getValue(command, 1, intSlice, pointer)

			intSlice[intSlice[pointer+3]] = value1 + value2
			pointer = pointer + 4
		} else if command.opcode == 2 {
			// multiplication
			value1 := getValue(command, 0, intSlice, pointer)
			value2 := getValue(command, 1, intSlice, pointer)

			intSlice[intSlice[pointer+3]] = value1 * value2
			pointer = pointer + 4
		} else if command.opcode == 3 {
			// receive input
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter an int: ")
			text, _ := reader.ReadString('\n')
			text = strings.Replace(text, "\n", "", -1)

			var value, _ = strconv.Atoi(text)

			intSlice[intSlice[pointer+1]] = value

			pointer = pointer + 2
		} else if command.opcode == 4 {
			// print output
			fmt.Printf("%d\n", getValue(command, 0, intSlice, pointer))

			pointer = pointer + 2
		} else if command.opcode == 5 {
			// jump-if-true
			if getValue(command, 0, intSlice, pointer) != 0 {
				pointer = getValue(command, 1, intSlice, pointer)
			} else {
				pointer = pointer + 3
			}
		} else if command.opcode == 6 {
			// jump-if-false
			if getValue(command, 0, intSlice, pointer) == 0 {
				pointer = getValue(command, 1, intSlice, pointer)
			} else {
				pointer = pointer + 3
			}
		} else if command.opcode == 7 {
			// less than
			if getValue(command, 0, intSlice, pointer) < getValue(command, 1, intSlice, pointer) {
				intSlice[intSlice[pointer+3]] = 1
			} else {
				intSlice[intSlice[pointer+3]] = 0
			}

			pointer = pointer + 4
		} else if command.opcode == 8 {
			// equals
			if getValue(command, 0, intSlice, pointer) == getValue(command, 1, intSlice, pointer) {
				intSlice[intSlice[pointer+3]] = 1
			} else {
				intSlice[intSlice[pointer+3]] = 0
			}

			pointer = pointer + 4
		} else {
			// something broke, bail out
			return intSliceToString(intSlice)
		}

		// fmt.Printf("%d %s\n", pointer, intSliceToString(intSlice))
	}

	return intSliceToString(intSlice)
}

func main() {
	arg := os.Args[1]
	RunOpcode(arg)
}
