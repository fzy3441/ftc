package tt

import "time"

func Tick()int64  {
	return time.Now().UnixNano()/1e6
}

func TimeFmt()string  {
	return time.Now().Format("2006-01-02 15:04:05")

}