package printbitswithspace

import (
	"fmt"
	"strings"
)

func FormatBitsWithSpace(boolArray []bool, spaces []int, maxBits int) string {
	var result strings.Builder

	for i, b := range boolArray {
		if min(len(spaces), maxBits-i) > 0 {
			if (spaces[0] == i) {
				spaces = spaces[1:]
				result.WriteString(" ")
			}
		}
		if b {
			result.WriteString("1")
		} else {
			result.WriteString("0")
		}
	}

	return result.String()
}

func PrintBitsWithSpace(boolArray []bool, spaces []int, maxBits int) {
	fmt.Println(FormatBitsWithSpace(boolArray, spaces, maxBits))
}
