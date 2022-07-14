// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package caseconv_test

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/pfmt/caseconv"
)

var newTests = []struct {
	name      string
	line      string
	input     string
	fromText  []string
	fromCamel []string
	bench     bool
}{
	{
		name:      "two words",
		line:      testline(),
		input:     "Hello, World!",
		fromText:  []string{"Hello", "World"},
		fromCamel: []string{"Hello", "World"},
	}, {
		name:      "two words with extra exclamations",
		line:      testline(),
		input:     "Hello, World!!",
		fromText:  []string{"Hello", "World"},
		fromCamel: []string{"Hello", "World"},
	}, {
		name:      "two words with extra spaces",
		line:      testline(),
		input:     " Hello,  World! ",
		fromText:  []string{"Hello", "World"},
		fromCamel: []string{"Hello", "World"},
		bench:     true,
	}, {
		name:      "two camel words",
		line:      testline(),
		input:     "HelloWorld",
		fromText:  []string{"HelloWorld"},
		fromCamel: []string{"Hello", "World"},
		bench:     true,
	}, {
		name:      "two lower camel words",
		line:      testline(),
		input:     "helloWorld",
		fromText:  []string{"helloWorld"},
		fromCamel: []string{"hello", "World"},
	}, {
		name:      "two snake words",
		line:      testline(),
		input:     "hello_world",
		fromText:  []string{"hello", "world"},
		fromCamel: []string{"hello", "world"},
	}, {
		name:      "two snake words with extra underscores",
		line:      testline(),
		input:     "_hello__world_",
		fromText:  []string{"hello", "world"},
		fromCamel: []string{"hello", "world"},
		bench:     true,
	}, {
		name:      "two kebab words",
		line:      testline(),
		input:     "hello-world",
		fromText:  []string{"hello", "world"},
		fromCamel: []string{"hello", "world"},
	}, {
		name:      "two kebab words with extra hyphens",
		line:      testline(),
		input:     "-hello--world-",
		fromText:  []string{"hello", "world"},
		fromCamel: []string{"hello", "world"},
		bench:     true,
	}, {
		name:      "one number by 3 digits",
		line:      testline(),
		input:     "123",
		fromText:  []string{"123"},
		fromCamel: []string{"123"},
	},
}

func TestNew(t *testing.T) {
	for _, tt := range newTests {
		tt := tt

		t.Run(tt.line+"/"+tt.name, func(t *testing.T) {
			t.Parallel()

			got := caseconv.New(tt.input)
			ok := len(got) == len(tt.fromText)
			if ok {
				for i := 0; i < len(got); i++ {
					if got[i] != tt.fromText[i] {
						ok = false
						break
					}
				}
			}
			if !ok {
				t.Errorf("\nwant text: %s\n got text: %s\ntest: %s", tt.fromText, got, tt.line)
			}

			got = caseconv.New(tt.input, caseconv.FromCamel())
			ok = len(got) == len(tt.fromCamel)
			if ok {
				for i := 0; i < len(got); i++ {
					if got[i] != tt.fromCamel[i] {
						ok = false
						break
					}
				}
			}
			if !ok {
				t.Errorf("\nwant camel: %s\n got camel: %s\n      test: %s", tt.fromCamel, got, tt.line)
			}
		})
	}
}

func BenchmarkNew(b *testing.B) {
	b.ReportAllocs()

	for _, tt := range newTests {
		if !tt.bench {
			continue
		}

		b.Run(tt.line+"/"+tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = caseconv.New(tt.input)
			}
		})
	}
}

func BenchmarkNewFromCamel(b *testing.B) {
	b.ReportAllocs()

	for _, tt := range newTests {
		if !tt.bench {
			continue
		}

		b.Run(tt.line+"/"+tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = caseconv.New(tt.input, caseconv.FromCamel())
			}
		})
	}
}

var newTextTests = []struct {
	name      string
	line      string
	input     string
	fromText  string
	fromCamel string
	bench     bool
}{
	{
		name:      "two words",
		line:      testline(),
		input:     "Hello, World!",
		fromText:  "Hello World",
		fromCamel: "Hello World",
	}, {
		name:      "two words with extra exclamations",
		line:      testline(),
		input:     "Hello, World!!",
		fromText:  "Hello World",
		fromCamel: "Hello World",
	}, {
		name:      "two words with extra spaces",
		line:      testline(),
		input:     " Hello,  World! ",
		fromText:  "Hello World",
		fromCamel: "Hello World",
		bench:     true,
	}, {
		name:      "two camel words",
		line:      testline(),
		input:     "HelloWorld",
		fromText:  "HelloWorld",
		fromCamel: "Hello World",
	}, {
		name:      "two lower camel words",
		line:      testline(),
		input:     "helloWorld",
		fromText:  "helloWorld",
		fromCamel: "hello World",
		bench:     true,
	}, {
		name:      "two snake words",
		line:      testline(),
		input:     "hello_world",
		fromText:  "hello world",
		fromCamel: "hello world",
	}, {
		name:      "two snake words with extra underscores",
		line:      testline(),
		input:     "_hello__world_",
		fromText:  "hello world",
		fromCamel: "hello world",
		bench:     true,
	}, {
		name:      "two kebab words",
		line:      testline(),
		input:     "hello-world",
		fromText:  "hello world",
		fromCamel: "hello world",
	}, {
		name:      "two kebab words with extra hyphens",
		line:      testline(),
		input:     "-hello--world-",
		fromText:  "hello world",
		fromCamel: "hello world",
		bench:     true,
	}, {
		name:      "one number by 3 digits",
		line:      testline(),
		input:     "123",
		fromText:  "123",
		fromCamel: "123",
	},
}

func TestNewText(t *testing.T) {
	for _, tt := range newTextTests {
		tt := tt
		t.Run(tt.line+"/"+tt.name, func(t *testing.T) {
			t.Parallel()

			got := caseconv.Text(tt.input)
			if got != tt.fromText {
				t.Errorf("\nwant text: %s\n got text: %s\n     test: %s", tt.fromText, got, tt.line)
			}

			got = caseconv.Text(tt.input, caseconv.FromCamel())
			if got != tt.fromCamel {
				t.Errorf("\nwant camel: %s\n got camel: %s\n      test: %s", tt.fromCamel, got, tt.line)
			}
		})
	}
}

func BenchmarkNewText(b *testing.B) {
	b.ReportAllocs()

	for _, tt := range newTextTests {
		if !tt.bench {
			continue
		}

		b.Run(tt.line+"/"+tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = caseconv.Text(tt.input)
			}
		})
	}
}

func BenchmarkNewTextFromCamel(b *testing.B) {
	b.ReportAllocs()

	for _, tt := range newTextTests {
		if !tt.bench {
			continue
		}

		b.Run(tt.line+"/"+tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = caseconv.Text(tt.input, caseconv.FromCamel())
			}
		})
	}
}

var newCamelTests = []struct {
	name      string
	line      string
	input     string
	fromText  string
	fromCamel string
	bench     bool
}{
	{
		name:      "two words",
		line:      testline(),
		input:     "Hello, World!",
		fromText:  "HelloWorld",
		fromCamel: "HelloWorld",
	}, {
		name:      "two words with extra exclamations",
		line:      testline(),
		input:     "Hello, World!!",
		fromText:  "HelloWorld",
		fromCamel: "HelloWorld",
	}, {
		name:      "two words with extra spaces",
		line:      testline(),
		input:     " Hello,  World! ",
		fromText:  "HelloWorld",
		fromCamel: "HelloWorld",
		bench:     true,
	}, {
		name:      "two camel words",
		line:      testline(),
		input:     "HelloWorld",
		fromText:  "HelloWorld",
		fromCamel: "HelloWorld",
		bench:     true,
	}, {
		name:      "two lower camel words",
		line:      testline(),
		input:     "helloWorld",
		fromText:  "HelloWorld",
		fromCamel: "HelloWorld",
	}, {
		name:      "two snake words",
		line:      testline(),
		input:     "hello_world",
		fromText:  "HelloWorld",
		fromCamel: "HelloWorld",
	}, {
		name:      "two snake words with extra underscores",
		line:      testline(),
		input:     "_hello__world_",
		fromText:  "HelloWorld",
		fromCamel: "HelloWorld",
		bench:     true,
	}, {
		name:      "two kebab words",
		line:      testline(),
		input:     "hello-world",
		fromText:  "HelloWorld",
		fromCamel: "HelloWorld",
	}, {
		name:      "two kebab words with extra hyphens",
		line:      testline(),
		input:     "-hello--world-",
		fromText:  "HelloWorld",
		fromCamel: "HelloWorld",
		bench:     true,
	}, {
		name:      "one number of 3 digits",
		line:      testline(),
		input:     "123",
		fromText:  "123",
		fromCamel: "123",
	},
}

func TestNewCamel(t *testing.T) {
	for _, tt := range newCamelTests {
		tt := tt
		t.Run(tt.line+"/"+tt.name, func(t *testing.T) {
			t.Parallel()

			got := caseconv.Camel(tt.input)
			if got != tt.fromText {
				t.Errorf("\nwant text: %s\n got text: %s\n     test: %s", tt.fromText, got, tt.line)
			}

			got = caseconv.Camel(tt.input, caseconv.FromCamel())
			if got != tt.fromCamel {
				t.Errorf("\nwant camel: %s\n got camel: %s\n      test: %s", tt.fromCamel, got, tt.line)
			}
		})
	}
}

func BenchmarkNewCamel(b *testing.B) {
	b.ReportAllocs()

	for _, tt := range newCamelTests {
		if !tt.bench {
			continue
		}

		b.Run(tt.line+"/"+tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = caseconv.Camel(tt.input)
			}
		})
	}
}

func BenchmarkNewCamelFromCamel(b *testing.B) {
	b.ReportAllocs()

	for _, tt := range newCamelTests {
		if !tt.bench {
			continue
		}

		b.Run(tt.line+"/"+tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = caseconv.Camel(tt.input, caseconv.FromCamel())
			}
		})
	}
}

var camelTests = []struct {
	name  string
	line  string
	input []string
	want  string
	bench bool
}{
	{
		name:  "two words",
		line:  testline(),
		input: []string{"Hello", "World"},
		want:  "HelloWorld",
	}, {
		name:  "two words with blanks",
		line:  testline(),
		input: []string{"Hello", "", "World", ""},
		want:  "HelloWorld",
		bench: true,
	},
}

func TestCamel(t *testing.T) {
	for _, tt := range camelTests {
		tt := tt
		t.Run(tt.line+"/"+tt.name, func(t *testing.T) {
			t.Parallel()

			got := caseconv.Txt(tt.input).Camel()
			if got != tt.want {
				t.Errorf("\nwant: %s\nget:  %s\ntest: %s", tt.want, got, tt.line)
			}
		})
	}
}

func BenchmarkCamel(b *testing.B) {
	b.ReportAllocs()

	for _, tt := range camelTests {
		if !tt.bench {
			continue
		}

		b.Run(tt.line+"/"+tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = caseconv.Txt(tt.input).Camel()
			}
		})
	}
}

var newSnakeTests = []struct {
	name      string
	line      string
	input     string
	fromText  string
	fromCamel string
	bench     bool
}{
	{
		name:      "two words",
		line:      testline(),
		input:     "Hello, World!",
		fromText:  "Hello_World",
		fromCamel: "Hello_World",
	}, {
		name:      "two words with extra exclamations",
		line:      testline(),
		input:     "Hello, World!!",
		fromText:  "Hello_World",
		fromCamel: "Hello_World",
	}, {
		name:      "two words with extra spaces",
		line:      testline(),
		input:     " Hello,  World!! ",
		fromText:  "Hello_World",
		fromCamel: "Hello_World",
		bench:     true,
	}, {
		name:      "two camel words",
		line:      testline(),
		input:     "HelloWorld",
		fromText:  "HelloWorld",
		fromCamel: "Hello_World",
		bench:     true,
	}, {
		name:      "two lower camel words",
		line:      testline(),
		input:     "helloWorld",
		fromText:  "helloWorld",
		fromCamel: "hello_World",
	}, {
		name:      "two snake words",
		line:      testline(),
		input:     "hello_world",
		fromText:  "hello_world",
		fromCamel: "hello_world",
	}, {
		name:      "two snake words with extra underscores",
		line:      testline(),
		input:     "_hello__world_",
		fromText:  "hello_world",
		fromCamel: "hello_world",
		bench:     true,
	}, {
		name:      "two kebab words",
		line:      testline(),
		input:     "hello-world",
		fromText:  "hello_world",
		fromCamel: "hello_world",
	}, {
		name:      "two kebab words with extra hyphens",
		line:      testline(),
		input:     "-hello--world-",
		fromText:  "hello_world",
		fromCamel: "hello_world",
		bench:     true,
	}, {
		name:      "one number of 3 digits",
		line:      testline(),
		input:     "123",
		fromText:  "123",
		fromCamel: "123",
	},
}

func TestNewSnake(t *testing.T) {
	for _, tt := range newSnakeTests {
		tt := tt
		t.Run(tt.line+"/"+tt.name, func(t *testing.T) {
			t.Parallel()

			got := caseconv.Snake(tt.input)
			if got != tt.fromText {
				t.Errorf("\nwant text: %s\n got text: %s\n     test: %s", tt.fromText, got, tt.line)
			}

			got = caseconv.Snake(tt.input, caseconv.FromCamel())
			if got != tt.fromCamel {
				t.Errorf("\nwant camel: %s\n got camel: %s\n      test: %s", tt.fromCamel, got, tt.line)
			}
		})
	}
}

func BenchmarkNewSnake(b *testing.B) {
	b.ReportAllocs()

	for _, tt := range newSnakeTests {
		if !tt.bench {
			continue
		}

		b.Run(tt.line+"/"+tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = caseconv.Snake(tt.input)
			}
		})
	}
}

func BenchmarkNewSnakeFromCamel(b *testing.B) {
	b.ReportAllocs()

	for _, tt := range newSnakeTests {
		if !tt.bench {
			continue
		}

		b.Run(tt.line+"/"+tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = caseconv.Snake(tt.input, caseconv.FromCamel())
			}
		})
	}
}

var snakeTests = []struct {
	name  string
	line  string
	input []string
	want  string
	bench bool
}{
	{
		name:  "two words",
		line:  testline(),
		input: []string{"Hello", "World"},
		want:  "Hello_World",
	}, {
		name:  "two words with blanks",
		line:  testline(),
		input: []string{"Hello", "", "World", ""},
		want:  "Hello_World",
		bench: true,
	},
}

func TestSnake(t *testing.T) {
	for _, tt := range snakeTests {
		tt := tt
		t.Run(tt.line+"/"+tt.name, func(t *testing.T) {
			t.Parallel()

			got := caseconv.Txt(tt.input).Snake()
			if got != tt.want {
				t.Errorf("\nwant: %s\nget:  %s\ntest: %s", tt.want, got, tt.line)
			}
		})
	}
}

func BenchmarkSnake(b *testing.B) {
	b.ReportAllocs()

	for _, tt := range snakeTests {
		if !tt.bench {
			continue
		}

		b.Run(tt.line+"/"+tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = caseconv.Txt(tt.input).Snake()
			}
		})
	}
}

var newKebabTests = []struct {
	name      string
	line      string
	input     string
	fromText  string
	fromCamel string
	bench     bool
}{
	{
		name:      "two words",
		line:      testline(),
		input:     "Hello, World!!",
		fromText:  "Hello-World",
		fromCamel: "Hello-World",
	}, {
		name:      "two words with extra exclamations",
		line:      testline(),
		input:     "Hello, World!!",
		fromText:  "Hello-World",
		fromCamel: "Hello-World",
	}, {
		name:      "two words with extra spaces",
		line:      testline(),
		input:     " Hello,  World! ",
		fromText:  "Hello-World",
		fromCamel: "Hello-World",
		bench:     true,
	}, {
		name:      "two camel words",
		line:      testline(),
		input:     "HelloWorld",
		fromText:  "HelloWorld",
		fromCamel: "Hello-World",
		bench:     true,
	}, {
		name:      "two lower camel words",
		line:      testline(),
		input:     "helloWorld",
		fromText:  "helloWorld",
		fromCamel: "hello-World",
	}, {
		name:      "two snake words",
		line:      testline(),
		input:     "hello_world",
		fromText:  "hello-world",
		fromCamel: "hello-world",
	}, {
		name:      "two snake words with extra underscores",
		line:      testline(),
		input:     "_hello__world_",
		fromText:  "hello-world",
		fromCamel: "hello-world",
		bench:     true,
	}, {
		name:      "two kebab words",
		line:      testline(),
		input:     "hello-world",
		fromText:  "hello-world",
		fromCamel: "hello-world",
	}, {
		name:      "two kebab words with extra hyphens",
		line:      testline(),
		input:     "-hello--world-",
		fromText:  "hello-world",
		fromCamel: "hello-world",
		bench:     true,
	}, {
		name:      "one number of 3 digits",
		line:      testline(),
		input:     "123",
		fromText:  "123",
		fromCamel: "123",
	},
}

func TestNewKebab(t *testing.T) {
	for _, tt := range newKebabTests {
		tt := tt
		t.Run(tt.line, func(t *testing.T) {
			t.Parallel()

			got := caseconv.Kebab(tt.input)
			if got != tt.fromText {
				t.Errorf("\nwant text: %s\n got text: %s\n     test: %s", tt.fromText, got, tt.line)
			}

			got = caseconv.Kebab(tt.input, caseconv.FromCamel())
			if got != tt.fromCamel {
				t.Errorf("\nwant camel: %s\n got camel: %s\n      test: %s", tt.fromCamel, got, tt.line)
			}
		})
	}
}

func BenchmarkNewKebab(b *testing.B) {
	b.ReportAllocs()

	for _, tt := range newKebabTests {
		if !tt.bench {
			continue
		}

		b.Run(tt.line, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = caseconv.Kebab(tt.input)
			}
		})
	}
}

var kebabTests = []struct {
	name  string
	line  string
	input []string
	want  string
	bench bool
}{
	{
		name:  "two words",
		line:  testline(),
		input: []string{"Hello", "World"},
		want:  "Hello-World",
	}, {
		name:  "two words with extra spaces",
		line:  testline(),
		input: []string{"Hello", "", "World", ""},
		want:  "Hello-World",
		bench: true,
	},
}

func TestKebab(t *testing.T) {
	for _, tt := range kebabTests {
		tt := tt
		t.Run(tt.line, func(t *testing.T) {
			t.Parallel()

			got := caseconv.Txt(tt.input).Kebab()
			if got != tt.want {
				t.Errorf("\nwant: %s\nget:  %s\ntest: %s", tt.want, got, tt.line)
			}
		})
	}
}

func BenchmarkKebab(b *testing.B) {
	b.ReportAllocs()

	for _, tt := range kebabTests {
		if !tt.bench {
			continue
		}

		b.Run(tt.line, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = caseconv.Txt(tt.input).Kebab()
			}
		})
	}
}

func testline() string {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		return fmt.Sprintf("%s:%d", filepath.Base(file), line)
	}
	return "it was not possible to recover file and line number information about function invocations"
}
