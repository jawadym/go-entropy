package main

import (
	"fmt"
	"os"
)

const (
	chunksize int = 1024
)

var (
	data  *os.File
	part  []byte
	err   error
	count int
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./entropy [path]file")
		os.Exit(-1)
	}

	var freqList []float64
	fileSize, bytesFile := openFile(os.Args[1])

	// construct frequencies array
	for b := 0; b < 256; b++ {
		ctr := 0
		for _, byteChunk := range bytesFile.Bytes() {
			if int(byteChunk) == b {
				ctr++
			}
		}
		freqList = append(freqList, float64(ctr)/float64(fileSize))

	}

	fmt.Println(entropySum(freqList))

}
