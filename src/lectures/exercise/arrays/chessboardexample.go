package main 

import (
	"fmt"
)

func main() {
var board [8][8]string

board[0][0] = "kqrbnp"
board[7][0] = "KQRBNP"

for column := range board[1] {
    board[1][column] = "p"
	board[6][column] = "p"
}

fmt.Print(board)
}