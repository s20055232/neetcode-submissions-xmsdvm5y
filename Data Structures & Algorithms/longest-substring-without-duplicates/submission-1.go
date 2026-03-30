func lengthOfLongestSubstring(s string) int {
    // "" or "1"
    if len(s) == 0 || len(s) == 1{
        return len(s)
    } 
    // "aa"
    head, now := 0, 0
    chars := []rune(s)
    pos := map[rune]int{}
    maxLength := 0

    for now < len(chars) {
        if last, ok := pos[chars[now]]; ok && last >= head {
            head = last + 1
            pos[chars[now]] = now
        }else{
            pos[chars[now]] = now
        }
        maxLength = max(now - head + 1, maxLength)
        now++
    }

    return maxLength
}
