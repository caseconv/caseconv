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
BenchmarkNew/pcase_test.go:38/two_words_with_extra_spaces-8              1259240           940.9 ns/op
BenchmarkNew/pcase_test.go:45/two_camel_words-8                          1840171           660.1 ns/op
BenchmarkNew/pcase_test.go:64/two_snake_words_with_extra_underscores-8           1411021           859.4 ns/op
BenchmarkNew/pcase_test.go:77/two_kebab_words_with_extra_hyphens-8               1286624           949.0 ns/op
BenchmarkNewFromCamel/pcase_test.go:38/two_words_with_extra_spaces-8               22574         54589 ns/op
BenchmarkNewFromCamel/pcase_test.go:45/two_camel_words-8                           27642         41346 ns/op
BenchmarkNewFromCamel/pcase_test.go:64/two_snake_words_with_extra_underscores-8                26362         43109 ns/op
BenchmarkNewFromCamel/pcase_test.go:77/two_kebab_words_with_extra_hyphens-8                    27056         44742 ns/op
BenchmarkNewText/pcase_test.go:183/two_words_with_extra_spaces-8                              603895          1968 ns/op
BenchmarkNewText/pcase_test.go:196/two_lower_camel_words-8                                    976288          1238 ns/op
BenchmarkNewText/pcase_test.go:209/two_snake_words_with_extra_underscores-8                   641882          1896 ns/op
BenchmarkNewText/pcase_test.go:222/two_kebab_words_with_extra_hyphens-8                       589873          1984 ns/op
BenchmarkNewTextFromCamel/pcase_test.go:183/two_words_with_extra_spaces-8                      20730         54830 ns/op
BenchmarkNewTextFromCamel/pcase_test.go:196/two_lower_camel_words-8                            29925         48053 ns/op
BenchmarkNewTextFromCamel/pcase_test.go:209/two_snake_words_with_extra_underscores-8           24733         45833 ns/op
BenchmarkNewTextFromCamel/pcase_test.go:222/two_kebab_words_with_extra_hyphens-8               24672         48951 ns/op
BenchmarkNewCamel/pcase_test.go:309/two_words_with_extra_spaces-8                             517555          2566 ns/op
BenchmarkNewCamel/pcase_test.go:316/two_camel_words-8                                         817165          1438 ns/op
BenchmarkNewCamel/pcase_test.go:335/two_snake_words_with_extra_underscores-8                  507556          2298 ns/op
BenchmarkNewCamel/pcase_test.go:348/two_kebab_words_with_extra_hyphens-8                      511054          2370 ns/op
BenchmarkNewCamelFromCamel/pcase_test.go:309/two_words_with_extra_spaces-8                     19822         57375 ns/op
BenchmarkNewCamelFromCamel/pcase_test.go:316/two_camel_words-8                                 27086         44381 ns/op
BenchmarkNewCamelFromCamel/pcase_test.go:335/two_snake_words_with_extra_underscores-8          25658         45265 ns/op
BenchmarkNewCamelFromCamel/pcase_test.go:348/two_kebab_words_with_extra_hyphens-8              25555         46415 ns/op
BenchmarkCamel/pcase_test.go:427/two_words_with_blanks-8                                      795375          1542 ns/op
BenchmarkNewSnake/pcase_test.go:486/two_words_with_extra_spaces-8                             601503          2055 ns/op
BenchmarkNewSnake/pcase_test.go:493/two_camel_words-8                                         954084          1242 ns/op
BenchmarkNewSnake/pcase_test.go:512/two_snake_words_with_extra_underscores-8                  640178          1920 ns/op
BenchmarkNewSnake/pcase_test.go:525/two_kebab_words_with_extra_hyphens-8                      619274          1991 ns/op
BenchmarkNewSnakeFromCamel/pcase_test.go:486/two_words_with_extra_spaces-8                     17046         62143 ns/op
BenchmarkNewSnakeFromCamel/pcase_test.go:493/two_camel_words-8                                 24355         43223 ns/op
BenchmarkNewSnakeFromCamel/pcase_test.go:512/two_snake_words_with_extra_underscores-8          26589         45760 ns/op
BenchmarkNewSnakeFromCamel/pcase_test.go:525/two_kebab_words_with_extra_hyphens-8              23811         46552 ns/op
BenchmarkSnake/pcase_test.go:604/two_words_with_blanks-8                                     1000000          1109 ns/op
BenchmarkNewKebab/pcase_test.go:663-8                                                         582560          1984 ns/op
BenchmarkNewKebab/pcase_test.go:670-8                                                         952430          1242 ns/op
BenchmarkNewKebab/pcase_test.go:689-8                                                         616113          1927 ns/op
BenchmarkNewKebab/pcase_test.go:702-8                                                         583974          2008 ns/op
BenchmarkKebab/pcase_test.go:765-8                                                            974294          1130 ns/op
PASS
ok      github.com/pfmt/casefmt 62.138s
```

## TODO

* Add fmt.Formatter
