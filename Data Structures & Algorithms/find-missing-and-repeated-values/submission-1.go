func findMissingAndRepeatedValues(grid [][]int) []int {
    n := len(grid)
	gridSum := 0
	gridSqSum := 0
	for i := 0; i < n; i++ {
		for j := 0; j< n; j++ {
			gridSum += grid[i][j]
			gridSqSum += grid[i][j] * grid[i][j]
		}
	}
	total := n * n * (n*n + 1) / 2
	diff := gridSum - total // a-b grid is current values, total is from 0 to n number s sum
	totalSqSum := n * n *   (n*n + 1) * (2*n*n + 1) / 6
	diffSq := gridSqSum - totalSqSum // (a*a ) - (b*b)
	sum := diffSq / diff // a + b
	a := (sum + diff) / 2
	b := sum - a
	return []int{a,b}
}