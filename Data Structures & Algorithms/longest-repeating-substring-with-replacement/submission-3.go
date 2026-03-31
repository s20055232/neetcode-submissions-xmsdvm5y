func characterReplacement(s string, k int) int {
	count := [26]int{}
	result := 0
	left := 0

	for right := 0; right < len(s); right++ {
		count[s[right] - 'A']++

		maxFreq := 0
		for _, c := range count {
			maxFreq = max(maxFreq, c)
		}

		for right - left + 1 - maxFreq > k {
			count[s[left] - 'A']--
			left++
		} 

		result = max(result, right - left + 1)
	}
	
	return result
}
