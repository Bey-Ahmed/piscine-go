package piscine

import "fmt"

func EightQueens() {
	chessboard := [8][8]int{}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			chessboard[i][j] = 1
			fmt.Print(chessboard[i][j])
		}
		fmt.Println("")
	}
}
