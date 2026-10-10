func isValidSudoku(board [][]byte) bool {
	row := make([]int, 9)
	col := make([]int, 9)
	box := make([]int, 9)
	for r := 0; r < 9; r++ {
		for c:=0; c<9; c++ {
			if board[r][c] == '.' {
				continue
			}

			val := board[r][c] - '1'
			bit := 1 << val
			sqIdx := (r/3)*3 + c/3
			if row[r] & bit != 0  || col[c] & bit != 0 || box[sqIdx] & bit != 0{
				return false
			}
			row[r] |= bit
			col[c] |= bit
			box[sqIdx] |= bit
		}
	}
	return true
}
