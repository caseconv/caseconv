// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package idfmt

import (
	"regexp"
	"strings"
)

// ID represents identifier consisting of words, that can be represented as
// a space-separated text, as a text from words in Camel case, Snake case or
// Kebab case.
type ID []string

// New splits a identifier phrase into words.
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

// Configuration holds values changeable by options.
type Configuration struct{ camel bool }

// Option changes configuration.
type Option func(*Configuration)

// FromCamel respects camel case.
func FromCamel() Option {
	return func(repl *Configuration) { repl.camel = true }
}
