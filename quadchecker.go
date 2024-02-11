package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Initialize an empty rune slice to store characters
	var arr []rune

	// Create a reader to read characters from standard input
	reader := bufio.NewReader(os.Stdin)

	// Read characters until an error occurs (e.g., end of input)
	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		arr = append(arr, char)
	}

	// Convert rune slice to string for easier handling
	input := string(arr)

	// Count the width and height of the shape
	width := 0
	height := 0
	for i := 0; i < len(arr); i++ {
		// Check for characters other than newline and increment width if height is still 0
		if arr[i] != '\n' && height == 0 {
			width++
		}
		// Check for newline characters and increment height
		if arr[i] == '\n' {
			height++
		}
	}

	// Check if the dimensions are invalid
	if width == 0 || height == 0 {
		fmt.Println("Not a quad function")
		return
	}

	// Check for various quad patterns and print corresponding information
	if isQd(input, width, height, 'o', 'o', 'o', 'o', '-', '|') {
		fmt.Printf("[quadA] [%v] [%v]\n", width, height)
		return
	}

	if isQd(input, width, height, '/', '\\', '\\', '/', '*', '*') {
		fmt.Printf("[quadB] [%v] [%v]\n", width, height)
		return
	}

	var count int = 0
	if isQd(input, width, height, 'A', 'A', 'C', 'C', 'B', 'B') {
		count++
		fmt.Printf("[quadC] [%v] [%v]", width, height)
	}

	if isQd(input, width, height, 'A', 'C', 'A', 'C', 'B', 'B') {
		if count > 0 {
			fmt.Print(" || ")
		}
		count++
		fmt.Printf("[quadD] [%v] [%v]", width, height)
	}

	if isQd(input, width, height, 'A', 'C', 'C', 'A', 'B', 'B') {
		if count > 0 {
			fmt.Print(" || ")
		}
		count++
		fmt.Printf("[quadE] [%v] [%v]", width, height)
	}

	// Print a newline if at least one quad was matched
	if count > 0 {
		fmt.Println()
		return
	}

	// Print an error message if no recognized quad pattern was matched
	fmt.Println("Not a quad function")
}

// Function to check if the input shape matches a specific quad pattern
func isQd(input string, width, height int, topleft, topright, bottomleft, bottomright, hr, vat rune) bool {
	var arrq []rune

	// Iterate through the height and width of the shape
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			// Determine the character based on the position in the quad
			switch i {
			case 0:
				switch j {
				case 0:
					arrq = append(arrq, topleft)
				case width - 1:
					arrq = append(arrq, topright)
				default:
					arrq = append(arrq, hr)
				}
			case height - 1:
				switch j {
				case 0:
					arrq = append(arrq, bottomleft)
				case width - 1:
					arrq = append(arrq, bottomright)
				default:
					arrq = append(arrq, hr)
				}
			default:
				switch j {
				case 0, width - 1:
					arrq = append(arrq, vat)
				default:
					arrq = append(arrq, ' ')
				}
			}
		}
		// Add newline character to separate rows
		arrq = append(arrq, '\n')
	}

	// Convert rune slice to string for easier comparison
	output := string(arrq)

	// Check if the generated quad matches the input shape
	return output == input
}
