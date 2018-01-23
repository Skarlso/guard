package main

import (
	"testing"
)

func TestMin(t *testing.T) {
	if 1 != min(1, 2) {
		t.Error("min(1, 2) should return 1")
	}

	if 1 != min(2, 1) {
		t.Error("min(2, 1) should return 1")
	}
}

func TestSetMethods(t *testing.T) {
	n := &node{}

	if n.methods != NONE {
		t.Errorf("n.methods should be NONE, but got: %x", n.methods)
	}

	methods := []HTTPMethod{GET, POST, PUT, DELETE, HEAD, OPTIONS, CONNECT, TRACE, PATCH}
	for _, m := range methods {
		n.setMethods(m)

		if !n.hasMethod(m) {
			t.Errorf("n should have HTTP method %x, but got: %x", m, n.methods)
		}
	}
}

func TestInsertLeaf(t *testing.T) {
	n := &node{}

	n.insertChild("this", "/use/this")

	if n.path != "this" {
		t.Errorf("n.path should be 'this' but got: %s", n.path)
	}

	if n.nType != static {
		t.Errorf("n.nType should be %x but got: %x", static, n.nType)
	}

	if n.methods != NONE {
		t.Errorf("n.methods should be %x but got: %x", NONE, n.methods)
	}

	if n.wildChild != false {
		t.Errorf("n.wildChild should be false but got: %t", n.wildChild)
	}

	if n.indices != "" {
		t.Errorf("n.indices should be empty but got: %s", n.indices)
	}

	if len(n.children) != 0 {
		t.Errorf("n should have no children, but got: %+v", n.children)
	}

	if n.leaf != true {
		t.Errorf("n.leaf should be true but got: %t", n.leaf)
	}

	if n.status == nil {
		t.Error("n.status should not be nil")
	}

	n.insertChild("this", "/use/this", GET)

	if !n.hasMethod(GET) {
		t.Error("n should have HTTP method `GET` been set, but not")
	}
}

func TestInsertChild(t *testing.T) {
	n := &node{}

	n.insertChild("/:name", "/user/:name", GET)

	// check n first
	if n.path != "/" {
		t.Errorf("n.path should be '/' but got: %s", n.path)
	}

	if n.nType != static {
		t.Errorf("n.nType should be %x but got: %x", static, n.nType)
	}

	if n.methods != NONE {
		t.Errorf("n.methods should be %x but got: %x", NONE, n.methods)
	}

	if n.wildChild != true {
		t.Errorf("n.wildChild should be false but got: %t", n.wildChild)
	}

	if n.indices != "" {
		t.Errorf("n.indices should be empty but got: %s", n.indices)
	}

	if len(n.children) != 1 {
		t.Errorf("n should have one children, but got: %+v", n.children)
	}

	if n.leaf != false {
		t.Errorf("n.leaf should be true but got: %t", n.leaf)
	}

	if n.status != nil {
		t.Errorf("n.status should not be nil, but got: %+v", n.status)
	}

	// check it's child, then
	n = n.children[0]

	if n.path != ":name" {
		t.Errorf("n.path should be ':name' but got: %s", n.path)
	}

	if n.nType != param {
		t.Errorf("n.nType should be %x but got: %x", param, n.nType)
	}

	if n.methods != GET {
		t.Errorf("n.methods should be %x but got: %x", GET, n.methods)
	}

	if n.wildChild != false {
		t.Errorf("n.wildChild should be false but got: %t", n.wildChild)
	}

	if n.indices != "" {
		t.Errorf("n.indices should be empty but got: %s", n.indices)
	}

	if len(n.children) != 0 {
		t.Errorf("n should have no children, but got: %+v", n.children)
	}

	if n.leaf != true {
		t.Errorf("n.leaf should be true but got: %t", n.leaf)
	}

	if n.status == nil {
		t.Errorf("n.status should not be nil, but got: %+v", n.status)
	}
}
