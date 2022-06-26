// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package idfmt

import (
	"strings"
	"sync"
	"unicode"
)

// Text returns all words concatenated together.
func Text(text string, opts ...Option) string {
	return ID(New(text, opts...)).Text()
}

// Text returns all words concatenated together.
func (id ID) Text() string {
	a := id.compact()
	s := strings.Join(a, " ")
	pool.Put(&a)
	return s
}

// Camel returns all words concatenated together.
func Camel(text string, opts ...Option) string {
	return ID(New(text, opts...)).Camel()
}

// Camel returns all words concatenated together.
func (id ID) Camel() string {
	a := id.clone()
	for i := 0; i < len(a); i++ {
		if len(a[i]) == 0 {
			continue
		}
		runes := []rune(a[i])
		runes[0] = unicode.ToUpper(runes[0])
		a[i] = string(runes)
	}
	s := strings.Join(a, "")
	pool.Put(&a)
	return s
}

// Snake returns all words concatenated together.
func Snake(text string, opts ...Option) string {
	return ID(New(text, opts...)).Snake()
}

// Snake returns all words concatenated together.
func (id ID) Snake() string {
	a := id.compact()
	s := strings.Join(a, "_")
	pool.Put(&a)
	return s
}

// Kebab returns all words concatenated together.
func Kebab(text string, opts ...Option) string {
	return ID(New(text, opts...)).Kebab()
}

// Kebab returns all words concatenated together.
func (id ID) Kebab() string {
	a := id.compact()
	s := strings.Join(a, "-")
	pool.Put(&a)
	return s
}

func (id ID) compact() []string {
	a := id.clone()
	a = a[:0]
	for i := 0; i < len(id); i++ {
		if len(id[i]) == 0 {
			continue
		}
		a = append(a, id[i])
	}
	return a
}

func (id ID) clone() []string {
	a := *pool.Get().(*[]string)
	if cap(a) < len(id) {
		a = append(a, make([]string, len(id)-cap(a))...)
	}
	if len(a) < len(id) {
		a = a[:len(id)]
	}
	n := copy(a, id)
	return a[:n]
}

var pool = sync.Pool{New: func() interface{} { a := []string{}; return &a }}
