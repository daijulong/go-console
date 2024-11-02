package console

import (
	"fmt"
)

type Ques struct {
	Key           string
	Question      string
	Options       []string
	Validator     func(answer *Answer) error
	ValidatorFail string
}

func (q *Ques) WithValidator(validator func(answer *Answer) error, fail ...string) *Ques {
	q.Validator = validator
	if len(fail) > 0 {
		q.ValidatorFail = fail[0]
	}
	return q
}

func (q *Ques) showQuestion() {
	Line(q.Question)
	if len(q.Question) > 0 {
		for _, option := range q.Options {
			Line(option)
		}
	}
}

func (q *Ques) Wait() (*Answer, error) {
	q.showQuestion()
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		return nil, err
	}
	answer := newAnswer(input)
	if q.Validator != nil {
		return answer, q.Validator(answer)
	}
	return answer, nil
}

func (q *Ques) WaitUntil() (*Answer, error) {
	var input string
	for {
		q.showQuestion()
		_, err := fmt.Scan(&input)
		if err != nil {
			return nil, err
		}
		answer := newAnswer(input)
		if q.Validator != nil {
			err = q.Validator(answer)
			if err == nil {
				return answer, nil
			}
			if q.ValidatorFail == "" {
				Warning(err.Error())
			} else {
				Warning(q.ValidatorFail)
			}
			continue
		}
		return answer, nil
	}
}
