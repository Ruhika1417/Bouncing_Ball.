package main

import "fmt"

func main() {

	const (
		width      = 50 //x
		height     = 10 //y
		cellEmppty = ' '
		cellBall   = '⚾'
	)
	var cell rune //rune type is an alias to int 32 and can store unicode emoji
	board := make([][]bool, width)

	//for multidimensional slices we also need to initialize each one of the inner slices

	for row := range board {
		board[row] = make([]bool, height)

	}

	//drawing the board
	//slice
	board[0][0] = true
	board[1][1] = true
	board[2][2] = true
	board[3][3] = true

	for y := range board[0] { //y   similar to for y:=0;y<height;y++  //column
		for x := range board { //x  for x:=0;x<height;x++  //row
			cell = cellEmppty
			if board[x][y] {
				cell = cellBall
				//fmt.Print(" ")

			}
			fmt.Print(string(cell), " ")
			//fmt.Print(" ")
		}
		fmt.Println()

	}

}
