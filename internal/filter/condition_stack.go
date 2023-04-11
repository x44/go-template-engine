package filter

import "fmt"

type ConditionStack struct {
	pos int
	buf []bool
}

func NewConditionStack() *ConditionStack {
	return &ConditionStack{
		pos: 0,
		buf: make([]bool, 16),
	}
}

func (s *ConditionStack) Clear() {
	s.pos = 0
}

func (s *ConditionStack) Size() int {
	return s.pos
}

func (s *ConditionStack) Push(b bool) {
	if s.pos == len(s.buf) {
		grow := 32
		s.buf = append(s.buf, make([]bool, grow)...)
	}
	s.buf[s.pos] = b
	s.pos++
}

func (s *ConditionStack) Peek() (bool, error) {
	if s.pos < 1 {
		return false, fmt.Errorf("stack is empty")
	}
	return s.buf[s.pos-1], nil
}

func (s *ConditionStack) Pop() (bool, error) {
	if s.pos < 1 {
		return false, fmt.Errorf("stack is empty")
	}
	s.pos--
	return s.buf[s.pos], nil
}

// Returns true is and only if all stack entries are true
func (s *ConditionStack) Sum() (bool, error) {
	if s.pos < 1 {
		return false, fmt.Errorf("stack is empty")
	}
	for i := 0; i < s.pos; i++ {
		if !s.buf[i] {
			return false, nil
		}
	}
	return true, nil
}
