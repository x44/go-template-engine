package lines

import "strings"

func FromBytes(in []byte) []string {
	return FromString(string(in))
}

func FromString(in string) []string {
	out := []string{}
	if len(in) > 0 {
		out = strings.Split(in, "\n")
		n := len(out)
		for i := 0; i < n-1; i++ {
			out[i] = out[i] + "\n"
		}
		if strings.HasSuffix(out[n-1], "\r") {
			out[n-1] += "\n"
		}
	}
	return out
}

func FromStrings(lines []string) []string {
	return append([]string{}, lines...)
}

// func FromString(lines string) []string {
// 	if len(lines) == 0 {
// 		return []string{}
// 	}
// 	lines = strings.TrimSuffix(lines, "\r\n")
// 	lines = strings.TrimSuffix(lines, "\n")
// 	return FromStrings(strings.Split(lines, "\n"))
// }

// func FromStrings(lines []string) []string {
// 	out := make([]string, len(lines))
// 	for i, line := range lines {
// 		out[i] = strings.TrimRight(line, " \r")
// 	}
// 	return out
// }
