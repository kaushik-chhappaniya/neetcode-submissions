func isValidSudoku(board [][]byte) bool {
    rows, cols := make([]int, 9), make([]int, 9)
    squares := make([]int, 9)

    for r:=0; r<9; r++ {
        for c := 0; c<9; c++ {
            if board[r][c] == '.' {
                continue
            }

            val := board[r][c] - '1'
            bit := 1<<val
            squareIndex := (r/3)*3 + c/3 // Formula to find out the index of the particular cell
            if rows[r]&bit != 0 || cols[c]&bit != 0 ||
            squares[squareIndex]&bit != 0 {
                return false
            }
            rows[r] |= bit
            cols[c] |= bit
            squares[squareIndex] |= bit
        }
    }
    return true
}
