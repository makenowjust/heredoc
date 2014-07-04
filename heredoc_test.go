// Copyright (c) 2014 TSUYUSATO Kitsune
// This software is released under the MIT License.
// http://opensource.org/licenses/mit-license.php

package heredoc_test

import (
	"testing"
)

import "github.com/MakeNowJust/heredoc"

type testCase struct {
	raw, expect string
}

var tests = []testCase{
	{`
		Foo
		Bar
		`,
		"Foo\nBar\n"},
	{`Foo
		Bar`,
		"Foo\nBar"},
	{`Foo
			
		Bar
		`,
		"Foo\n\t\nBar\n"},
	{`
		Foo
			Bar
				Hoge
					`,
		"Foo\n\tBar\n\t\tHoge\n\t\t\t"},
	{`Foo Bar`, "Foo Bar"},
}

func TestDoc(t *testing.T) {
	for i, test := range tests {
		result := heredoc.Doc(test.raw)
		if result != test.expect {
			t.Errorf("tests[%d] failed: expected=> %#v, result=> %#v", i, test.expect, result)
		}
	}
}
