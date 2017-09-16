package main

import (
	"fmt"
	"math/rand"
	"time"
	"sort"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func createBoard(size int) []int {
	return make([]int, size)
}

func main() {
	var boardSize int
	fmt.Printf("Board size: ")
	fmt.Scanln(&boardSize)

	nQueensGenetic(boardSize)
}

type byFitness [][]int

func (s byFitness) Len() int {
	return len(s)
}
func (s byFitness) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byFitness) Less(i, j int) bool {
	return getFitness(&s[i]) < getFitness(&s[j])
}

func nQueensGenetic(boardSize int) {
	batch_size := 5
	board_batch := make([][]int, batch_size)
	for i := 0; i< batch_size; i++ {
		board_batch[i] = createBoard(boardSize)
		setBoard(&board_batch[i])
	}

	for {
		sort.Sort(sort.Reverse(byFitness(board_batch)))
		fmt.Println(board_batch)
		for _, i := range board_batch {
			fmt.Println(getFitness(&i))
			printBoard(&i)
		}
		//breed(&board_batch)
		break
	}
}

func breed(board_batch *[][]int) {
	mutationRate := 0.01
	batchSize := len(*board_batch)
	totalFitness := 0
	fitness := make([]int, batchSize)
	for i, board := range *board_batch {
		fitness[i] = getFitness(&board)
		totalFitness += fitness[i]
	}

	probDist := make([]int, batchSize)
	for i, sum := range probDist {
		probDist[i] = fitness[i]
	}



}

func getFitness(board *[]int) (fitness int) {
	length := len(*board)
	fitness = length
	for i := 1; i < len(*board); i++ {
		if !isValid(board, i) {
			fitness--
		}
	}
	return
}

func isValid(board *[]int, i int) (fitness bool) {
	fitness = true
	for j := i-1; j >= 0; j-- {
		if (*board)[i] == (*board)[j] ||
			(*board)[i] == (*board)[j] + (i-j) ||
			(*board)[i] == (*board)[j] - (i-j) {
			fitness = false
			break
		}
	}
	return
}

func setBoard(board *[]int) {
	length := len(*board)
	for i := 0; i < length; i++ {
		(*board)[i] = r.Intn(length)
	}
}

func printBoard(queens *[]int) {
	length := len(*queens)

	board := make([][]string, length)
	for i := range board {
		board[i] = make([]string, length)
	}

	for i := 0; i<length; i++ {
		for j := 0; j<length; j++ {
			if j == (*queens)[i] {
				board[i][j] = "Q"
			} else {
				board[i][j] = "."
			}
		}

		fmt.Println(board[i])
	}


}

// func nQueens(boardSize)  {
//     const board [boardSize][boardSize]int
//
//     for i := 0; i < boardSize; i++ {
//         placeQueen(board, i)
//     }
// }
//
// func placeQueen(board [][]int, i int) {
//     for j := 0; j < len(board); j++ {
//         if checkValidQueen(board, i, j) {
//             board[i][j] = 1
//             return
//         } else {
//
//         }
//     }
//
// }
