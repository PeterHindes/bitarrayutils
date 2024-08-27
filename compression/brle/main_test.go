package brle

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/PeterHindes/bitarrayutils/debug/printbitswithspace"
	"github.com/PeterHindes/bitarrayutils/debug/printbrle"
)

func TestBlankRunEncode(t *testing.T) {
	tests := []struct {
		input            []bool
		powerOfTwo       int
		expectedEncoding []bool
	}{
		{[]bool{false, false, true, false, false, false, true},
			2,
			[]bool{
				// bits for extracted size = 3
				false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, true,
				true, true, true, // extracted size 7
				false, false, false, true, false, // Runlength 2
				true, false, true, // 2 then true
				true, true, true, // 3 then true
			},
		},
		// {[]bool{true, false, false, false, true, false, false}, 3, 69},
		// {[]bool{false, false, false, false, false, false, false}, 4, 66},
	}

	for _, test := range tests {
		encoded := BlankRunEncode(test.input, test.powerOfTwo)
		for i := 0; i < len(test.input); i++ {
			if encoded[i] != test.expectedEncoding[i] {
				t.Errorf("\nExpected %s,\n     got %s", printbrle.FormatBRLE(test.expectedEncoding, 300), printbrle.FormatBRLE(encoded, 300))
				break
			}
		}
	}
}

func TestBlankRunDecode(t *testing.T) {
	tests := []struct {
		input   []bool
		baseTwo int
	}{
		// {
		//     []bool{false,true,false,false,true,true,true,false,true,true,true,false,true,true,true,true,false,false,true,false},
		//     4,
		// },
		// {
		//     []bool{true, false, true, true, false, false, true, true, false, true, false, true},
		//     2,
		// },
		{
			[]bool{false, false, true, false, false, false, true},
			2,
		},
		{
			[]bool{false, false, true, false, false, false, false},
			4,
		},
		{
			[]bool{false, false, true, false, false, false, false, false, false, false, false, true, false},
			3,
		},
	}

	for _, test := range tests {
		encoded := BlankRunEncode(test.input, test.baseTwo)
		decoded := BlankRunDecode(encoded)
		for i := 0; i < len(test.input); i++ {
			if decoded[i] != test.input[i] {
				t.Errorf("BlankRunDecode:\nexp %v\ngot %v\n\nEncoding: %v", printbitswithspace.FormatBitsWithSpace(test.input, []int{}, 300), printbitswithspace.FormatBitsWithSpace(decoded, []int{}, 300), printbrle.FormatBRLE(encoded, 300))
				break
			}
		}
	}
}

func generateRandomBoolArray(length int) []bool {
	array := make([]bool, length)
	for i := 0; i < length/2; i++ {
		array[rand.Intn(length)] = true
	}

	return array
}

func generateRandomPowerOfTwo(min, max int) int {
	return rand.Intn(max-min) + min
}

func TestRandomBlankRunEncode(t *testing.T) { // TODO some sort of memory leak is here
	// Define the range and number of random tests
	numTests := 100000

	for i := 0; i < numTests; i++ {
		// Generate a random input array
		length := rand.Intn(1000)+1
		input := generateRandomBoolArray(length)

		// t.Log("Done generating random bool array")

		// Generate a random power of two min 2 max 32
		powerOfTwo := generateRandomPowerOfTwo(2, 32)
		// powerOfTwo := generateRandomPowerOfTwo(2, 20)

        defer func() {
            // recover from panic if one occurred. Set err to nil otherwise.
            if recover() != nil {
                fmt.Println("params:", length, powerOfTwo)
            }
        }()

		// Encode the input array
		encoded := BlankRunEncode(input, powerOfTwo)

		// t.Log("Done encoding")

		// Decode the encoded array
		decoded := BlankRunDecode(encoded)

		// t.Log("Done decoding")

		// Compare the decoded array with the original input array
		for j := 0; j < len(input); j++ {
			if decoded[j] != input[j] {

				t.Errorf("\nFaild At Index: %d\npowerOfTwo: %d\nexp %v\ngot %v\n\nEncoding: %v",
					j,
					powerOfTwo,
					printbitswithspace.FormatBitsWithSpace(input, []int{}, 100),
					printbitswithspace.FormatBitsWithSpace(decoded, []int{}, 100),
					printbrle.FormatBRLE(encoded, 300))
				// t.Errorf("Randoms Failed: %d missmatched",j)
				return
			}
		}
	}

    fmt.Printf("finished %d\n", numTests)
}
