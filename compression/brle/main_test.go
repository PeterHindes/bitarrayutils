package brle

import (
	"testing"
    "math/rand"
	// "fmt"
	// "time"
)

func TestBlankRunEncode(t *testing.T) {
    tests := []struct {
        input       []bool
        powerOfTwo  int
        expectedLen int
    }{
        {[]bool{false, false, true, false, false, false, true}, 2, 67},
        {[]bool{true, false, false, false, true, false, false}, 3, 69},
        {[]bool{false, false, false, false, false, false, false}, 4, 66},
    }

    for _, test := range tests {
        encoded := BlankRunEncode(test.input, test.powerOfTwo)
        if len(encoded) != test.expectedLen {
            t.Errorf("BlankRunEncode(%v, %d) = %v, expected length %d", test.input, test.powerOfTwo, encoded, test.expectedLen)
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