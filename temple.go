package temple

import (
	"fmt"
	"os"
	"strings"

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
	dryRun       bool
	inFile       string
	outFile      string
	lines        []string
	eol          string
	isEolSet     bool
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
	return &Temple{
		dryRun: false,
	}
}

func (t *Temple) DryRun(dryRun bool) *Temple {
	t.dryRun = dryRun
	return t
}

// Sets both the input and output file.
func (t *Temple) File(fn string) *Temple {
	t.InputFile(fn)
	t.OutputFile(fn)
	return t
}

func (t *Temple) InputFile(fn string) *Temple {
	t.inFile = fn
	return t
}

func (t *Temple) InputString(s string) *Temple {
	t.lines = lines.FromString(s)
	return t
}

func (t *Temple) InputStrings(s []string) *Temple {
	t.lines = lines.FromStrings(s)
	return t
}

func (t *Temple) OutputFile(fn string) *Temple {
	t.outFile = fn
	return t
}

func (t *Temple) InputBytes(b []byte) *Temple {
	t.lines = lines.FromBytes(b)
	return t
}

func (t *Temple) EndOfLine(eol string) *Temple {
	t.eol = eol
	t.isEolSet = true
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
	if t.dryRun {
		src := "mem"
		dst := "mem"
		if len(t.inFile) > 0 {
			src = "file " + t.inFile
		}
		if len(t.outFile) > 0 {
			dst = "file " + t.outFile
		}
		fmt.Printf("Temple dry run: %s -> %s\n", src, dst)
		return nil, nil
	}

	if len(t.inFile) > 0 {
		b, err := os.ReadFile(t.inFile)
		if err != nil {
			return nil, err
		}
		t.InputBytes(b)
	}

	var err error = nil
	// t.lines already is a copy
	lines := t.lines

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
	if t.isEolSet {
		if len(lines) > 0 {
			// set EOL of lines except the last line
			// change EOL of last line only if last line has any EOL
			last := lines[len(lines)-1]
			out = append(out, eol.SetAll(lines[:len(lines)-1], t.eol)...)
			out = append(out, eol.Change(last, t.eol))
		}
	} else {
		out = lines
	}

	if t.outFile != "" {
		err = os.WriteFile(t.outFile, []byte(strings.Join(out, "")), os.ModePerm)
		if err != nil {
			return nil, err
		}
	}

	return out, nil
}

func (t *Temple) ProcessWalker(walker *Walker) error {
	if t.inFile != "" || t.outFile != "" || t.lines != nil {
		return fmt.Errorf("input is already set")
	}

	var err error

	walker.Walk(func(path string, isFile bool) {
		if err != nil {
			return
		}
		if !isFile {
			return
		}
		t.File(path)
		_, err = t.Process()
	})
	return err
}
