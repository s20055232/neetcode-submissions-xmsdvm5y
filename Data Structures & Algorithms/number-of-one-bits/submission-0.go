func hammingWeight(n int) int {
	// the key "bit hack" here is "n & (n - 1)"
	// will remove the rightmost number from the binary
	
	count := 0
	for n != 0 {
		n = n & (n - 1)
		count++
	}
	return count
}
