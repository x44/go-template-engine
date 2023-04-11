package lines

import "strings"

func FromBytes(lines []byte) []string {
	return FromString(string(lines))
}

func FromString(lines string) []string {
	if len(lines) == 0 {
		return []string{}
	}
	lines = strings.TrimSuffix(lines, "\r\n")
	lines = strings.TrimSuffix(lines, "\n")
	return FromStrings(strings.Split(lines, "\n"))
}

func FromStrings(lines []string) []string {
	out := make([]string, len(lines))
	for i, line := range lines {
		out[i] = strings.TrimRight(line, " \r")
	}
	return out
}
