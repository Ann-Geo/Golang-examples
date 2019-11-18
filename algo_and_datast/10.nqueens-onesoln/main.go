package main

import (
	"fmt"
)

//
const n = 4

type board struct {
	brd [n][n]int
}

func (b *board) initBoard() {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			b.brd[i][j] = 0
		}
	}
}

func (b *board) printBoard() {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(b.brd[i][j])
		}
		fmt.Printf("\n")
	}

	fmt.Printf("\n")
}

func nqueenhelperOneSoln(b *board, col int) bool {
	//if all queens are placed return true
	if col >= n {
		b.printBoard()
		return true
	}

	//Try all rows in the current column
	for row := 0; row < n; row++ {
		//if the queen can be placed in this row,
		//mark this (row, col) as part of the solution
		//
		if b.isSafe(row, col) {
			b.brd[row][col] = 1

			//recursively check placing queen here leads to a solution
			success := nqueenhelperOneSoln(b, col+1)
			if success {
				//if placing queen here leads to a solution return true
				return true
			}
			//if placing queen doesn't lead to a solution
			//then unmark this (row,col)
			b.brd[row][col] = 0

		}
		//try other rows
	}

	//If all rows have been tried and nothing worked
	//return false to trigger backtacking
	return false

}

//we need to check only left side for attacking queens
func (b *board) isSafe(row, col int) bool {
	//check this row on left side
	for i := 0; i < col; i++ {
		if b.brd[row][i] == 1 {
			return false
		}

	}

	//check for upper left diagonal
	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if b.brd[i][j] == 1 {
			return false
		}

	}

	//check for this row on lower left diagonal
	for i, j := row, col; i < n && j >= 0; i, j = i+1, j-1 {
		if b.brd[i][j] == 1 {
			return false
		}

	}
	//it is safe to place the queen here
	return true

}

func solve(b *board) {
	//start with left most column
	nqueenhelperOneSoln(b, 0)
}

func main() {
	var b board
	b.initBoard()
	b.printBoard()
	solve(&b)

}
