type Solution struct{}


func (s *Solution) Encode(strs []string) string {
	var builder strings.Builder
	for _, str := range strs {
		builder.WriteString(str)
		builder.WriteString("你")
	}
	return builder.String()
}

func (s *Solution) Decode(encoded string) []string {
	if len(encoded) == 0 {
		return []string{}
	}
	
	return strings.Split(encoded[:len(encoded)-3], "你")
}
