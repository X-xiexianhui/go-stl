// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list

import (
	"fmt"
	"testing"
)

func TestIntList(t *testing.T) {
	l := New[int]()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	l.PushBack(5)
	l.PushBack(6)
	l.PushBack(7)
	l.PushBack(8)
	l.PushBack(9)
}
func TestIterator(t *testing.T) {
	l := New[int]()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	l.PushBack(5)
	l.PushBack(6)
	l.PushBack(7)
	l.PushBack(8)
	l.PushBack(9)
	for e := range l.Iterator() {
		fmt.Println(e)
	}
}
