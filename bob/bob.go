// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"regexp"
	"strings"
)

func isSilence(remark string) bool {
	return len(strings.TrimSpace(remark)) == 0
}

func isShouting(remark string) bool {
	var re = regexp.MustCompile("[^a-zA-Z]")
	return len(re.ReplaceAllString(remark, "")) != 0 && strings.ToUpper(remark) == remark
}

func isShoutingQuestion(remark string) bool {
	return isShouting(remark) && strings.HasSuffix(remark, "?")
}

func isQuestion(remark string) bool {
	return strings.HasSuffix(remark, "?")
}

// Hey should have a comment documenting it.
func Hey(remark string) string {

	message := strings.TrimSpace(remark)

	if isSilence(message) {
		return "Fine. Be that way!"
	}

	if isShoutingQuestion(message) {
		return "Calm down, I know what I'm doing!"
	}

	if isShouting(message) {
		return "Whoa, chill out!"
	}

	if isQuestion(message) {
		return "Sure."
	}

	return "Whatever."
}
