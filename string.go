package types

import "strconv"

// convert from string to other types

type String string

func (s String) Int() int{
	i,_:=strconv.Atoi(s.String())
	return i
}

func (s String) Int64() int64{
	i,_:=strconv.ParseInt(s.String(),10,64)
	return i
}

func (s String) Float64() float64{
	f,_:=strconv.ParseFloat(s.String(),64)
	return f
}


func (s String)String() string{
	return string(s)
}
