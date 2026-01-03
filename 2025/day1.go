package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const DEFUALT_DIAL_LOCATION = 50
const LOWEST_NUM_ON_DIAL = 0
const HIGHEST_NUM_ON_DIAL = 99

func main() {
	input, err := getInputFromFile("input.txt")
	if err != nil {
		fmt.Print("Unable to read file")
	}
	output := 0
	currentDialPosition := DEFUALT_DIAL_LOCATION
	for rotation := range input {
		// We are basically rotating the dial at this point
		newPosition := positionTheDial(currentDialPosition, processInput(input[rotation]))
		if newPosition == 0 {
			output++
		}
		currentDialPosition = newPosition
	}
	fmt.Println("Final Output :", output)
}

// L39 -> Move Left towards 0
// R29 -> Move right towards 99
func processInput(rotationStep string) int {
	// Get the first charcacter that defines if the input is +ve or -ve
	// Then rest of the string is just integer
	isPositive := false
	if rotationStep[0] == 'R' {
		isPositive = true
	}
	rotationCount := rotationStep[1:]
	result, err := strconv.Atoi(rotationCount)
	// The steps might be more than 100 so ignore a complete roation
	if result > 100 {
		result = result % 100
	}
	if err != nil {
		// I am not sure what to do here
	}
	if isPositive {
		return result
	} else {
		return -result
	}
}

// if the input forces to go below 0 or over 99 handle that here, so just position the dial correctly
func positionTheDial(currentDialPosition int, dialChange int) int {
	// If going below 0 subtract the rest from 99
	// If going above 99 add the rest from 0
	// Base case is between 0 - 99
	result := currentDialPosition + dialChange
	if result == 100 { // This is basically zero
		return 0
	}
	if result >= LOWEST_NUM_ON_DIAL && result <= HIGHEST_NUM_ON_DIAL {
		return result
	} else {
		if result < LOWEST_NUM_ON_DIAL {
			return HIGHEST_NUM_ON_DIAL + (result + 1) // + 1 To account for 0
		} else {
			return LOWEST_NUM_ON_DIAL + (result - (HIGHEST_NUM_ON_DIAL + 1))
		}
	}
}

// Helper function to read input from the input.txt
func getInputFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			lines = append(lines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}
