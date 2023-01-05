// Package snickerboa - stack_test.go Test interface implementation of stack
package snickerboa

import "testing"

func TestStackIsEmpty(t *testing.T) {
	s := NewStack()
	if s.IsEmpty() != true {
		t.Errorf("New stack IsEmpty is false; wanted true")
	}
}

func TestStackSize(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	if s.Size() != 3 {
		t.Errorf("New stack size is %d ; wanted 3", s.Size())
	}
}

func TestStackPeek(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(33)
	item, _ := s.Peek()
	if item != 33 {
		t.Errorf("New stack peek is %d ; wanted 33", item)
	}
}

func TestStackPop(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(66)
	item, _ := s.Pop()
	if item != 66 {
		t.Errorf("New stack peek is %d ; wanted 66", item)
	}
}

func TestStackPopString(t *testing.T) {
	s := NewStack()
	s.Push("Kalle")
	s.Push("Olle Olofsson")
	s.Push("foooooooobar 33")
	item, _ := s.Peek()
	if item != "foooooooobar 33" {
		t.Errorf("New stack peek is %s ; wanted long string", item)
	}
}

func TestStackPopMix(t *testing.T) {
	s := NewStack()
	s.Push("Kalle")
	s.Push(100)
	s.Push("foooooooobar 33")
	item, _ := s.Pop()
	if item != "foooooooobar 33" {
		t.Errorf("New stack peek is %s ; wanted long string", item)
	}
	item2, _ := s.Pop()
	if item2 != 100 {
		t.Errorf("New stack peek is %s ; wanted 100", item)
	}
}

func FuzzStackInt(f *testing.F) {
	f.Add(10)

	s := NewStack()

	f.Fuzz(func(t *testing.T, a int) {
		t.Log(a)
		s.Push(a)
		res, _ := s.Pop()
		if res != a {
			t.Errorf("Stack push %d pop value %d", a, res)
		}
	})
}

func FuzzStackString(f *testing.F) {
	f.Add("foobaaar")
	s := NewStack()
	f.Fuzz(func(t *testing.T, a string) {
		t.Log(a)
		s.Push(a)
		res, _ := s.Pop()
		if res != a {
			t.Errorf("Stack push %s pop value %s", a, res)
		}
	})
}
