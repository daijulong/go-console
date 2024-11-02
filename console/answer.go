package console

import (
	"strconv"
	"strings"
)

type Answer struct {
	val string
}

func newAnswer(answer string) *Answer {
	return &Answer{val: answer}
}

func (a *Answer) String() string {
	return a.val
}

func (a *Answer) Bool() bool {
	val := strings.ToLower(a.val)
	if val == "true" || val == "1" || val == "yes" || val == "y" {
		return true
	} else if val == "false" || val == "0" || val == "no" || val == "n" {
		return false
	}

	return false
}

func (a *Answer) Int() int {
	i, _ := strconv.Atoi(a.val)
	return i
}

func (a *Answer) Float64() float64 {
	f, _ := strconv.ParseFloat(a.val, 64)
	return f
}
