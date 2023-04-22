package testutil

import "strings"

func ReplaceControlChars(in string) string {
	out := in
	out = strings.ReplaceAll(out, "\r", "\\r")
	out = strings.ReplaceAll(out, "\n", "\\n")
	out = strings.ReplaceAll(out, "\t", "\\t")
	return out
}
