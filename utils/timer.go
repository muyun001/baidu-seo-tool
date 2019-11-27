package utils

import "time"

func TimeFormat(time time.Time) string {
	return time.Format("2006-01-02 15:04:05")
}

// 获取两个时间点相差多少秒
func SecondDiffer(start_time, end_time time.Time) int {
	var minutes int
	t1, err := time.ParseInLocation("2006-01-02 15:04:05", TimeFormat(start_time), time.Local)
	t2, err := time.ParseInLocation("2006-01-02 15:04:05", TimeFormat(end_time), time.Local)
	if err == nil && t1.Before(t2) {
		diff := t2.Unix() - t1.Unix()
		minutes = int(diff / 1)
		return minutes
	}
	return minutes
}
