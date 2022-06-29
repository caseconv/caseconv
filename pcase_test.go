// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pcase_test

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/gorelib/pcase"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name      string
		line      string
		input     string
		fromText  []string
		fromCamel []string
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
		}, {
			name:      "two camel words",
			line:      testline(),
			input:     "HelloWorld",
			fromText:  []string{"HelloWorld"},
			fromCamel: []string{"Hello", "World"},
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
		}, {
			name:      "one number by 3 digits",
			line:      testline(),
			input:     "123",
			fromText:  []string{"123"},
			fromCamel: []string{"123"},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.line+"/"+tt.name, func(t *testing.T) {
			t.Parallel()

			get := pcase.New(tt.input)
			ok := len(get) == len(tt.fromText)
			if ok {
				for i := 0; i < len(get); i++ {
					if get[i] != tt.fromText[i] {
						ok = false
						break
					}
				}
			}
			if !ok {
				t.Errorf("\nwant text: %s\n get text: %s\ntest: %s", tt.fromText, get, tt.line)
			}

			get = pcase.New(tt.input, pcase.FromCamel())
			ok = len(get) == len(tt.fromCamel)
			if ok {
				for i := 0; i < len(get); i++ {
					if get[i] != tt.fromCamel[i] {
						ok = false
						break
					}
				}
			}
			if !ok {
				t.Errorf("\nwant camel: %s\n get camel: %s\n      test: %s", tt.fromCamel, get, tt.line)
			}
		})
	}
}

func TestNewText(t *testing.T) {
	tests := []struct {
		name      string
		line      string
		input     string
		fromText  string
		fromCamel string
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
		}, {
			name:      "one number by 3 digits",
			line:      testline(),
			input:     "123",
			fromText:  "123",
			fromCamel: "123",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.line+"/"+tt.name, func(t *testing.T) {
			t.Parallel()

			get := pcase.Text(tt.input)
			if get != tt.fromText {
				t.Errorf("\nwant text: %s\n get text: %s\n     test: %s", tt.fromText, get, tt.line)
			}

			get = pcase.Text(tt.input, pcase.FromCamel())
			if get != tt.fromCamel {
				t.Errorf("\nwant camel: %s\n get camel: %s\n      test: %s", tt.fromCamel, get, tt.line)
			}
		})
	}
}

func TestNewCamel(t *testing.T) {
	tests := []struct {
		name      string
		line      string
		input     string
		fromText  string
		fromCamel string
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
		}, {
			name:      "two camel words",
			line:      testline(),
			input:     "HelloWorld",
			fromText:  "HelloWorld",
			fromCamel: "HelloWorld",
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
		}, {
			name:      "one number of 3 digits",
			line:      testline(),
			input:     "123",
			fromText:  "123",
			fromCamel: "123",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.line+"/"+tt.name, func(t *testing.T) {
			t.Parallel()

			get := pcase.Camel(tt.input)
			if get != tt.fromText {
				t.Errorf("\nwant text: %s\n get text: %s\n     test: %s", tt.fromText, get, tt.line)
			}

			get = pcase.Camel(tt.input, pcase.FromCamel())
			if get != tt.fromCamel {
				t.Errorf("\nwant camel: %s\n get camel: %s\n      test: %s", tt.fromCamel, get, tt.line)
			}
		})
	}
}

func TestCamel(t *testing.T) {
	tests := []struct {
		name  string
		line  string
		input []string
		want  string
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
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.line+"/"+tt.name, func(t *testing.T) {
			t.Parallel()

			get := pcase.Txt(tt.input).Camel()
			if get != tt.want {
				t.Errorf("\nwant: %s\nget:  %s\ntest: %s", tt.want, get, tt.line)
			}
		})
	}
}

func TestNewSnake(t *testing.T) {
	tests := []struct {
		name      string
		line      string
		input     string
		fromText  string
		fromCamel string
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
		}, {
			name:      "two camel words",
			line:      testline(),
			input:     "HelloWorld",
			fromText:  "HelloWorld",
			fromCamel: "Hello_World",
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
		}, {
			name:      "one number of 3 digits",
			line:      testline(),
			input:     "123",
			fromText:  "123",
			fromCamel: "123",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.line+"/"+tt.name, func(t *testing.T) {
			t.Parallel()

			get := pcase.Snake(tt.input)
			if get != tt.fromText {
				t.Errorf("\nwant text: %s\n get text: %s\n     test: %s", tt.fromText, get, tt.line)
			}

			get = pcase.Snake(tt.input, pcase.FromCamel())
			if get != tt.fromCamel {
				t.Errorf("\nwant camel: %s\n get camel: %s\n      test: %s", tt.fromCamel, get, tt.line)
			}
		})
	}
}

func TestSnake(t *testing.T) {
	tests := []struct {
		name  string
		line  string
		input []string
		want  string
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
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.line+"/"+tt.name, func(t *testing.T) {
			t.Parallel()

			get := pcase.Txt(tt.input).Snake()
			if get != tt.want {
				t.Errorf("\nwant: %s\nget:  %s\ntest: %s", tt.want, get, tt.line)
			}
		})
	}
}

func TestNewKebab(t *testing.T) {
	tests := []struct {
		name      string
		line      string
		input     string
		fromText  string
		fromCamel string
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
		}, {
			name:      "two camel words",
			line:      testline(),
			input:     "HelloWorld",
			fromText:  "HelloWorld",
			fromCamel: "Hello-World",
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
		}, {
			name:      "one number of 3 digits",
			line:      testline(),
			input:     "123",
			fromText:  "123",
			fromCamel: "123",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.line, func(t *testing.T) {
			t.Parallel()

			get := pcase.Kebab(tt.input)
			if get != tt.fromText {
				t.Errorf("\nwant text: %s\n get text: %s\n     test: %s", tt.fromText, get, tt.line)
			}

			get = pcase.Kebab(tt.input, pcase.FromCamel())
			if get != tt.fromCamel {
				t.Errorf("\nwant camel: %s\n get camel: %s\n      test: %s", tt.fromCamel, get, tt.line)
			}
		})
	}
}

func TestKebab(t *testing.T) {
	tests := []struct {
		name  string
		line  string
		input []string
		want  string
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
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.line, func(t *testing.T) {
			t.Parallel()

			get := pcase.Txt(tt.input).Kebab()
			if get != tt.want {
				t.Errorf("\nwant: %s\nget:  %s\ntest: %s", tt.want, get, tt.line)
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
