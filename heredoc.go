// Copyright (c) 2014 TSUYUSATO Kitsune
// This software is released under the MIT License.
// http://opensource.org/licenses/mit-license.php

// Package heredoc provides the here-document with keeping indent.
//
// Golang supports raw-string syntax.
//     doc := `
//     	Foo
//     	Bar
//     `
// But raw-string cannot recognize indent. Thus such content is indented string, equivalent to
//     "\n\tFoo\n\tBar\n"
// I dont't want this!
//
// However this problem is solved by package heredoc.
//     doc := heredoc.Doc(`
//     	Foo
//     	Bar
//     `)
// It is equivalent to
//     "Foo\nBar\n"
package heredoc

import (
	"fmt"
	"strings"
	"unicode"
)

const maxInt = int(^uint(0) >> 1)

// heredoc.Doc retutns unindented string as here-document.
//
// Process of making here-document:
//     1. Find most little indent size. (Skip empty lines)
//     2. Remove this indents of lines.
func Doc(raw string) string {
	skipFirstLine := false
	if raw[0] == '\n' {
		raw = raw[1:]
	} else {
		skipFirstLine = true
	}

	lines := strings.Split(raw, "\n")

	minIndentSize := getMinIndent(lines, skipFirstLine)
	lines = removeIndentation(lines, minIndentSize, skipFirstLine)

	return strings.Join(lines, "\n")
}

// getMinIndent calculates the minimum indentation in lines, excluding empty lines.
func getMinIndent(lines []string, skipFirstLine bool) int {
	minIndentSize := maxInt

	for i, line := range lines {
		if i == 0 && skipFirstLine {
			continue
		}

		indentSize := 0
		for _, r := range []rune(line) {
			if unicode.IsSpace(r) {
				indentSize += 1
			} else {
				break
			}
		}

		if len(line) == indentSize {
			if i == len(lines)-1 && indentSize < minIndentSize {
				lines[i] = ""
			}
		} else if indentSize < minIndentSize {
			minIndentSize = indentSize
		}
	}
	return minIndentSize
}

// removeIndentation removes n characters from the front of each line in lines.
// Skips first line if skipFirstLine is true, skips empty lines.
func removeIndentation(lines []string, n int, skipFirstLine bool) []string {
	for i, line := range lines {
		if i == 0 && skipFirstLine {
			continue
		}

		if len(lines[i]) >= n {
			lines[i] = line[n:]
		}
	}
	return lines
}

// heredoc.Docf returns unindented and formatted string as here-document.
// This format is same with package fmt's format.
func Docf(raw string, args ...interface{}) string {
	return fmt.Sprintf(Doc(raw), args...)
}
