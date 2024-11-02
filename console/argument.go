package console

import (
	"strconv"
	"strings"
)

type ConArgument struct {
	val string
}

func (a *ConArgument) String() string {
	return a.val
}

func (a *ConArgument) Bool() bool {
	val := strings.ToLower(a.val)
	if val == "true" || val == "1" {
		return true
	} else if val == "false" || val == "0" {
		return false
	}

	return false
}

func (a *ConArgument) Int() int {
	i, _ := strconv.Atoi(a.val)
	return i
}

func (a *ConArgument) Float64() float64 {
	f, _ := strconv.ParseFloat(a.val, 64)
	return f
}
