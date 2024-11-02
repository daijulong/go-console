package console

import (
	"os"
	"strings"
)

var con *console

func init() {
	con = &console{
		parsed:    false,
		args:      []*ConArgument{},
		options:   []*ConOption{},
		optionMap: map[string]*ConOption{},
	}
}

type console struct {
	parsed bool

	args      []*ConArgument
	options   []*ConOption
	optionMap map[string]*ConOption
}

func (c *console) parse() {
	if c.parsed {
		return
	}

	defer func() {
		c.parsed = true
	}()

	args := os.Args
	if len(args) < 2 {
		return
	}

	for _, arg := range args[1:] {
		if strings.Index(arg, "-") == 0 {
			opt := strings.Split(arg, "=")
			optionKey := strings.TrimLeft(opt[0], "-")
			optionVal := ""
			if len(opt) > 1 {
				optionVal = strings.Join(opt[1:], "=")
			}
			o := &ConOption{
				name: optionKey,
				val:  optionVal,
				full: arg,
			}
			c.options = append(c.options, o)
			c.optionMap[optionKey] = o
		} else {
			c.args = append(c.args, &ConArgument{val: arg})
		}
	}
}
