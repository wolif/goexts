package goexts

import (
	"fmt"
	"sync"
)

type Class struct {
	mut    sync.Mutex
	name   string
	super  *Class
	derive map[string]*Class
}

func New(name string) *Class {
	return &Class{
		name:   name,
		derive: map[string]*Class{},
	}
}

func (c *Class) Name() string {
	return c.name
}

func (c *Class) Is(another *Class, strict ...bool) bool {
	if c == another {
		return true
	}
	if len(strict) > 0 && strict[0] {
		return false
	}
	return c.IsDescendantOf(another)
}

func (c *Class) IsBaseClass() bool {
	return c.super == nil
}

func (c *Class) IsDescendantOf(another *Class) bool {
	if c.IsBaseClass() {
		return false
	}
	if c.super == another {
		return true
	}
	return c.super.IsDescendantOf(another)
}

func (c *Class) Extends(name string, cover ...bool) *Class {
	c.mut.Lock()
	defer c.mut.Unlock()
	if _, ok := c.derive[name]; ok {
		if len(cover) > 0 && !cover[0] {
			panic(fmt.Sprintf("class named [%s] exists", name))
		}
	}
	another := &Class{
		name:   name,
		super:  c,
		derive: map[string]*Class{},
	}
	c.derive[name] = another
	return another
}
