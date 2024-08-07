package t

import "time"

func GetNowTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
func FormatMysqlDate(dateInt int) string {
	t2, _ := time.ParseInLocation("2006-01-02", "1970-01-01", time.Local)
	return t2.AddDate(0, 0, dateInt).Format("2006-01-02")
}
