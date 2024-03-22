package isvalidsudoku

func isValidSudoku(board [][]byte) bool {

	mRow := make(map[byte]int)
	mCol := make([]map[byte]int, 9)
	mBox := make([]map[byte]int, 3)

	for i := 0; i < 9; i++ {
		mCol[i] = map[byte]int{}
	}

	for i := 0; i < 3; i++ {
		mBox[i] = map[byte]int{}
	}

	for i, row := range board {

		if i == 3 {
			mBox[0] = map[byte]int{}
			mBox[1] = map[byte]int{}
			mBox[2] = map[byte]int{}
		}
		if i == 6 {
			mBox[0] = map[byte]int{}
			mBox[1] = map[byte]int{}
			mBox[2] = map[byte]int{}
		}

		for j, col := range row {

			if col == 46 {
				continue
			}

			//Each row must contain the digits 1-9 without repetition.
			if _, ok := mRow[col]; ok {
				return false
			}
			mRow[col]++

			//Each column must contain the digits 1-9 without repetition.
			if _, ok := mCol[j][col]; ok {
				return false
			}
			mCol[j][col]++

			if j < 3 {
				if _, ok := mBox[0][col]; ok {
					return false
				}
				mBox[0][col]++
			} else if j < 6 {
				if _, ok := mBox[1][col]; ok {
					return false
				}
				mBox[1][col]++
			} else {
				if _, ok := mBox[2][col]; ok {
					return false
				}
				mBox[2][col]++
			}
		}
		mRow = make(map[byte]int)
	}

	return true
}
