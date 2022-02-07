package sudoku

/*
	given a 2d array 9x9
	return true if the array is valid sudoku
*/
func IsValidSudoku(board [][]byte) bool {
	var existingMap map[byte]bool

	// check dim 0
	for i := 0; i < len(board); i++ {
		// create a map to store existance
		existingMap = make(map[byte]bool, 9)

		//  check current row
		for j := 0; j < len(board[i]); j++ {
			_, ok := existingMap[board[i][j]]
			if ok && board[i][j] != '.' {
				return false
			}

			existingMap[board[i][j]] = true
		}

	}

	for i := 0; i < len(board); i++ {
		// create a map to store existance
		existingMap = make(map[byte]bool, 9)

		//  check current row
		for j := 0; j < len(board[i]); j++ {
			_, ok := existingMap[board[j][i]]
			if ok && board[j][i] != '.' {
				return false
			}

			existingMap[board[j][i]] = true
		}

	}

	// check blocks
	for xoffset := 0; xoffset < 9; xoffset += 3 {
		for yoffset := 0; yoffset < 9; yoffset += 3 {
			existingMap = make(map[byte]bool, 9)

			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					_, ok := existingMap[board[xoffset+i][yoffset+j]]
					if ok && board[xoffset+i][yoffset+j] != '.' {
						return false
					}

					existingMap[board[xoffset+i][yoffset+j]] = true
				}
			}

		}
	}

	return true

}
