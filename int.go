
package types

// convert from int to other types
import "strconv"

type Int int

func (i Int) Int() int{
	return int(i)
}
//convert to bool ,0 is false ,1 is true
func (i Int) Bool() bool{
	if i.Int() == 0{
		return false
	}
	return true
}


func (i Int)Int64() int64{
	return int64(i)
}

func (i Int) Float64() float64{
	return float64(i)
}

func (i Int) String() string{
	return strconv.Itoa(i.Int())
}
