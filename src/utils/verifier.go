package utils

import (
	"math"
)

func check_blocks(board *[9][9]int) bool {
	/*
		Block Locations:
		  0 | 1 | 2
		  ---------
		  3 | 4 | 5
		  ---------
		  6 | 7 | 8
	*/
	block_sums := [9]int{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			// Use floor division to quickly find the respective blocks that
			// i and j values point to. Added together, this should point to
			// the current block without needing if or switch statements.
			i_block := math.Floor(float64(i)/float64(3)) * 3
			j_block := math.Floor(float64(j) / float64(3))

			block_sums[int(i_block+j_block)] += board[i][j]
		}
	}
	for _, s := range block_sums {
		if s != 45 {
			return false
		}
	}
	return true
}

func check_rows(board *[9][9]int) bool {
	for i := 0; i < 9; i++ {
		sum := 0
		for j := 0; j < 9; j++ {
			sum += board[i][j]
		}
		if sum != 45 {
			return false
		}
	}
	return true
}

func check_cols(board *[9][9]int) bool {
	for i := 0; i < 9; i++ {
		sum := 0
		for j := 0; j < 9; j++ {
			sum += board[j][i]
		}
		if sum != 45 {
			return false
		}
	}
	return true
}

func Solved(board *[9][9]int) bool {
	if !check_blocks(board) {
		return false
	}
	if !check_rows(board) {
		return false
	}
	if !check_cols(board) {
		return false
	}
	return true
}