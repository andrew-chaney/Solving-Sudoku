package utils

import "fmt"

func Assemble_board(data string) [9][9]int {
	board := [9][9]int{}
	ptr := 0
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			board[i][j] = int(data[ptr] - '0')
			ptr++
		}
	}
	return board
}

func Display_board(board *[9][9]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				fmt.Print("   ")
			} else {
				fmt.Printf(" %d ", board[i][j])
			}

			if j%3 == 2 && j != 8 {
				fmt.Print("|")
			}
		}
		fmt.Println()

		if i%3 == 2 && i != 8 {
			fmt.Println("----------------------------")
		}
	}
}