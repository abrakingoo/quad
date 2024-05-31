package functions

import (
	"fmt"
)

func QuadA(x, y int) {
	if x > 0 && y > 0 {
		fmt.Print("o")
		if x > 1 {
			for i := 1; i < x - 1; i++ {
				fmt.Print("-")
			}
			fmt.Println("o")
		} else {
			fmt.Println()
		}

		for i := 0; i < y-2; i++ {
			fmt.Print("|")
			if x > 1 {
				for j := 0; j < x-2; j++ {
					fmt.Print(" ")
				}
				fmt.Println("|")
			} else {
				fmt.Println()
			}
		}

		if y > 1 {
			fmt.Print("o")
			if x > 1 {
				for i := 1; i < x-1; i++ {
					fmt.Print("-")
				}
				fmt.Println("o")

			} else {
				fmt.Println()
			}
		}
	}

}