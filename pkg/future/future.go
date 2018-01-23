package future

import (
	"errors"
)

type SuccessFunc func(string)

type FailFunc func(error)

type ExecuteStringFunc func(string, error)

type MaybeString struct {
	successFunc SuccessFunc
	failFunc    FailFunc
}

func (s *MaybeString) Success(f SuccessFunc) *MaybeString {
	s.successFunc = f
	return s
}

func (s *MaybeString) Fail(f FailFunc) *MaybeString {
	s.failFunc = f
	return s
}

func (s *MaybeString) Execute(f ExecuteStringFunc) *MaybeString {
	go func(s *MaybeString) {
		err := errors.New("Error ocurred")
		str := ""
		if err != nil {
			s.failFunc(err)
		} else {
			s.successFunc(str)
		}
	}(s)
	return nil
}

func main() {

}
