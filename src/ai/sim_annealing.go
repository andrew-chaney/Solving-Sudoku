package ai

import "math/rand"

/*
1. Generate random states (choose a 3x3, swap two values, calculate which is better)
2. Write a cost function (goal is to bring cost down to zero by having no duplicates)
3. Selecting starting temperature (standard dev of the cost of 200 starting states)
4. Calculate the number of iterations per temperature (T)
5. Choose a cooling rate (0.99, but try experimentally to get most accurate but also balance with speed)
*/

func randomly_fill(board *[9][9]int) {
	for i := 0; i < 9; i++ {

	}
}