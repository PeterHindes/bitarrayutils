package brle

import (
	"math/rand"
	"testing"

	// "github.com/PeterHindes/bitarrayutils/debug/printbitswithspace"
	"github.com/PeterHindes/bitarrayutils/debug/printbrle"
)

func TestBlankRunEncode(t *testing.T) {
    tests := []struct {
        input       []bool
        powerOfTwo  int
        expectedEncoding []bool
    }{
        {[]bool{false, false, true, false, false, false, true}, 
            2, 
            []bool{
                // bits for extracted size = 3
                false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, true,
                true,true,true, // extracted size 7
                false, false, false, true, false, // Runlength 2
                true,false, true, // 2 then true
                true,true, true, // 3 then true
            },
        },
        // {[]bool{true, false, false, false, true, false, false}, 3, 69},
        // {[]bool{false, false, false, false, false, false, false}, 4, 66},
    }

    for _, test := range tests {
        encoded := StepBlankRunEncode(test.input, test.powerOfTwo)
        for i := 0; i < len(test.input); i++ {
            if encoded[i] != test.expectedEncoding[i] {
                // t.Errorf("got %d, expected length %d\nPowerof2: %d\nArrayLen: %d\nArray: %s\nEncode: %s", len(encoded), test.expectedLen, test.powerOfTwo, len(test.input), printbitswithspace.FormatBitsWithSpace(test.input, []int{}, 300), printbrle.FormatBRLE(encoded, 300))
                t.Errorf("\nExpected %s,\n     got %s", printbrle.FormatBRLE(test.expectedEncoding, 300), printbrle.FormatBRLE(encoded, 300))
                break
            }
        }
    }
}

// func TestBlankRunDecode(t *testing.T) {
//     tests := []struct {
//         input    []bool
//         expected []bool
//     }{
//         {[]bool{false, false, true, false, false, false, true}, []bool{false, false, true, false, false, false, true}},
//         {[]bool{true, false, false, false, true, false, false}, []bool{true, false, false, false, true, false, false}},
//         {[]bool{false, false, false, false, false, false, false}, []bool{false, false, false, false, false, false, false}},
//     }

//     for _, test := range tests {
//         decoded := BlankRunDecode(test.input)
//         if len(decoded) != len(test.expected) {
//             t.Errorf("BlankRunDecode(%v) = %v, expected %v", test.input, decoded, test.expected)
//         }
//         for i := range decoded {
//             if decoded[i] != test.expected[i] {
//                 t.Errorf("BlankRunDecode(%v) = %v, expected %v", test.input, decoded, test.expected)
//                 break
//             }
//         }
//     }
// }

// func TestEncodeDecodeConsistency(t *testing.T) {
//     rand.Seed(time.Now().UnixNano())

//     tests := []struct {
//         input      []bool
//         powerOfTwo int
//     }{
//         {generateRandomBoolArray(100), 2},
//         {generateRandomBoolArray(200), 3},
//         {generateRandomBoolArray(300), 4},
//     }

//     for _, test := range tests {
//         encoded := BlankRunEncode(test.input, test.powerOfTwo)
//         decoded := BlankRunDecode(encoded)
//         if len(decoded) != len(test.input) {
//             t.Errorf("EncodeDecodeConsistency: length mismatch for input %v, got %v", test.input, decoded)
//         }
//         for i := range decoded {
//             if decoded[i] != test.input[i] {
//                 t.Errorf("EncodeDecodeConsistency: value mismatch at index %d for input %v, got %v", i, test.input, decoded)
//                 break
//             }
//         }
//     }
// }

func generateRandomBoolArray(size int) []bool {
    array := make([]bool, size)
    for i := range array {
        array[i] = rand.Intn(2) == 1
    }
    return array
}