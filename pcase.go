// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package casefmt

import (
	"regexp"
	"strings"
	"sync"
	"unicode"
)

// Txt represents phrase consisting of words, that can be represented as
// a space-separated text, as a text in Camel case, Snake case or Kebab case.
type Txt []string

// New splits a phrase into words.
func New(text string, opts ...Option) []string {
	var cfg Configuration

	for _, opt := range opts {
		opt(&cfg)
	}

	if cfg.camel {
		text = camelFirst.ReplaceAllString(text, "${1} ${2}")
		text = camelAll.ReplaceAllString(text, "${1} ${2}")
	}

	f := func(c rune) bool {
		return contains(stops, c)
	}

	return strings.FieldsFunc(text, f)
}

func contains(a []rune, v rune) bool {
	for _, c := range a {
		if c == v {
			return true
		}
	}
	return false
}

var (
	stops      = []rune{' ', '_', '.', ',', '!', '-'}
	camelFirst = regexp.MustCompile("(.)([A-Z][a-z]+)")
	camelAll   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

// Text returns all words concatenated together by spaces.
func Text(text string, opts ...Option) string {
	return Txt(New(text, opts...)).Text()
}

// Text returns all words concatenated together by spaces.
func (txt Txt) Text() string {
	a := txt.compact()
	s := strings.Join(a, " ")
	pool.Put(&a)
	return s
}

// Camel returns all words concatenated together through camel case.
func Camel(text string, opts ...Option) string {
	return Txt(New(text, opts...)).Camel()
}

// Camel returns all words concatenated together through camel case.
func (txt Txt) Camel() string {
	a := txt.clone()
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

// Snake returns all words concatenated together through underscores.
func Snake(text string, opts ...Option) string {
	return Txt(New(text, opts...)).Snake()
}

// Snake returns all words concatenated together through underscores.
func (txt Txt) Snake() string {
	a := txt.compact()
	s := strings.Join(a, "_")
	pool.Put(&a)
	return s
}

// Kebab returns all words concatenated together through hyphens.
func Kebab(text string, opts ...Option) string {
	return Txt(New(text, opts...)).Kebab()
}

// Kebab returns all words concatenated together through hyphens.
func (txt Txt) Kebab() string {
	a := txt.compact()
	s := strings.Join(a, "-")
	pool.Put(&a)
	return s
}

func (txt Txt) compact() []string {
	a := txt.clone()
	a = a[:0]
	for i := 0; i < len(txt); i++ {
		if len(txt[i]) == 0 {
			continue
		}
		a = append(a, txt[i])
	}
	return a
}

func (txt Txt) clone() []string {
	a := *pool.Get().(*[]string)
	if len(a) < len(txt) {
		a = append(a[:0], make([]string, len(txt))...)
	}
	n := copy(a, txt)
	return a[:n]
}

var pool = sync.Pool{New: func() interface{} { a := []string{}; return &a }}

// Configuration holds values changeable by options.
type Configuration struct{ camel bool }

// Option changes configuration.
type Option func(*Configuration)

// FromCamel respects camel case.
func FromCamel() Option {
	return func(repl *Configuration) { repl.camel = true }
}
