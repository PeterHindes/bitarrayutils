package bit-array-utils/file_utils

import (
    "io"
    "os"
)

func SaveBinaryFile(array []bool, filename string) error {
    // Create a new file
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    // Save the array to the file
    for i := 0; i < len(array); i += 8 {
        microarray := make([]byte, 1)
        for j := 0; j < 8 && i+j < len(array); j++ {
            if array[i+j] {
                microarray[0] |= 1 << uint8(7-j)
            }
        }
        _, err := file.Write(microarray)
        return err
    }
    return nil
}

func LoadBinaryFile(filename string) ([]bool, error) {
    // Open the file
	file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    // Read the file
    bitArray := make([]bool, 0)
    for {
        byteArray := make([]byte, 1)
        _, err := file.Read(byteArray)
        if err != nil {
            if err == io.EOF {
                break
            }
            return nil, err
        }

        for j := 0; j < 8; j++ {
            bit := (byteArray[0] >> uint8(7-j)) & 1
            bitArray = append(bitArray, bit == 1)
        }
    }
    return bitArray, nil
}