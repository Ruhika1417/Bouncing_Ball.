package main

import (
	"fmt"
	"time"
	screen "github.com/aditya43/clear-shell-screen-golang"

)

func main() {

	const (
		width      = 50 //x
		height     = 10 //y
		cellEmppty = ' '
		cellBall   = '⚾'
		maxFrame   = 1200
		speed      = time.Second / 20
	)
	var (
		px, py int
		vx, vy = 1, 1
		cell   rune
	) //rune type is an alias to int 32 and can store unicode emoji

	board := make([][]bool, width)

	//for multidimensional slices we also need to initialize each one of the inner slices

	for row := range board {
		board[row] = make([]bool, height)

	}
	screen.clear()

	for i := 0; i < maxFrame; i++ {

		px += vx
		py += vy

		//if ball hits the left or right edges (here horizontal)
		if px <= 0 || py >= width-1 {
			vx *= -1 //reverses velocity
		}
		if py <= 0 || px >= height-1 {
			vy *= -1 //reverses velocity
		}
		//for removing the previous ball
		for y := range board[0] { //y   similar to for y:=0;y<height;y++  //column
			for x := range board {
				board[x][y] = false
			}
		}

		buf := make([]rune, 0, width*height)
		buf = buf[:0]

		for y := range board[0] { //y   similar to for y:=0;y<height;y++  //column
			for x := range board { //x  for x:=0;x<height;x++  //row
				cell = cellEmppty
				if board[x][y] {
					cell = cellBall
					//fmt.Print(" ")

				}
				buf = append(buf, cell, ' ')
				//fmt.Print(" ")
			}
			buf = append(buf, '\n')

		}
		screen.MoveTopLeft()
		fmt.Print(string(buf))
		time.Sleep(speed)
	}
}
