# Errors

[![Go Reference](https://pkg.go.dev/badge/github.com/nicolasparada/go-errs.svg)](https://pkg.go.dev/github.com/nicolasparada/go-errs)

Golang package to create constant sentinel errors.

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

You can use `errors.As` to check if your error is of type `errs.Error` so you can extract the error kind.

```go
package main

import (
    "errors"
    "fmt"

    "myapp"
    errs "github.com/nicolasparada/go-errs"
)

func main() {
    var e errs.Error
    ok := errors.As(myapp.ErrInvalidEmail, &e)
    fmt.Println(ok)
    // Output: true

    ok = e.Kind() == errs.KindInvalidArgument
    fmt.Println(ok)
    // Output: true
}
```

## HTTP Errors

You can use the `httperrs` subpackage to quickly convert an error
defined using this package into an HTTP status code.

```go
package main

import (
    "errors"
    "fmt"

    "myapp"
    "github.com/nicolasparada/go-errs/httperrs"
)

func main() {
    got := httperrs.Code(myapp.ErrInvalidEmail)
    fmt.Println(got)
    // Output: 422
}
```
