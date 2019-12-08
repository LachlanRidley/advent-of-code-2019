package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func count(slice []int, number int) int {
	count := 0
	for _, v := range slice {
		if v == number {
			count++
		}
	}

	return count
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func solvePartOne(layers [][]int, layerSize int) int {
	fewestZeroes := -1
	fewestZeroesCount := layerSize + 1
	for i, layer := range layers {
		count := count(layer, 0)

		if count < fewestZeroesCount {
			fewestZeroesCount = count
			fewestZeroes = i
		}
	}

	return count(layers[fewestZeroes], 1) * count(layers[fewestZeroes], 2)
}

func solvePartTwo(layers [][]int, layerSize int) {
	// create an array[layerSize] for the final, flattened image
	finalLayer := make([]int, layerSize)

	// iterate through layers backwards
	for i := len(layers) - 1; i >= 0; i-- {
		// for each pixel, overwrite position in image array unless transparent
		for j := layerSize - 1; j >= 0; j-- {
			if layers[i][j] == 2 {
				continue
			}

			finalLayer[j] = layers[i][j]
		}
	}

	// turn into 2D array
	// decide if you can be fucked turning into an image...
	// otherwise just print it out to the terminal and work it out

	for i, v := range finalLayer {
		if i > 0 && i%25 == 0 {
			fmt.Printf("\n")
		}
		fmt.Printf("%d,", v)
	}
}

func main() {
	const width = 25
	const height = 6
	const layerSize = width * height

	file, err := os.Open("puzzle.txt")
	check(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	var layers [][]int
	var buffer []int

	for i := 0; scanner.Scan(); i++ {

		if i > 0 && i%layerSize == 0 {
			layers = append(layers, buffer)
			buffer = []int{}
		}

		str := scanner.Text()

		digit, _ := strconv.Atoi(str)

		buffer = append(buffer, digit)
	}

	// fmt.Println(solvePartOne(layers, layerSize))
	solvePartTwo(layers, layerSize)
}
