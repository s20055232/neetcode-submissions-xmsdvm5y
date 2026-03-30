type Solution struct{}


func (s *Solution) Encode(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var b strings.Builder
	for _, str := range strs {
		b.WriteString(strconv.Itoa(len(str)))
		b.WriteString("#")
		b.WriteString(str)
	}
	return b.String()
}

func (s *Solution) Decode(encoded string) []string {
	if encoded == "" {
		return []string{}
	}
	result := []string{}
	l, r := 0, 0
	for r < len(encoded){
		if encoded[r] != '#' {
			r++
		}else {
			length, _ := strconv.Atoi(encoded[l:r])
			l = r + 1
			r = r + 1 + length
			result = append(result, encoded[l:r])
			l = r
			r++
		}
	}
	return result
}
