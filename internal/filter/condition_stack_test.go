package filter

import "testing"

func TestPeek(t *testing.T) {
	stack := NewConditionStack()
	_, err := stack.Peek()
	if err == nil {
		t.Errorf("expected empty stack error")
	}
	stack.Push(true)
	b, err := stack.Peek()
	if err != nil {
		t.Errorf("expected no error")
	}
	if b != true {
		t.Errorf("expected true got false")
	}
	stack.Push(false)
	b, err = stack.Peek()
	if err != nil {
		t.Errorf("expected no error")
	}
	if b != false {
		t.Errorf("expected false got true")
	}
}

func TestPop(t *testing.T) {
	stack := NewConditionStack()
	_, err := stack.Pop()
	if err == nil {
		t.Errorf("expected empty stack error")
	}
	stack.Push(true)
	b, err := stack.Pop()
	if err != nil {
		t.Errorf("expected no error")
	}
	if b != true {
		t.Errorf("expected true got false")
	}
	_, err = stack.Pop()
	if err == nil {
		t.Errorf("expected empty stack error")
	}

	stack.Push(false)
	b, err = stack.Pop()
	if err != nil {
		t.Errorf("expected no error")
	}
	if b != false {
		t.Errorf("expected false got true")
	}

	stack.Push(false)
	stack.Push(true)
	b, err = stack.Pop()
	if err != nil {
		t.Errorf("expected no error")
	}
	if b != true {
		t.Errorf("expected true got false")
	}
	b, err = stack.Pop()
	if err != nil {
		t.Errorf("expected no error")
	}
	if b != false {
		t.Errorf("expected false got true")
	}
}

func TestSum(t *testing.T) {
	stack := NewConditionStack()
	_, err := stack.Sum()
	if err == nil {
		t.Errorf("expected empty stack error")
	}
	stack.Push(true)
	b, err := stack.Sum()
	if err != nil {
		t.Errorf("expected no error")
	}
	if b != true {
		t.Errorf("expected true got false")
	}
	stack.Push(true)
	b, err = stack.Sum()
	if err != nil {
		t.Errorf("expected no error")
	}
	if b != true {
		t.Errorf("expected true got false")
	}
	stack.Push(false)
	b, err = stack.Sum()
	if err != nil {
		t.Errorf("expected no error")
	}
	if b != false {
		t.Errorf("expected false got true")
	}
}
