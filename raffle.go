package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func doRaffle(filePath string, seed int) error {
	// Read the file contents
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Split the file contents into lines
	lines := strings.Split(string(data), "\n")

	// Initialize the random source with the provided seed
	rand.Seed(int64(seed))

	// Shuffle the lines
	rand.Shuffle(len(lines), func(i, j int) {
		lines[i], lines[j] = lines[j], lines[i]
	})

	// Output all of the lines
	for _, line := range lines {
		fmt.Println(line)
	}

	return nil
}
