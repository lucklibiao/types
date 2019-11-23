package types

import (
	"strings"
	"time"
)

var m= map[string]string{
	"%Y":"2006",
	"%M":"01",
	"%D":"02",
	"%h":"15",
	"%m":"04",
	"%s":"05",
}

type Time time.Time

func Now() Time{
	return Time(time.Now())
}
//
func TimeFormat(str string) string{
	for k,v:=range m {
		str= strings.Replace(str,k,time.Now().Format(v),-1)
	}
	return str
}

var timeFormat ="2006-01-02 15:04:06"

// convert Time t to time.Time ,and then can use format function
func (t Time) String() string{
	return time.Time(t).Format(timeFormat)
}

func (t Time) Int64() int64{
	return time.Time(t).UnixNano()
}

func (t Time) Int() int{
	return int(time.Time(t).UnixNano())
}

func (t Time) Float64() float64{
	return float64(time.Time(t).UnixNano())
}

const(
	Second = int64(time.Second)
	Minute=int64(time.Minute)
	Hour = int64(time.Minute)

	Day=Hour*24
	Month=Day*30
	Year=Month*12
)

