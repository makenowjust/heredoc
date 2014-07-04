#heredoc [![Build Status](https://drone.io/github.com/MakeNowJust/heredoc/status.png)](https://drone.io/github.com/MakeNowJust/heredoc/latest) [![github.com/MakeNowJust/heredoc - Go Walker](http://gowalker.org/api/v1/badge)](https://gowalker.org/github.com/MakeNowJust/heredoc)

##About

Package heredoc provides the here-document with keeping indent.

##Install

```console
$ go get github.com/MakeNowJust/heredoc
```

##Import

```go
import "github.com/MakeNowJust/heredoc"
```

##Example

```go
package main

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
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

##License

This software is released under the MIT License, see LICENSE.
