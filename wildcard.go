package main

import (
	"regexp"
	"strings"
)

func wildcardToRegexp(pattern string) (*regexp.Regexp, error) {
	pattern = strings.TrimSpace(pattern)

	var b strings.Builder
	b.WriteString("(?i)^")

	for _, r := range pattern {
		switch r {
		case '*':
			b.WriteString(".*")
		case '?':
			b.WriteString(".")
		default:
			b.WriteString(regexp.QuoteMeta(string(r)))
		}
	}

	b.WriteString("$")

	return regexp.Compile(b.String())
}
