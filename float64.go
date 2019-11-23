package types

import "strconv"

type Float64 float64

func (f Float64) Int() int{
	return int(f)
}

func (f Float64) Int64() int64{
	return int64(f)
}

func (f Float64) Float64() float64{
	return float64(f)
}

func (f Float64) String() string{
	return strconv.FormatFloat(f.Float64(),'E',-1,64)
}

