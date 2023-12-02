package stack

import (
	"bytes"
)

type Stack struct {
	values []interface{}
}

func (s *Stack) Push(el interface{}) {
	s.values = append(s.values, el)
}

func (s *Stack) Pop() interface{} {
	v := s.Peek()
	s.values = s.values[:s.Len()-1]
	return v
}

func (s *Stack) Peek() interface{} {
	v := s.values[len(s.values)-1]
	return v
}

func (s *Stack) Len() int {
	return len(s.values)
}

func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Stack) String() string {

	var buffer bytes.Buffer

	for _, l := range s.values {
		if v, ok := l.(string); ok {
			buffer.WriteString(v)
		}
	}

	return buffer.String()
}
