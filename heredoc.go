// Copyright (c) 2014-2024 TSUYUSATO Kitsune
// This software is released under the MIT License.
// http://opensource.org/licenses/mit-license.php

// Package heredoc provides creation of here-documents from raw strings.
//
// Golang supports raw-string syntax.
//
//     doc := `
//     	Foo
//     	Bar
//     `
//
// But raw-string cannot recognize indentation. Thus such content is an indented string, equivalent to
//
//     "\n\tFoo\n\tBar\n"
//
// I dont't want this!
//
// However this problem is solved by package heredoc.
//
//     doc := heredoc.Doc(`
//     	Foo
//     	Bar
//     `)
//
// Is equivalent to
//
//     "Foo\nBar\n"
package heredoc

import (
	"fmt"
	"strings"
)

const maxInt = int(^uint(0) >> 1)

// Doc returns un-indented string as here-document.
func Doc(raw string) string {
	skipFirstLine := false
	if len(raw) > 0 && raw[0] == '\n' {
		raw = raw[1:]
	} else {
		skipFirstLine = true
	}

	rawLines := strings.Split(raw, "\n")
	lines := rawLines
	if skipFirstLine {
		lines = lines[1:]
	}

	minIndentSize := getMinIndent(lines)
	lines = removeIndentation(lines, minIndentSize)

	return strings.Join(rawLines, "\n")
}

// isSpace checks whether the rune represents space or not.
// Only white spcaes (U+0020) and horizontal tabs are treated as space character.
// It is the same as Go.
//
// See https://github.com/MakeNowJust/heredoc/issues/6#issuecomment-524231625.
func isSpace(r rune) bool {
	switch r {
	case ' ', '\t':
		return true
	default:
		return false
	}
}

// getMinIndent calculates the minimum indentation in lines, excluding empty lines.
func getMinIndent(lines []string) int {
	minIndentSize := maxInt

	for i, line := range lines {
		indentSize := 0
		for _, r := range line {
			if isSpace(r) {
				indentSize++
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
func removeIndentation(lines []string, n int) []string {
	for i, line := range lines {
		if len(lines[i]) >= n {
			lines[i] = line[n:]
		}
	}
	return lines
}

// Docf returns unindented and formatted string as here-document.
// Formatting is done as for fmt.Printf().
func Docf(raw string, args ...interface{}) string {
	return fmt.Sprintf(Doc(raw), args...)
}
