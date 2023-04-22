package replacer

import (
	"strings"

	"github.com/x44/go-template-engine/internal/lines"
)

type Replacer struct {
	lines []string
}

const (
	VAR_START = "${"
	VAR_END   = "}"
)

func New(lines []string) *Replacer {
	return &Replacer{
		lines: append([]string{}, lines...),
	}
}

func (r *Replacer) GetOutput() []string {
	return r.lines
}

// Replaces all occurences of ${variable}
// If with is a multi-line string, the lines are inserted with the same indent as the replaced ${variable}
// If with is an empty string and the resulting line is empty after trimming, the resulting line gets removed
func (r *Replacer) Replace(variable string, with string) {
	r.process(variable, with, true)
}

// Replaces first occurence of ${variable}
func (r *Replacer) ReplaceFirst(variable string, with string) {
	r.process(variable, with, false)
}

func (r *Replacer) process(variable string, with string, all bool) {
	out := []string{}
	variable = VAR_START + variable + VAR_END
	withLines := lines.FromString(with)

	firstFound := false
	for _, line := range r.lines {
		if !all && firstFound {
			out = append(out, line)
		} else {
			curLine := line
			added := false
			for {
				resultLines, found, processLastLine := r.processLine(curLine, variable, withLines, all)
				if len(resultLines) == 1 {
					if !added {
						out = append(out, resultLines[0])
						added = true
					} else {
						out[len(out)-1] = resultLines[0]
					}
				} else {
					out = append(out, resultLines...)
					added = true
				}
				if !found {
					break
				}
				firstFound = true

				if (!all && firstFound) || len(resultLines) == 0 {
					break
				}
				if !processLastLine {
					break
				}
				curLine = resultLines[len(resultLines)-1]
			}
		}
	}

	r.lines = out
}

func (r *Replacer) processLine(line, variable string, with []string, all bool) ([]string, bool, bool) {
	pos := strings.Index(line, variable)
	if pos == -1 {
		return []string{line}, false, false
	}
	if len(with) == 0 {
		// remove
		line = strings.Replace(line, variable, "", 1)
		if len(strings.TrimSpace(line)) == 0 {
			// resulting trimmed line is empty -> return nothing
			return []string{}, true, false
		}
		return []string{line}, true, true
	}
	if len(with) == 1 {
		// with is a single line string -> return one line
		line = strings.Replace(line, variable, with[0], 1)
		return []string{line}, true, true
	}
	// with is a multi line string -> return multiple lines
	before := line[:pos]
	after := line[pos+len(variable):]
	out := []string{}
	indent := getIndent(before)
	// append stuff before the replaced var + the first replacement line
	out = append(out, before+with[0])
	// append remaining replacement lines (except the last) with the same indent as the line of the replaced variable
	for i := 1; i < len(with)-1; i++ {
		out = append(out, indent+with[i])
	}
	// append last replacement line + stuff after the replaced var
	out = append(out, indent+with[len(with)-1]+after)
	// signal to caller if the last returned line needs further processing
	return out, true, len(after) > 0
}

func getIndent(s string) string {
	out := []byte{}
	l := len(s)
	i := 0
	for i < l {
		if s[i] == '\t' || s[i] == ' ' {
			out = append(out, s[i])
			i++
		} else {
			break
		}
	}
	return string(out)
}
