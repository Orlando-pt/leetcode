package main

import "fmt"

const BOARD_LENGHT int = 9

type void struct{}

var nothing void

func main() {
	first__ := []byte{'.', '4', '.', '.', '.', '.', '.', '.', '.'}
	second_ := []byte{'.', '.', '4', '.', '.', '.', '.', '.', '.'}
	third__ := []byte{'.', '.', '.', '1', '.', '.', '7', '.', '.'}
	fourth_ := []byte{'.', '.', '.', '.', '.', '.', '.', '.', '.'}
	fifth__ := []byte{'.', '.', '.', '3', '.', '.', '.', '6', '.'}
	sixth__ := []byte{'.', '.', '.', '.', '.', '6', '.', '9', '.'}
	seventh := []byte{'.', '.', '.', '.', '1', '.', '.', '.', '.'}
	eighth_ := []byte{'.', '.', '.', '.', '.', '.', '2', '.', '.'}
	ninth__ := []byte{'.', '.', '.', '8', '.', '.', '.', '.', '.'}
	board := [][]byte{first__, second_, third__, fourth_, fifth__, sixth__, seventh, eighth_, ninth__}

	fmt.Println(isValidSudoku(board))
}

func isValidSudokuFaster(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		m := make(map[byte]bool)

		for j := 0; j < len(board); j++ {
			if board[i][j] != '.' && m[board[i][j]] {
				return false
			} else {
				m[board[i][j]] = true
			}
		}
	}

	for i := 0; i < len(board); i++ {
		m := make(map[byte]bool)

		for j := 0; j < len(board); j++ {
			if board[j][i] != '.' && m[board[j][i]] {
				return false
			} else {
				m[board[j][i]] = true
			}
		}
	}

	for i := 0; i < len(board); i += 3 {
		for j := 0; j < len(board); j += 3 {
			if !validateBox(board, i, j) {
				return false
			}
		}
	}

	return true
}

func validateBox(board [][]byte, posI int, posJ int) bool {
	m := make(map[byte]bool)

	for i := posI; i < posI+3; i++ {
		for j := posJ; j < posJ+3; j++ {
			if board[i][j] != '.' && m[board[i][j]] {
				return false
			} else {
				m[board[i][j]] = true
			}
		}
	}

	return true
}

func isValidSudoku(board [][]byte) bool {

	for i := 0; i < BOARD_LENGHT; i++ {

		// check horizontal
		if !checkHorizontal(board[i]) {
			return false
		}

		// check vertically
		if !checkVertically(board, i) {
			return false
		}

	}

	// check 3x3 (can be optimized eventually)
	return checkThree(board)
}

func checkThree(board [][]byte) bool {
	res := map[int]map[byte]void{}
	subBoxSize := 3

	for l := 0; l < BOARD_LENGHT; l++ {
		for c := 0; c < BOARD_LENGHT; c++ {
			if board[l][c] == '.' {
				continue
			}

			subBox := ((l / subBoxSize) * subBoxSize) + (c / subBoxSize)

			// fmt.Println(subBox)
			// fmt.Println("char: " + string(board[l][c]))
			box := res[subBox]

			if box == nil {
				empty := map[byte]void{}
				res[subBox] = empty
				box = res[subBox]
			}

			if board[l][c] == '.' {
				continue
			}

			// hello, ip := box[board[l][c]]
			//
			// fmt.Println(hello)
			// fmt.Println(ip)

			if _, ok := box[board[l][c]]; ok {
				return false
			}

			box[board[l][c]] = nothing
		}
	}
	return true
}

func checkVertically(board [][]byte, column int) bool {

	res := map[byte]void{}

	for i := 0; i < BOARD_LENGHT; i++ {

		if board[i][column] == '.' {
			continue
		}

		if _, ok := res[board[i][column]]; ok {
			return false
		}

		res[board[i][column]] = nothing
	}

	return true
}

func checkHorizontal(line []byte) bool {

	res := map[byte]void{}

	for i := 0; i < BOARD_LENGHT; i++ {

		if line[i] == '.' {
			continue
		}

		if _, ok := res[line[i]]; ok {
			return false
		}

		res[line[i]] = nothing
	}

	return true
}
