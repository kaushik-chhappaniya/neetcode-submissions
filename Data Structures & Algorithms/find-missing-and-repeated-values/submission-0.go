func findMissingAndRepeatedValues(grid [][]int) []int {
    // total length, iterate over 2d arr, create bool arr if present then just append to result
	colLen, rowLen := len(grid), len(grid[0])
	var res []int
	mark := make([]bool, colLen*rowLen+1)
	for i:=0; i<colLen; i++ {
		for j := 0; j< rowLen; j++ {
			if mark[grid[i][j]] {
				// if already true that mean ele doubled
				res = append(res, grid[i][j])
			} else {
				mark[grid[i][j]] = true
			}
		}
	}
	for i:=1; i<=colLen*rowLen; i++ {
		if !mark[i] {
			res = append(res, i)
		}
	}
return res
}