package printbrle

import (
	"github.com/PeterHindes/bitarrayutils/debug/printbitswithspace"
)

func PrintBRLE(brleEncodedArray []bool) {
	spaces := []int{53, 56, 61}

	// Extract the runlength bit count from 56 to 61
	runLengthBitCount := 0
	for i := 56; i <= 61; i++ {
		if brleEncodedArray[i] {
			runLengthBitCount += 1 << uint(61-i)
		}
	}

	// add spaces after 61+ runlength bit count +1 cyclicly
	for i := 62 + runLengthBitCount; i < len(brleEncodedArray); i += runLengthBitCount + 1 {
		spaces = append(spaces, i)
	}

	printbitswithspace.PrintBitsWithSpace(brleEncodedArray, spaces, 300)
}