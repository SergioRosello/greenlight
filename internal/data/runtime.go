package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

// Implements the Marshaller interface
func (r Runtime) MarshalJSON() ([]byte, error) {
	str := fmt.Sprintf(`%v mins`, r)
	str = strconv.Quote(str)
	return []byte(str), nil
}
