package unionfind

import (
	"container/list"
	"testing"
)

type StringSlice []string
type IntSlice []int

func (s IntSlice) Len() int    { return len(s) }
func (s StringSlice) Len() int { return len(s) }

var ss StringSlice
var is IntSlice
var il *list.List

func initTests() {
	ss = StringSlice{"foo", "bar", "root", "all evil"}
	is = IntSlice{-1, 10, 5, 0, 9}
	il = list.New()
	for i := 0; i < is.Len(); i++ {
		il.PushBack(is[i])
	}
}

func TestFind(t *testing.T) {
	initTests()
	s := New(ss)
	p := s.Find(-2)
	if p != -1 {
		t.Errorf("on invalid index got %d, expected: -1", p)
	}
	p = s.Find(ss.Len() - 1)
	if p != ss.Len()-1 {
		t.Errorf("on last index got %d, expected: %d", p, ss.Len()-1)
	}
	s = New(is)
	p = s.Find(0)
	if p != 0 {
		t.Errorf("on last index got %d, expected: 0", p)
	}
	s = New(il)
	p = s.Find(1)
	if p != 1 {
		t.Errorf("on last index got %d, expected: 1", p)
	}
	s.Union(1, 2)
}

func TestUnion(t *testing.T) {
	initTests()
	s := New(ss)
	err := s.Union(0, -1)
	if err == nil {
		t.Errorf("expected invalid index error")
	}
	err = s.Union(0, 1)
	if err != nil {
		t.Errorf("unexpected error")
	}
	if s.Find(0) != s.Find(1) {
		t.Errorf("got different roots")
	}

	err = s.Union(2, 3)
	if err != nil {
		t.Errorf("unexpected error")
	}
	err = s.Union(2, 2)
	if err != nil {
		t.Errorf("unexpected error")
	}
	if s.Find(2) != s.Find(3) {
		t.Errorf("got different roots")
	}
	err = s.Union(2, 1)
	if err != nil {
		t.Errorf("unexpected error")
	}
	if s.Find(0) != s.Find(3) {
		t.Errorf("got different roots")
	}

	s = New(is)
	s.Union(0, 1)
	s.Union(1, 2)
	s.Union(3, 0)
	if s.Find(0) != s.Find(3) {
		t.Errorf("got different roots")
	}

}
