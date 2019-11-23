
package types

// convert from int to other types
import "strconv"

type Int64 int64

func (i Int64) Int() int{
	return int(i)
}
//convert to bool ,0 is false ,1 is true
func (i Int64) Bool() bool{
	if i.Int() == 0{
		return false
	}
	return true
}


func (i Int64)Int64() int64{
	return int64(i)
}

func (i Int64) Float64() float64{
	return float64(i)
}

func (i Int64) String() string{
	return strconv.Itoa(i.Int())
}
