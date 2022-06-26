# Errors

[![Go Reference](https://pkg.go.dev/badge/github.com/nicolasparada/go-errs.svg)](https://pkg.go.dev/github.com/nicolasparada/go-errs)

Golang package to create constant sentinel errors with `errors.Is` support.

## Install

```bash
go get github.com/nicolasparada/go-errs
```

## Usage

You can create your own constant sentinel errors as if they were a string.

```go
package myapp

import (
    errs "github.com/nicolasparada/go-errs"
)

const (
    ErrInvalidEmail = errs.InvalidArgumentError("myapp: invalid email")
)
```

You can use `errors.Is` to check if your error is known.

```go
package main

import (
    "errors"
    "fmt"

    "myapp"
    errs "github.com/nicolasparada/go-errs"
)

func main() {
    ok := errors.Is(myapp.ErrInvalidEmail, errs.ErrInvalidArgument)
    fmt.Println(ok)
    // Output: true
}
```
