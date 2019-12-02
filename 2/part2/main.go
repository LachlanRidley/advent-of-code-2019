package main

import (
	"fmt"
	"strings"
	"strconv"
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

func RunOpcode(opcode string) string {
	var intSlice = stringToIntSlice(opcode)

	var pointer = 0

	for intSlice[pointer] != 99 {
		var command = intSlice[pointer]
		var value1 = intSlice[intSlice[pointer + 1]]
		var value2 = intSlice[intSlice[pointer + 2]]
		var resultIndex = intSlice[pointer + 3]

		if command == 1 {
			intSlice[resultIndex] = value1 + value2
		} else if command == 2 {
			intSlice[resultIndex] = value1 * value2
		}

		pointer = pointer + 4
	}

	return intSliceToString(intSlice)
}

func main() {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			var sb strings.Builder

			sb.WriteString("1,")
			sb.WriteString(strconv.Itoa(noun))
			sb.WriteString(",")
			sb.WriteString(strconv.Itoa(verb))
			sb.WriteString(",3,1,1,2,3,1,3,4,3,1,5,0,3,2,1,13,19,1,10,19,23,2,9,23,27,1,6,27,31,1,10,31,35,1,35,10,39,1,9,39,43,1,6,43,47,1,10,47,51,1,6,51,55,2,13,55,59,1,6,59,63,1,10,63,67,2,67,9,71,1,71,5,75,1,13,75,79,2,79,13,83,1,83,9,87,2,10,87,91,2,91,6,95,2,13,95,99,1,10,99,103,2,9,103,107,1,107,5,111,2,9,111,115,1,5,115,119,1,9,119,123,2,123,6,127,1,5,127,131,1,10,131,135,1,135,6,139,1,139,5,143,1,143,9,147,1,5,147,151,1,151,13,155,1,5,155,159,1,2,159,163,1,163,6,0,99,2,0,14,0")

			result := RunOpcode(sb.String())

			if stringToIntSlice(result)[0] == 19690720 {
				fmt.Printf("%d\n", 100 * noun + verb)
			}
		}
	}
}
