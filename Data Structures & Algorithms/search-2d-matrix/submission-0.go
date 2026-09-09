func searchMatrix(matrix [][]int, target int) bool {
 rowLen , colLen := len(matrix), len(matrix[0])
 row, col := 0,colLen-1 

 for row < rowLen && col >= 0 {
	if target < matrix[row][col] {
		col--
	} else if target > matrix[row][col] {
		row++ 
	} else {
		return true
	}
}
return false
}