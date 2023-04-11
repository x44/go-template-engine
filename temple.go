package temple

import (
	"os"

	"github.com/x44/go-template-engine/internal/eol"
	"github.com/x44/go-template-engine/internal/filter"
	"github.com/x44/go-template-engine/internal/lines"
	"github.com/x44/go-template-engine/internal/replacer"
)

const (
	LF   = "\n"
	CRLF = "\r\n"
)

type Temple struct {
	file         string
	lines        []string
	eol          string
	filters      []*templeFilter
	replacements []*templeReplacement
}

type templeFilter struct {
	variable string
	value    bool
}

type templeReplacement struct {
	variable string
	value    string
	all      bool
}

func New() *Temple {
	return &Temple{}
}

func Hello() {

}

func (t *Temple) SetInputFile(fn string) *Temple {
	t.file = fn
	return t
}

func (t *Temple) SetInputString(s string) *Temple {
	t.lines = lines.FromString(s)
	return t
}

func (t *Temple) SetInputStrings(s []string) *Temple {
	t.lines = lines.FromStrings(s)
	return t
}

func (t *Temple) SetInputBytes(b []byte) *Temple {
	t.lines = lines.FromBytes(b)
	return t
}

func (t *Temple) SetOutputEndOfLine(eol string) *Temple {
	t.eol = eol
	return t
}

func (t *Temple) Filter(variable string, value bool) *Temple {
	t.filters = append(t.filters, &templeFilter{
		variable: variable,
		value:    value,
	})
	return t
}

func (t *Temple) Replace(variable string, value string) *Temple {
	t.replacements = append(t.replacements, &templeReplacement{
		variable: variable,
		value:    value,
		all:      true,
	})
	return t
}

func (t *Temple) ReplaceFirst(variable string, value string) *Temple {
	t.replacements = append(t.replacements, &templeReplacement{
		variable: variable,
		value:    value,
		all:      false,
	})
	return t
}

func (t *Temple) Process() ([]string, error) {
	if len(t.file) > 0 {
		b, err := os.ReadFile(t.file)
		if err != nil {
			return nil, err
		}
		t.SetInputBytes(b)
	}

	var err error = nil
	lines := t.lines[:]

	if len(t.filters) > 0 {
		filter := filter.New(lines)
		for _, f := range t.filters {
			filter.SetVar(f.variable, f.value)
		}
		lines, err = filter.Process()
		if err != nil {
			return nil, err
		}
	}

	if len(t.replacements) > 0 {
		replacer := replacer.New(lines)
		for _, r := range t.replacements {
			if r.all {
				replacer.Replace(r.variable, r.value)
			} else {
				replacer.ReplaceFirst(r.variable, r.value)
			}
		}
		lines = replacer.GetOutput()
	}

	out := []string{}
	if len(lines) > 0 {
		// set EOL of lines except the last line
		// change EOL of last line only if last line has any EOL
		last := lines[len(lines)-1]
		out = append(out, eol.SetAll(lines[:len(lines)-1], t.eol)...)
		out = append(out, eol.Change(last, t.eol))
	}

	return out, nil
}
