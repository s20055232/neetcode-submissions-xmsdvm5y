func hammingWeight(n int) int {
	// The problem I'm going to solve today is called Hamming Weight.
	// The goal is to count the number of 1 bits in the binary representation of a given integer n.
	// To solve this, we can use a bit hack.
	// The key observation is: when we subtract 1 from n, it flips the rightmost 1 bit to 0, and turns all the bits to its right into 1s.
	// So when we perform n AND n minus one, those changed bits cancel each other out, and the result is n with the rightmost 1 bit removed.
	// We keep repeating this operation until n equals 0, because that means there are no more 1 bits left.
	// The answer is simply the total number of iterations we ran, since each iteration removes exactly one 1 bit.
	// The time complexity is O of k, where k is the number of 1 bits in n, which makes this more efficient than checking every single bit.
	
	count := 0
	for n != 0 {
		n = n & (n - 1)
		count++
	}
	return count
}
