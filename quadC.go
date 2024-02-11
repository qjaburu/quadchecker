package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func QuadC(x, y int) {
	if x > 0 && y > 0 {
		if x == 1 && y == 1 {
			z01.PrintRune('A')
			z01.PrintRune('\n')
			return
		} else if x == 1 {
			z01.PrintRune('A')
			for s := 1; s < y-1; s++ {
				if s <= 0 {
					break
				} else {
					z01.PrintRune('\n')
					z01.PrintRune('B')
				}
				if y == 1 {
					z01.PrintRune('\n')
					z01.PrintRune('A')
					return
				}
			}
			z01.PrintRune('\n')
			z01.PrintRune('C')
			z01.PrintRune('\n')
		} else if y == 1 {
			z01.PrintRune('A')
			for i := 1; i < x-1; i++ {
				z01.PrintRune('B')
			}
			z01.PrintRune('A')
			z01.PrintRune('\n')
			return
		} else {
			z01.PrintRune('A')
			for i := 1; i < x-1; i++ {
				z01.PrintRune('B')
			}
			z01.PrintRune('A')
			z01.PrintRune('\n')
			for s := 1; s < y-1; s++ {
				z01.PrintRune('B')
				for j := 1; j < x-1; j++ {
					z01.PrintRune(' ')
				}
				z01.PrintRune('B')
				z01.PrintRune('\n')
			}
			z01.PrintRune('C')
			for i := 1; i < x-1; i++ {
				z01.PrintRune('B')
			}
			z01.PrintRune('C')
			z01.PrintRune('\n')
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
	QuadC(width, height)
}
