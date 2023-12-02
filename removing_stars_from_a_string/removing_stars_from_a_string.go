package removing_stars_from_a_string

import (
	"bytes"
)

type Stack []interface{}

func (s *Stack) Push(el interface{}) {
	*s = append(*s, el)
}

func (s *Stack) Pop() interface{} {
	v := s.Peek()
	*s = (*s)[:s.Len()-1]
	return v
}

func (s *Stack) Peek() interface{} {
	v := (*s)[len(*s)-1]
	return v
}

func (s *Stack) Len() int {
	return len(*s)
}

func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Stack) String() string {

	var buffer bytes.Buffer

	for _, l := range *s {
		if v, ok := l.(string); ok {
			buffer.WriteString(v)
		}
	}

	return buffer.String()
}

func RemoveStars(s string) string {

	st := Stack{}

	for _, l := range s {
		if l != rune('*') {
			st.Push(string(l))
		} else {
			st.Pop()
		}
	}

	return st.String()

}

func RemoveStars_V2(s string) string {

	r := make([]rune, 0)

	for _, l := range s {
		if l != rune('*') {
			r = append(r, l)
		} else if len(r) > 0 {
			r = r[:len(r)-1]
		}
	}

	return string(r)

}
