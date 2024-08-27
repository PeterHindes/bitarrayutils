package printbrle

import (
	"fmt"

	"github.com/PeterHindes/bitarrayutils/debug/printbitswithspace"
)

func PrintBRLE(brleEncodedArray []bool, maxBits int) {
	spaces := brleSpaces(brleEncodedArray)
	printbitswithspace.PrintBitsWithSpace(brleEncodedArray, spaces, maxBits)
}

func FormatBRLE(brleEncodedArray []bool, maxBits int) string {
	spaces := brleSpaces(brleEncodedArray)
	return printbitswithspace.FormatBitsWithSpace(brleEncodedArray, spaces, maxBits)
}

func brleSpaces(brleEncodedArray []bool) []int {
	spaces := []int{}

	i := 0

	// Extract the find the sizeofextractedfilesizebitcount from 0 to 52
	sizeofExtractedFileSizeBitCount := 0
	for j := 0; j <= 52 && i < len(brleEncodedArray); j++ {
		if brleEncodedArray[i] {
			sizeofExtractedFileSizeBitCount += 1 << uint(52-i)
		}
		i++
	}

	spaces = append(spaces, i)

	i += int(sizeofExtractedFileSizeBitCount)

	spaces = append(spaces, i)

	// Extract the runlength bit count from 53+sizeofextractedfilesizebitcount to 53+sizeofextractedfilesizebitcount+5
	runLengthBitCount := 0
	for j := 0; j < 5 && i < len(brleEncodedArray); j++ {
		if brleEncodedArray[i] {
			runLengthBitCount += 1 << uint(4-j)
		}
		i++
	}

	// add spaces after every rulength bit count+1 bits
	for j := 0; j < runLengthBitCount+1 && i < len(brleEncodedArray); j++ {
		spaces = append(spaces, i)
		i += runLengthBitCount + 1
	}

	fmt.Println("Spaces:", spaces)

	return spaces
}

// func brleSpaces(brleEncodedArray []bool, powerOfTwo int) []int {
// 	spaces := []int{53}