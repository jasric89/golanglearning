package main 

import "fmt"

//func chessboard(board [8][8]string ){
//	board[0][0] = "kqrbnp"
//	board[0][7] = "KQRBNP"
//	fmt.Println(board[0][0], board[0][7])
//}

func drawBoard(board [8][8]string){
   j := 0
   for range board {
	fmt.Println(board[j])
	j ++
   }  
}

func main (){
	var board [8][8]string 
	//chessboard(board)
	board[0][0] = "kqrbnp"
	board[0][7] = "KQRBNP"
	fmt.Println(board[0][0], board[0][7])
	drawBoard(board[0])
	
}