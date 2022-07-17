# caseconv

[![Build Status](https://cloud.drone.io/api/badges/pfmt/caseconv/status.svg)](https://cloud.drone.io/pfmt/caseconv)
[![Go Reference](https://pkg.go.dev/badge/github.com/pfmt/caseconv.svg)](https://pkg.go.dev/github.com/pfmt/caseconv)

Case converter for Go.  
Source files are distributed under the BSD-style license.

## About

The software is considered to be at a alpha level of readiness,
its extremely slow and allocates a lots of memory.

## Benchmark

```sh
$ go test -count=1 -race -bench ./... 
goos: linux
goarch: amd64
pkg: github.com/pfmt/caseconv
cpu: 11th Gen Intel(R) Core(TM) i7-1165G7 @ 2.80GHz
BenchmarkNew/case_test.go:38/two_words_with_extra_spaces-8           1296159           913.9 ns/op
BenchmarkNew/case_test.go:45/two_camel_words-8                       1828016           655.6 ns/op
BenchmarkNew/case_test.go:64/two_snake_words_with_extra_underscores-8            1393654           857.1 ns/op
BenchmarkNew/case_test.go:77/two_kebab_words_with_extra_hyphens-8                1282837           925.3 ns/op
BenchmarkNewFromCamel/case_test.go:38/two_words_with_extra_spaces-8                21603         56251 ns/op
BenchmarkNewFromCamel/case_test.go:45/two_camel_words-8                            27385         40277 ns/op
BenchmarkNewFromCamel/case_test.go:64/two_snake_words_with_extra_underscores-8             28636         44022 ns/op
BenchmarkNewFromCamel/case_test.go:77/two_kebab_words_with_extra_hyphens-8                 26858         43789 ns/op
BenchmarkNewText/case_test.go:183/two_words_with_extra_spaces-8                           608253          1960 ns/op
BenchmarkNewText/case_test.go:196/two_lower_camel_words-8                                 922262          1239 ns/op
BenchmarkNewText/case_test.go:209/two_snake_words_with_extra_underscores-8                541579          1905 ns/op
BenchmarkNewText/case_test.go:222/two_kebab_words_with_extra_hyphens-8                    516870          2024 ns/op
BenchmarkNewTextFromCamel/case_test.go:183/two_words_with_extra_spaces-8                   20266         61152 ns/op
BenchmarkNewTextFromCamel/case_test.go:196/two_lower_camel_words-8                         25473         44641 ns/op
BenchmarkNewTextFromCamel/case_test.go:209/two_snake_words_with_extra_underscores-8        27824         45162 ns/op
BenchmarkNewTextFromCamel/case_test.go:222/two_kebab_words_with_extra_hyphens-8            27364         44596 ns/op
BenchmarkNewCamel/case_test.go:309/two_words_with_extra_spaces-8                          468693          2350 ns/op
BenchmarkNewCamel/case_test.go:316/two_camel_words-8                                      796428          1421 ns/op
BenchmarkNewCamel/case_test.go:335/two_snake_words_with_extra_underscores-8               511460          2281 ns/op
BenchmarkNewCamel/case_test.go:348/two_kebab_words_with_extra_hyphens-8                   497002          2358 ns/op
BenchmarkNewCamelFromCamel/case_test.go:309/two_words_with_extra_spaces-8                  21795         65303 ns/op
BenchmarkNewCamelFromCamel/case_test.go:316/two_camel_words-8                              27828         44177 ns/op
BenchmarkNewCamelFromCamel/case_test.go:335/two_snake_words_with_extra_underscores-8       26343         45620 ns/op
BenchmarkNewCamelFromCamel/case_test.go:348/two_kebab_words_with_extra_hyphens-8           26497         45724 ns/op
BenchmarkCamel/case_test.go:427/two_words_with_blanks-8                                   788952          1497 ns/op
BenchmarkNewSnake/case_test.go:486/two_words_with_extra_spaces-8                          582966          2017 ns/op
BenchmarkNewSnake/case_test.go:493/two_camel_words-8                                      915932          1248 ns/op
BenchmarkNewSnake/case_test.go:512/two_snake_words_with_extra_underscores-8               585519          1897 ns/op
BenchmarkNewSnake/case_test.go:525/two_kebab_words_with_extra_hyphens-8                   571680          1996 ns/op
BenchmarkNewSnakeFromCamel/case_test.go:486/two_words_with_extra_spaces-8                  20046         55738 ns/op
BenchmarkNewSnakeFromCamel/case_test.go:493/two_camel_words-8                              28462         42816 ns/op
BenchmarkNewSnakeFromCamel/case_test.go:512/two_snake_words_with_extra_underscores-8       25154         45698 ns/op
BenchmarkNewSnakeFromCamel/case_test.go:525/two_kebab_words_with_extra_hyphens-8           23715         45124 ns/op
BenchmarkSnake/case_test.go:604/two_words_with_blanks-8                                  1000000          1098 ns/op
BenchmarkNewKebab/case_test.go:663-8                                                      558657          1988 ns/op
BenchmarkNewKebab/case_test.go:670-8                                                      965186          1244 ns/op
BenchmarkNewKebab/case_test.go:689-8                                                      606800          1927 ns/op
BenchmarkNewKebab/case_test.go:702-8                                                      575348          2039 ns/op
BenchmarkKebab/case_test.go:765-8                                                        1000000          1108 ns/op
PASS
ok      github.com/pfmt/caseconv    58.215s
```

## TODO

* Add fmt.Formatter
