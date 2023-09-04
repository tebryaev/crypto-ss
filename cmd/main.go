package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
)

func crypt(input string) []int {
	hash := sha256.Sum256([]byte(input))
	sequence := make([]int, 24)
	for i := range sequence {
		sequence[i] = i + 1
	}

	result := make([]int, 24)
	for i := range hash {
		if len(sequence) == 0 {
			break
		}
		idx := int(hash[i]) % len(sequence)
		result[i] = sequence[idx]
		sequence = append(sequence[:idx], sequence[idx+1:]...)
	}

	return result
}

func main() {
	fmt.Print("Please enter the password string: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	sequence := crypt(input)
	for i, val := range sequence {
		fmt.Printf("%2d -> %2d\n", i+1, val)
	}
}
