package brle

import (
	"math"
)

func BlankRunEncode(array []bool, powerOfTwo int) []bool {
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

	// use bitLenOfArraySize bits to store the arraySize
	arraySize := len(array)
	arraySizeBits := make([]bool, 0, int(bitLenOfArraySize))
	for j := int(bitLenOfArraySize) - 1; j >= 0; j-- {
		arraySizeBits = append(arraySizeBits, (arraySize>>uint(j))&1 == 1)
	}
	encoded = append(encoded, arraySizeBits...)

	// convert the power of two to a boolean array and append them to the encoded array
	// 5 bits
	powerOfTwoBits := make([]bool, 0, 5)
	for j := 4; j >= 0; j-- {
		powerOfTwoBits = append(powerOfTwoBits, (powerOfTwo>>j)&1 == 1)
	}
	encoded = append(encoded, powerOfTwoBits...)

	// loop thorugh the input array to find zeros
	run := 0
	for i := 0; i < len(array); i++ {
		// if the next one is false and run is less than powerOfTwo bits long increment run
		if (!array[i] && run < int(math.Pow(2, float64(powerOfTwo)))-1 && i != len(array)-1) {
			run++
		} else {
			// convert the run to a boolean array and append them to the encoded array
			add := []bool{}
			for j := powerOfTwo - 1; j >= 0; j-- {
				add = append(add, (run>>j)&1 == 1)
			}
			encoded = append(encoded, add...)
			encoded = append(encoded, array[i])
			run = 0
		}
	}

	return encoded

}

func BlankRunDecode(array []bool) []bool {
	// Initiate i, which we will increment as we move through the array
	i := 0

	// Get the bit length of arraySize by reading the first 53 bits last=msb
	bitLenOfArraySize := 0
	for j := 52; j >= 0; j-- {
		if array[i] {
			bitLenOfArraySize += 1 << uint(j)
		}
		i++
	}

	// Get the size of the array
	arraySize := 0
	for j := 0; j < bitLenOfArraySize; j++ {
		if array[i] {
			arraySize += 1 << uint(bitLenOfArraySize-j-1)
		}
		i++
	}

	// Get the number of bits in a run
	bitsPerRun := 0
	for j := 0; j < 5; j++ {
		if array[i] {
			bitsPerRun += 1 << uint(4-j)
		}
		i++
	}

	// Loop through the array to decode the data
	// Create a new array for the decoded data
	decoded := make([]bool, 0, arraySize)
	for i < len(array) {

		// Get the run
		run := 0
		for j := 0; j < bitsPerRun; j++ {
			if array[i] {
				run += (1 << uint(bitsPerRun-j-1))
			}
			i++
		}

		// Get the end value
		end := array[i]
		i++

		// Add the run of zeros to the decoded array
		for j := 0; j < run; j++ {
			decoded = append(decoded, false)
		}

		// Add the end value to the decoded array
		decoded = append(decoded, end)
	}

	return decoded
}