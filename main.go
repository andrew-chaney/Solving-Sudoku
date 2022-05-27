package main

import (
	"fmt"
	//"math/rand"
	//"time"

	"github.com/andrew-chaney/Solving-Sudoku/src/ai"
	//"github.com/andrew-chaney/Solving-Sudoku/src/bruteforce"
	"github.com/andrew-chaney/Solving-Sudoku/src/utils"
)

/* TODO:
- Sudoku Board Display................................DONE
- Sudoku Board Builder................................DONE
- CSV Parsing.........................................DONE
- Solution verification...............................DONE
- Implement backtracking solution.....................DONE
------------ Wikipedia Page for Ref ----------------------
- Implement Simulated Annealing.......................
- Implement Genetic Algorithm solution................
- Implement Constraint solution.......................
- Implement Exact Cover solution......................
- Implement Crook's Algorithm solution................
----------------------------------------------------------
- Status bar to show progress in data collection......
- Implement graphing..................................
- Minimax?............................................
- AI?.................................................
*/

func main() {
	fmt.Println("Puzzle Sample Sizes")
	fmt.Println("1. 1 million")
	fmt.Println("2. 9 million")
	var selection string
	for {
		fmt.Print("Enter # corresponding to desired sample size: ")
		var input string
		fmt.Scanln(&input)
		if input != "1" && input != "2" {
			fmt.Println("Error: Invalid input.")
		} else {
			selection = input
			break
		}
	}

	var file_path string
	switch selection {
	case "2":
		file_path = "puzzles/9m/sudoku_9mil.csv"
	default:
		file_path = "puzzles/1m/sudoku_1mil.csv"
	}

	fmt.Print("\nLoading puzzle input...\t")

	puzzles := utils.Read(file_path)
	fmt.Println("Puzzles loaded.")

	/*
		When it comes time to actually process all puzzles,
		make a deepcopy of the board and pass a pointer of
		that to each algorithm call. Then, reset that deep-
		copy and pass the pointer to the next algorithm. This
		way we can save time without passing the entire 2D-array
		while also not modifying our copy of the original board.
	*/

	b := utils.Assemble_board(puzzles[0])

	////////////////////////

	//start := time.Now()
	//bruteforce.Backtrack(&b)
	//stop := time.Since(start)
	//fmt.Printf("\nTime to Solve via Backtrack: %s\n\n", stop)

	//utils.Display_board(b)
	//fmt.Println()

	//start = time.Now()
	//solved := utils.Solved(&b)
	//stop = time.Since(start)
	//fmt.Printf("Board Solved Successfully: %t\n", solved)
	//if solved {
	//	fmt.Printf("Time taken to verify: %s\n", stop)
	//}

	fmt.Println("Starting....")
	ai.SimulatedAnnealing(&b)
	fmt.Println("DONE")
}