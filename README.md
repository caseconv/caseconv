# casefmt

[![Build Status](https://cloud.drone.io/api/badges/pfmt/casefmt/status.svg)](https://cloud.drone.io/pfmt/casefmt)
[![Go Reference](https://pkg.go.dev/badge/github.com/pfmt/casefmt.svg)](https://pkg.go.dev/github.com/pfmt/casefmt)

Source files are distributed under the BSD-style license.

## About

The software is considered to be at a alpha level of readiness -
its extremely slow and allocates a lots of memory)

## Benchmark

```sh
$ go test -count=1 -race -bench ./... 
goos: linux
goarch: amd64
pkg: github.com/pfmt/casefmt
cpu: 11th Gen Intel(R) Core(TM) i7-1165G7 @ 2.80GHz
BenchmarkNew/casefmt_test.go:38/two_words_with_extra_spaces-8              1274884           909.9 ns/op
BenchmarkNew/casefmt_test.go:45/two_camel_words-8                          1813984           658.9 ns/op
BenchmarkNew/casefmt_test.go:64/two_snake_words_with_extra_underscores-8           1393968           859.7 ns/op
BenchmarkNew/casefmt_test.go:77/two_kebab_words_with_extra_hyphens-8               1271848           958.9 ns/op
BenchmarkNewFromCamel/casefmt_test.go:38/two_words_with_extra_spaces-8               22382         55642 ns/op
BenchmarkNewFromCamel/casefmt_test.go:45/two_camel_words-8                           26655         41819 ns/op
BenchmarkNewFromCamel/casefmt_test.go:64/two_snake_words_with_extra_underscores-8                26662         43748 ns/op
BenchmarkNewFromCamel/casefmt_test.go:77/two_kebab_words_with_extra_hyphens-8                    26368         45600 ns/op
BenchmarkNewText/casefmt_test.go:183/two_words_with_extra_spaces-8                              585213          1999 ns/op
BenchmarkNewText/casefmt_test.go:196/two_lower_camel_words-8                                    940408          1241 ns/op
BenchmarkNewText/casefmt_test.go:209/two_snake_words_with_extra_underscores-8                   566400          1942 ns/op
BenchmarkNewText/casefmt_test.go:222/two_kebab_words_with_extra_hyphens-8                       548078          2032 ns/op
BenchmarkNewTextFromCamel/casefmt_test.go:183/two_words_with_extra_spaces-8                      20823         55103 ns/op
BenchmarkNewTextFromCamel/casefmt_test.go:196/two_lower_camel_words-8                            27100         42516 ns/op
BenchmarkNewTextFromCamel/casefmt_test.go:209/two_snake_words_with_extra_underscores-8           26200         45765 ns/op
BenchmarkNewTextFromCamel/casefmt_test.go:222/two_kebab_words_with_extra_hyphens-8               25105         46922 ns/op
BenchmarkNewCamel/casefmt_test.go:309/two_words_with_extra_spaces-8                             438352          2354 ns/op
BenchmarkNewCamel/casefmt_test.go:316/two_camel_words-8                                         717250          1422 ns/op
BenchmarkNewCamel/casefmt_test.go:335/two_snake_words_with_extra_underscores-8                  514124          2318 ns/op
BenchmarkNewCamel/casefmt_test.go:348/two_kebab_words_with_extra_hyphens-8                      477188          2393 ns/op
BenchmarkNewCamelFromCamel/casefmt_test.go:309/two_words_with_extra_spaces-8                     21337         58760 ns/op
BenchmarkNewCamelFromCamel/casefmt_test.go:316/two_camel_words-8                                 24538         41815 ns/op
BenchmarkNewCamelFromCamel/casefmt_test.go:335/two_snake_words_with_extra_underscores-8          24662         48660 ns/op
BenchmarkNewCamelFromCamel/casefmt_test.go:348/two_kebab_words_with_extra_hyphens-8              25239         46033 ns/op
BenchmarkCamel/casefmt_test.go:427/two_words_with_blanks-8                                      692812          1539 ns/op
BenchmarkNewSnake/casefmt_test.go:486/two_words_with_extra_spaces-8                             495900          2060 ns/op
BenchmarkNewSnake/casefmt_test.go:493/two_camel_words-8                                         880149          1255 ns/op
BenchmarkNewSnake/casefmt_test.go:512/two_snake_words_with_extra_underscores-8                  587776          1947 ns/op
BenchmarkNewSnake/casefmt_test.go:525/two_kebab_words_with_extra_hyphens-8                      597076          2046 ns/op
BenchmarkNewSnakeFromCamel/casefmt_test.go:486/two_words_with_extra_spaces-8                     19158         59589 ns/op
BenchmarkNewSnakeFromCamel/casefmt_test.go:493/two_camel_words-8                                 27078         41061 ns/op
BenchmarkNewSnakeFromCamel/casefmt_test.go:512/two_snake_words_with_extra_underscores-8          24880         45428 ns/op
BenchmarkNewSnakeFromCamel/casefmt_test.go:525/two_kebab_words_with_extra_hyphens-8              24924         45583 ns/op
BenchmarkSnake/casefmt_test.go:604/two_words_with_blanks-8                                     1000000          1112 ns/op
BenchmarkNewKebab/casefmt_test.go:663-8                                                         573470          1999 ns/op
BenchmarkNewKebab/casefmt_test.go:670-8                                                         972392          1262 ns/op
BenchmarkNewKebab/casefmt_test.go:689-8                                                         591441          1949 ns/op
BenchmarkNewKebab/casefmt_test.go:702-8                                                         566683          2234 ns/op
BenchmarkKebab/casefmt_test.go:765-8                                                            756885          1604 ns/op
PASS
ok      github.com/pfmt/casefmt    56.880s
```

## TODO

* Add fmt.Formatter
