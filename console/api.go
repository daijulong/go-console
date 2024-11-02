package console

import (
	"errors"
	"strings"
)

func Args() []*ConArgument {
	con.parse()
	return con.args
}

func Options() []*ConOption {
	con.parse()
	return con.options
}

func Arg(index int, def ...string) *ConArgument {
	con.parse()
	argsCount := len(con.args)
	if argsCount < 1 || index >= argsCount {
		if len(def) > 0 {
			return &ConArgument{val: def[0]}
		}
		return nil
	}
	return con.args[index]
}

func Option(name string, def ...string) *ConOption {
	name = strings.TrimLeft(name, "-")
	con.parse()
	if opt, ok := con.optionMap[name]; ok {
		return opt
	}
	if len(def) > 0 {
		return &ConOption{
			name: name,
			val:  def[0],
			full: "-" + name + "=" + def[0],
		}
	}
	return nil
}

func HasOption(name string) bool {
	name = strings.TrimLeft(name, "-")
	if name == "" {
		return false
	}
	con.parse()
	_, ok := con.optionMap[name]
	return ok
}

func Question(question string, selectOptions ...string) *Ques {
	if question == "" {
		return nil
	}
	return &Ques{
		Question: question,
		Options:  selectOptions,
	}
}

func Confirm(question string, def ...bool) (bool, error) {
	if question == "" {
		return false, errors.New("question missing")
	}
	result := false
	hasDef := false
	if len(def) > 0 {
		result = def[0]
		hasDef = true
	}
	q := Question(question)
	if q == nil {
		return result, errors.New("question missing")
	}
	q.WithValidator(func(answer *Answer) error {
		if answer.String() == "" {
			if hasDef {
				return nil
			}
		}
		a := strings.ToLower(answer.String())
		if a != "yes" && a != "no" && a != "y" && a != "n" {
			return errors.New("invalid answer")
		}
		if a == "yes" || a == "y" {
			result = true
		}
		return nil
	})
	_, err := q.WaitUntil()
	if err != nil {
		return result, err
	}
	return result, nil
}
