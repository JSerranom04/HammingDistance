// Package main implements a program to generate binary strings with a specific Hamming distance
package main

import (
	"fmt"
	"sync"
)

// generateHammingDistance generates all possible binary strings of length maxLong
// with exactly hammingDistance ones (Hamming distance from the zero string)
// Returns a slice of boolean slices representing the binary strings
func generateHammingDistance(maxLong, hammingDistance uint) [][]bool {
	if hammingDistance > maxLong {
		return nil
	}
	
	binaryChain := make([]bool, maxLong)
	result := make([][]bool, 0)
	var wg sync.WaitGroup
	var mu sync.Mutex
	
	wg.Add(1)
	hammingDistanceRec(binaryChain, 0, maxLong, hammingDistance, &result, &wg, &mu)
	wg.Wait()
	
	return result
}

func hammingDistanceRec(binaryChain []bool, pos, maxLong, remainingOnes uint, result *[][]bool, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	
	// Early termination: impossible to place remaining ones in available positions
	if remainingOnes > (maxLong - pos) {
		return
	}
	
	// Make recursive calls first
	if pos < maxLong {
		// Try with 0
		binaryChain[pos] = false
		wg.Add(1)
		go hammingDistanceRec(binaryChain, pos+1, maxLong, remainingOnes, result, wg, mu)
		
		// Try with 1 if we still need more ones
		if remainingOnes > 0 {
			binaryChain[pos] = true
			wg.Add(1)
			go hammingDistanceRec(binaryChain, pos+1, maxLong, remainingOnes-1, result, wg, mu)
		}
		return
	}
	
	// Check if this is a valid solution after reaching the end
	if pos == maxLong && remainingOnes == 0 {
		chainCopy := make([]bool, maxLong)
		copy(chainCopy, binaryChain)
		mu.Lock()
		*result = append(*result, chainCopy)
		mu.Unlock()
	}
}

func main() {
	var maxLong, hammingDistance uint
	
	// Ask for the length of the binary strings
	fmt.Print("Enter the length of the binary strings: ")
	_, err := fmt.Scanf("%d", &maxLong)
	if err != nil {
		fmt.Println("Error: Please enter a valid positive number")
		return
	}
	
	// Ask for the Hamming distance
	fmt.Print("Enter the desired Hamming distance: ")
	_, err = fmt.Scanf("%d", &hammingDistance)
	if err != nil {
		fmt.Println("Error: Please enter a valid positive number")
		return
	}
	
	// Validate input
	if hammingDistance > maxLong {
		fmt.Printf("Error: Hamming distance (%d) cannot be greater than string length (%d)\n", 
			hammingDistance, maxLong)
		return
	}
	
	// Generate and print results
	result := generateHammingDistance(maxLong, hammingDistance)
	
	if len(result) == 0 {
		fmt.Println("No valid binary strings found for the given parameters")
		return
	}
	
	fmt.Printf("\nFound %d binary strings:\n", len(result))
	for i, chain := range result {
		fmt.Printf("Chain %d: ", i+1)
		for _, bit := range chain {
			if bit {
				fmt.Print("1")
			} else {
				fmt.Print("0")
			}
		}
		fmt.Println()
	}
}