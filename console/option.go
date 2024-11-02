package console

import (
	"strconv"
	"strings"
)

type ConOption struct {
	name string
	val  string
	full string
}

func (o *ConOption) String() string {
	return o.val
}

func (o *ConOption) Bool() bool {
	val := strings.ToLower(o.val)
	if val == "true" || val == "1" || val == "yes" || val == "y" {
		return true
	} else if val == "false" || val == "0" || val == "no" || val == "n" {
		return false
	}

	return false
}

func (o *ConOption) Int() int {
	i, _ := strconv.Atoi(o.val)
	return i
}

func (o *ConOption) Float64() float64 {
	f, _ := strconv.ParseFloat(o.val, 64)
	return f
}
