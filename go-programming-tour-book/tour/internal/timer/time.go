package timer

import "time"

// GetNowTime 获取当前时间
// > cat /etc/localtime
// CST-8 ,即中国标准时间，UTC+8
func GetNowTime() time.Time {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(location)
}

// GetCalculateTime 对时间进行计算
func GetCalculateTime(currentTimer time.Time, d string) (time.Time, error) {
	// 解析出持续时间，支持的有效单位有"ns","us","ms","s","m","h"，例如："300ms","-1.5h","2h35m"
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}
	// 当前时间加上持续时间计算所得时间
	return currentTimer.Add(duration), nil
}
