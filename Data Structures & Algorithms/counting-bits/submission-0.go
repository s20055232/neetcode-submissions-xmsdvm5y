func countBits(n int) []int {
	result := make([]int, n+1)
	for i := 0 ; i <= n; i++ {
		num := i
		count := 0
	
		for num != 0 {
			num = num & (num - 1)
			count++
		}
		
		result[i] = count
	}
	return result
}
