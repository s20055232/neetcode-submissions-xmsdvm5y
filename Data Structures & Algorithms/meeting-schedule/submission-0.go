/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func canAttendMeetings(intervals []Interval) bool {
	if len(intervals) < 1 {
		return true
	}
	sort.Slice(intervals, func(i, j int) bool {
        return intervals[i].start < intervals[j].start
    })
	
	for i, interval := range intervals[1:] {
		if intervals[i].end > interval.start{
			return false
		}
	}

	return true
}
