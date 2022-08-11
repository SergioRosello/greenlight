package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Runtime int32

var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

// Implements the Marshaller interface
func (r Runtime) MarshalJSON() ([]byte, error) {
	str := fmt.Sprintf(`%v mins`, r)
	str = strconv.Quote(str)
	return []byte(str), nil
}

// Implements the Unmarshaller interface
// So that we intercept the unmarshal call
// for the Runtime type, and therefore are
// able to parse `%v mins` to a int32.
func (r *Runtime) UnmarshalJSON(data []byte) error {

	unquoted, err := strconv.Unquote(string(data))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	fields := strings.Split(unquoted, " ")

	// This means the string passed does not
	// meet '%d mins' format
	// The OR operator is important here, because
	// if we use a AND operator, without making sure
	// we have two strings, we will be accessing
	// unallcoated memory, and server will panic.
	if len(fields) != 2 || fields[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}

	// parse fields[0] to int32, so we can
	// parse it again into Runtime.
	i, err := strconv.ParseInt(fields[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	*r = Runtime(i)

	return nil
}
