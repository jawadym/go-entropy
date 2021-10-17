package main

import "math"

// Takes a list of byte frequencies, and calculates Entropy using Shannon's formula.
func entropySum(freqList []float64) float64 {
	var entropy float64
	for freqIndex := range freqList {
		if freqList[freqIndex] > 0 {
			entropy = entropy + freqList[freqIndex]*(math.Log2(freqList[freqIndex]))
		}
	}
	return -entropy
}
