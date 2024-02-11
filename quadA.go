package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func QuadA(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if i == 0 && j == 0 || i == 0 && j == x-1 || i == y-1 && j == 0 || i == y-1 && j == x-1 {
				z01.PrintRune('o')
			} else if i == 0 || i == y-1 {
				z01.PrintRune('-')
			} else if j == 0 || j == x-1 {
				z01.PrintRune('|')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Error: length of Arguments is not 3\n")
		return
	}

	width, err := strconv.Atoi(os.Args[1])
	if err != nil || width <= 0 {
		fmt.Println("Not a quad function")
		return
	}

	height, err := strconv.Atoi(os.Args[2])
	if err != nil || height <= 0 {
		fmt.Println("Not a quad function")
		return
	}
	QuadA(width, height)
}
