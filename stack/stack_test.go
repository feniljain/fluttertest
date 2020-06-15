package stack

import (
	"testing"
)

func TestStackInit(t *testing.T) {
	s := Init()
	if len(s.stack) != 0 {
		t.Error("Expected stack of size 0, got:", len(s.stack))
	}
	if s.Top != -1 {
		t.Error("Expected stack's top to be of size 0, but got:", len(s.stack))
	}
}

func TestStackPush(t *testing.T) {
	s := Init()
	s.Push("(")
	if s.stack[len(s.stack)-1] != "(" {
		t.Error("Expected stack to be [(], but got:", s.stack)
	}
	if s.Top != 0 {
		t.Error("Expected top to be 1, but got:", s.Top)
	}
}

func TestPopStack(t *testing.T) {
	s := Init()
	s.Push("(")
	s.Push(")")
	ch, err := s.Pop()
	if err != nil && ch != "(" {
		t.Error("Expected stack to be [(], but got:", err)
	}
	if s.stack[len(s.stack)-1] != "(" {
		t.Error("Expected stack to be [(], but got:", s.stack)
	}

	s.Pop()
	ch, err = s.Pop()
	_, ok := err.(*UnderFlowError)
	if !ok && ch != "" {
		t.Error("Expected UnderFlowError, but got:", err)
	}
}

func TestTopElement(t *testing.T) {
	s := Init()
	s.Push("(")
	str, _ := s.TopElement()
	if str != "(" {
		t.Error("Expected ( to be the top element but got:", str)
	}
}
