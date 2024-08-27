package printbitswithspace

import (
	"fmt"
)

func PrintBitsWithSpace(boolArray []bool, spaces []int, maxBits int) {
	for i, b := range boolArray {
		if b {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}

		if min(len(spaces), maxBits) > 0 && spaces[0] == i {
			spaces = spaces[1:]
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
