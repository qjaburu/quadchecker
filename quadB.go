package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func QuadB(x, y int) {
	if x > 0 && y > 0 {
		for i := 0; i < y; i++ {
			if i == 0 {
				z01.PrintRune('/')
			}
			if i == y-1 && y-1 != 0 {
				z01.PrintRune('\\')
			}
			if i == 0 {
				for j := 0; j < x-1; j++ {
					if j == x-2 {
						z01.PrintRune('\\')
					} else {
						z01.PrintRune('*')
					}
				}
				z01.PrintRune('\n')
			}
			if i == y-1 && y-1 != 0 {
				for j := 0; j < x-1; j++ {
					if j == x-2 {
						z01.PrintRune('/')
					} else {
						z01.PrintRune('*')
					}
				}
				z01.PrintRune('\n')
			}
			if i < y-1 && i > 0 {
				for j := 0; j < x; j++ {
					if j == x-1 || j == 0 {
						z01.PrintRune('*')
					} else {
						z01.PrintRune(' ')
					}
				}
				z01.PrintRune('\n')
			}
		}
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
	QuadB(width, height)
}
