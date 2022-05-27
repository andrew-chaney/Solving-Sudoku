package utils

import (
	"encoding/csv"
	"fmt"
	"os"
)

func Read(path string) []string {
	puzzles := []string{}

	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error: Couldn't open file - %s\n", path)
		os.Exit(0)
	}

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Printf("Error: Couldn't read csv vals from file - %s\n", path)
		os.Exit(0)
	}
	for i := 0; i < len(lines); i++ {
		// 0: unsolved puzzle, 1: solution
		if i != 0 {
			puzzles = append(puzzles, lines[i][0])
		}
	}
	return puzzles
}