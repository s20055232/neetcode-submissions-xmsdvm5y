func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    // [[1,1,1]
	// ,[1,1,0]
	// ,[1,0,1]]

	// 1
	originColor := image[sr][sc]
	visited := map[[2]int]bool{}
	// [1, 1]
	waiting := [][2]int{{sr, sc}}
	// directions := [][2]int{{sr+1, sc+1}, {sr-1, sc-1}, {sr+1, sc-1}, {sr-1, sc+1}}
	for len(waiting) > 0 {
		d := waiting[0]
		waiting = waiting[1:]
		visited[d] = true
		if d[0] < len(image) && d[0] >= 0 && d[1] < len(image[d[0]]) && d[1] >= 0{
			if image[d[0]][d[1]] == originColor {
				image[d[0]][d[1]] = color
				for _, dd := range [][2]int{{d[0]+1, d[1]}, {d[0]-1, d[1]}, {d[0], d[1]+1}, {d[0], d[1]-1}} {
					if _, ok := visited[dd]; ok {
						continue
					}
					waiting = append(waiting, dd)
				}
			}
		}
	}
	return image
}
