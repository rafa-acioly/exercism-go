// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"regexp"
	"strings"
)

const spaceString = " "

// Abbreviate return every first letter of each word
func Abbreviate(s string) string {
	var re = regexp.MustCompile(`(?m)[^a-zA-Z\s]`)
	result := re.ReplaceAllString(s, spaceString)

	splitted := strings.Fields(result)
	var response string

	for _, x := range splitted {
		response += strings.ToUpper(string(x[:1]))
	}

	return response
}
