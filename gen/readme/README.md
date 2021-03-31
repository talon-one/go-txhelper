# go-txhelper [![Actions Status](https://github.com/talon-one/go-txhelper/workflows/CI/badge.svg)](https://github.com/talon-one/go-txhelper/actions) [![Coverage Status](https://coveralls.io/repos/github/talon-one/go-txhelper/badge.svg?branch=master)](https://coveralls.io/github/talon-one/go-txhelper?branch=master) [![PkgGoDev](https://img.shields.io/badge/pkg.go.dev-reference-blue)](https://pkg.go.dev/github.com/talon-one/go-txhelper) [![GoDoc](https://godoc.org/github.com/talon-one/go-txhelper?status.svg)](https://godoc.org/github.com/talon-one/go-txhelper) [![go-report](https://goreportcard.com/badge/github.com/talon-one/go-txhelper)](https://goreportcard.com/report/github.com/talon-one/go-txhelper)
go-txhelper is a module that abstracts database transaction behaviour.
It does so by differentiating between Read and Write operations.

It is designed for the use of mappers where you do not know whether you need a transaction or not.

## Example
```go
<$
package main

import (
	"io"
	"os"
)

func main() {
	f, err := os.Open("./example/main.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = io.Copy(os.Stdout, f)
	if err != nil {
		panic(err)
	}
}
-$>
```
