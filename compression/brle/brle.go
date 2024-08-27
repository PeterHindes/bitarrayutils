package brle

import (
	"fmt"
	"math"
)

// Custom RLE encoding function
func BlankRunEncode(array []bool, powerOfTwo int) []bool {
	arraySize := len(array)
	// Create a new array for the runs
	runs := make([]int, 0)
	// Create a new array for the end values
	ends := make([]bool, 0)

	// keep track of our current run
	run := 0
	// loop through the input array to find zeros
	for i := 0; i < arraySize; i++ {
		if (array[i] == false) {
			run++
		} else {
			runs = append(runs, run)
			run = 0
			ends = append(ends, array[i])
		}
	}
	// Handle the case where the last run is a zero run
	if run > 0 {
		runs = append(runs, run-1)
		ends = append(ends, false)
	}

	powerInt := int(math.Pow(2, float64(powerOfTwo)))

	// Split runs longer than nearest power of two into multiple runs (the first of witch must have and end of false and the second stays true)
	// This requires looping through the runs array and checking if the run is longer than the power of two
	// if it is then we will split it into two runs
	newRuns := make([]int, 0)
	newEnds := make([]bool, 0)
	for i := 0; i < len(runs); i++ {
		newestRuns, newestEnds := splitByPowerOfTwo(runs[i], powerInt, ends[i])
		newRuns = append(newRuns, newestRuns...)
		newEnds = append(newEnds, newestEnds...)
	}


	// From below just encodes using the powerOfTwo and the newRuns and newEnds arrays

	// Create a new array for the change encoding
	// first 5 bits used to store the power of two we are using for max run length
	// then the data payload
	// the data payload consists of a number represented by (power of two) bits followed by a bit to indicate of that run ends with a true or false
	// the data payload is repeated until the end of the array that it represents
	encoded := make([]bool, 0)

	// find how many bits are needed to store arraySize
	bitLenOfArraySize := math.Ceil(math.Log2(float64(arraySize)))

	// if it is more than 53 bits (can store 1petabyte) then we will panic
	if bitLenOfArraySize > 53 {
		panic("Array size is too large")
	}

	// use 52 bits to store the bitLenOfArraySize
	for i := 51; i >= 0; i-- {
		encoded = append(encoded, (int(bitLenOfArraySize)>>i)&1 == 1)
	}

	// use bitLenOfArraySize bits to store the arraySize
	for i := int(bitLenOfArraySize) - 1; i >= 0; i-- {
		encoded = append(encoded, (arraySize>>uint(i))&1 == 1)
	}

	// convert the power of two to a boolean array and append them to the encoded array
	// 5 bits
	for j := 4; j >= 0; j-- {
		encoded = append(encoded, (powerOfTwo>>j)&1 == 1)
	}

	// loop through the alligned newRuns and newEnds arrays to encode the data
	for i := 0; i < len(newRuns); i++ {
	// for i := 0; i < 3; i++ {
		// convert the run to a boolean array and append them to the encoded array
		for j := powerOfTwo - 1; j >= 0; j-- {
			encoded = append(encoded, (newRuns[i]>>j)&1 == 1)
		}

		// then insert the current value of the array
		encoded = append(encoded, newEnds[i])
	}

	// Print the size diffrence between the two arrays
	// fmt.Println()
	// fmt.Println("Original array size:", arraySize)
	// fmt.Println("Encoded array size:  ", len(encoded))


	

	return encoded

}

func StepBlankRunEncode(array []bool, powerOfTwo int) []bool {
	encoded := make([]bool, 0) // TODO predict this to save memory allocation time

	// find how many bits are needed to store arraySize
	bitLenOfArraySize := math.Ceil(math.Log2(float64(len(array))))

	// if it is more than 53 bits (can store 1petabyte) then we will panic
	if bitLenOfArraySize > 53 {
		panic("Array size is too large")
	}

	// use 53 bits to store the bitLenOfArraySize
	bitLenOfArraySizeBits := make([]bool, 0)
	for j := 52; j >= 0; j-- {
		bitLenOfArraySizeBits = append(bitLenOfArraySizeBits, (int(bitLenOfArraySize)>>j)&1 == 1)
	}
	encoded = append(encoded, bitLenOfArraySizeBits...)

	fmt.Println("Sec1:", bitLenOfArraySizeBits)
	fmt.Println("Sec1 len:", len(bitLenOfArraySizeBits))

	// use bitLenOfArraySize bits to store the arraySize
	arraySize := len(array)
	fmt.Printf("Array size: %d\n", arraySize)
	arraySizeBits := make([]bool, 0, int(bitLenOfArraySize))
	for j := int(bitLenOfArraySize) - 1; j >= 0; j-- {
		arraySizeBits = append(arraySizeBits, (arraySize>>uint(j))&1 == 1)
	}
	encoded = append(encoded, arraySizeBits...)

	fmt.Println("Sec2:", arraySizeBits)

	// convert the power of two to a boolean array and append them to the encoded array
	// 5 bits
	powerOfTwoBits := make([]bool, 0, 5)
	for j := 4; j >= 0; j-- {
		powerOfTwoBits = append(powerOfTwoBits, (powerOfTwo>>j)&1 == 1)
	}
	encoded = append(encoded, powerOfTwoBits...)

	fmt.Println("Sec3:", powerOfTwoBits)

	// loop thorugh the input array to find zeros
	run := 0
	for i := 0; i < len(array); i++ {
		// if the next one is false and run is less than powerOfTwo bits long increment run
		if (!array[i] && run < int(math.Pow(2, float64(powerOfTwo)))) {
			run++
		} else {
			// convert the run to a boolean array and append them to the encoded array
			add := []bool{}
			for j := powerOfTwo - 1; j >= 0; j-- {
				add = append(add, (run>>j)&1 == 1)
			}
			fmt.Println("Add:", add)
			encoded = append(encoded, add...)
			encoded = append(encoded, array[i])
			run = 0
		}
	}

	return encoded

}

// Function to split a run into multiple runs
func splitByPowerOfTwo(run, power int, end bool) ([]int, []bool) {
	// Check if the run is smaller than the power of two
	if run <= power {
		return []int{run}, []bool{end}
	}

	// Split the run into two
	runs := make([]int, 0)
	ends := make([]bool, 0)
	runs = append(runs, power-1)
	ends = append(ends, false)
	newRuns, newEnds := splitByPowerOfTwo(run-power+1, power, end)
	runs = append(runs, newRuns...)
	ends = append(ends, newEnds...)

	return runs, ends
}

func BlankRunDecode(array []bool) []bool {

	// Initiate i, which we will increment as we move through the array
	i := 0

	// Get the bit length of arraySize
	bitLenOfArraySize := 0
	for j := 51; j >= 0; j-- {
		if array[i] {
			bitLenOfArraySize += int(math.Pow(2, float64(51-j)))
		}
		i++
	}

	// Get the size of the array
	arraySize := 0
	for j := 0; j < bitLenOfArraySize; j++ {
		if array[i] {
			arraySize += int(math.Pow(2, float64(51-j)))
		}
		i++
	}

	// Create a new array for the decoded data
	decoded := make([]bool, 0, arraySize)

	// Get the power of two used to encode the data
	powerOfTwo := 0
	for j := 4; j >= 0; j-- {
		if array[i] {
			powerOfTwo += int(math.Pow(2, float64(4-j)))
		}
		i++
	}

	// Loop through the array to decode the data
	for i < len(array) {
		// Get the run
		run := 0
		for j := 0; j < powerOfTwo; j++ {
			if array[i] {
				run += int(math.Pow(2, float64(powerOfTwo-j-1)))
			}
			i++
		}

		// Get the end value
		end := array[i]
		i++

		// Add the run of zeros to the decoded array
		for j := 0; j <= run; j++ {
			decoded = append(decoded, false)
		}

		// Add the end value to the decoded array
		decoded = append(decoded, end)
	}

	return decoded
}