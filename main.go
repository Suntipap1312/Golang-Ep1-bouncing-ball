package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {

	const (
		width     = 50
		height    = 10
		cellEmpty = ' '
		cellBall  = '🏐'
		maxFrame  = 1200
	)

	var (
		px, py int
		vx, vy = 1, 1
		cell   rune
	)

	board := make([][]bool, width) // สร้าง slice ตามแนวตั้ง
	for row := range board {
		board[row] = make([]bool, height) //  board index ที่ 1-50 (สร้าง slice ใส่ไปในแต่ละ index อีก 10)
		// fmt.Println(board[row])
	}

	// draw the board

	// for y := 0; y < height; y++ {
	// 	for x := 0; x < width; x++ {
	// 		fmt.Print("X")
	// 		fmt.Print(" ")
	// 	}
	// 	fmt.Println()
	// }

	for i := 0; i < maxFrame; i++ {

		px += vx
		py += vy
		if px <= 0 || px >= width-1 {
			vx *= -1
		}
		if py <= 0 || py >= height-1 {
			vy *= -1
		}
		for y := range board[0] { // จำนวน 10
			// fmt.Println("len", len(board[0]))
			for x := range board { // จำนวน 50
				board[x][y] = false
			}
		}
		board[px][py] = true
		buf := make([]rune, 0, width*height)
		buf = buf[:0]

		for y := range board[0] { // จำนวน 10
			// fmt.Println("len", len(board[0]))
			for x := range board { // จำนวน 50
				cell = ' '
				if board[x][y] { // == true
					cell = cellBall

				}
				// fmt.Print(string(cell), " ")
				buf = append(buf, cell, ' ')
			}
			// fmt.Println()
			buf = append(buf, '\n')
		}

		fmt.Print(string(buf))
		screen.MoveTopLeft()
		time.Sleep(time.Second / 30)

	}
}
