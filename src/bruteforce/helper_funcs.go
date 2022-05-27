package bruteforce

func in_row(board *[9][9]int, row int, val int) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == val {
			return true
		}
	}
	return false
}

func in_col(board *[9][9]int, col int, val int) bool {
	for i := 0; i < 9; i++ {
		if board[i][col] == val {
			return true
		}
	}
	return false
}

func in_square(board *[9][9]int, row int, col int, val int) bool {
	r := row - row%3
	c := col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[r+i][c+j] == val {
				return true
			}
		}
	}
	return false
}

func valid_value(board *[9][9]int, row int, col int, val int) bool {
	if in_row(board, row, val) {
		return false
	}
	if in_col(board, col, val) {
		return false
	}
	return !in_square(board, row, col, val)
}
