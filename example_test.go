// Copyright (c) 2014 TSUYUSATO Kitsune
// This software is released under the MIT License.
// http://opensource.org/licenses/mit-license.php

package heredoc_test

import (
	"fmt"
)

import "github.com/MakeNowJust/heredoc"

func ExampleDoc() {
	fmt.Print(heredoc.Doc(`
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
