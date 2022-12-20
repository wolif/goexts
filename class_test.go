package goexts

import "testing"

type test struct {
	Class *Class
}

func TestClass(t *testing.T) {
	base := New("base")
	baseAnother := New("base another")
	c1 := base.Extends("c1")
	c2 := c1.Extends("c2")

	t1 := &test{Class: c1}
	t2 := &test{Class: c2}
	t.Log(t2.Class.Is(t1.Class))
	t.Log(t2.Class.Is(t1.Class, true))
	t.Log(t2.Class.Is(t2.Class))
	t.Log(t2.Class.IsDescendantOf(t1.Class))
	t.Log(t2.Class.IsDescendantOf(base))
	t.Log(t2.Class.IsDescendantOf(baseAnother))
	t.Log(base.IsBaseClass())
	t.Log(baseAnother.IsBaseClass())
}
