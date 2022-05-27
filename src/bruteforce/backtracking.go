package bruteforce

func Backtrack(board *[9][9]int) bool {
	var row int
	var col int
	complete := true
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				complete = false
				row = i
				col = j
			}
		}
	}

	if complete {
		return true
	}

	for i := 1; i < 10; i++ {
		if valid_value(board, row, col, i) {
			board[row][col] = i
			if Backtrack(board) {
				return true
			}
		}
	}
	return false
}