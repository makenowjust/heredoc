# heredoc

[![Version](https://img.shields.io/github/v/release/makenowjust/heredoc)](https://github.com/makenowjust/heredoc/releases)
[![Build Status](https://circleci.com/gh/makenowjust/heredoc.svg?style=svg)](https://circleci.com/gh/makenowjust/heredoc)
[![Go Reference](https://pkg.go.dev/badge/github.com/MakeNowJust/heredoc/v2.svg)](https://pkg.go.dev/github.com/MakeNowJust/heredoc/v2)

## About

Package heredoc provides the here-document with keeping indent.

## Import

```go
import "github.com/MakeNowJust/heredoc/v2"
```

## Example

```go
package main

import (
	"fmt"

	"github.com/MakeNowJust/heredoc/v2"
)

func main() {
	fmt.Println(heredoc.Doc(`
		Lorem ipsum dolor sit amet, consectetur adipisicing elit,
		sed do eiusmod tempor incididunt ut labore et dolore magna
		aliqua. Ut enim ad minim veniam, ...
	`))
	// Output:
	// Lorem ipsum dolor sit amet, consectetur adipisicing elit,
	// sed do eiusmod tempor incididunt ut labore et dolore magna
	// aliqua. Ut enim ad minim veniam, ...
	//
}
```

## API Document

[heredoc package - github.com/MakeNowJust/heredoc/v2 - pkg.go.dev](https://pkg.go.dev/github.com/MakeNowJust/heredoc/v2)

## License

This software is released under the MIT License, see LICENSE.
