package temple

import (
	"fmt"
	"strings"
	"testing"
)

func TestDry(t *testing.T) {
	in := []string{
		"abc\n",
		"def",
	}

	expected := []string{
		"abc\n",
		"def",
	}

	out, err := New().
		SetInputStrings(in).
		SetOutputEndOfLine(LF).
		Process()

	if err != nil {
		t.Error(err)
	}

	if err := compare(expected, out); err != nil {
		t.Error(err)
	}
}

func TestFull(t *testing.T) {
	in := []string{
		"abc\n",
		"//?var1\n",
		"var1 is set and replaced with ${rep1}\n",
		"//-",
		"//!var2\n",
		"var2 is not set and replaced with ${rep2}\n",
		"//-",
		"//?var1\n",
		"//!var2\n",
		"var1 is set and replaced with ${rep1}\n",
		"var2 is not set and replaced with ${rep2}\n",
		"//-",
		"//-",
		"${empty}",
		"${multi}",
		"def",
	}

	expected := []string{
		"abc\n",
		"var1 is set and replaced with replacement1\n",
		"var2 is not set and replaced with replacement2\n",
		"var1 is set and replaced with replacement1\n",
		"var2 is not set and replaced with replacement2\n",
		"multi_line_1\n",
		"multi_line_2\n",
		"def",
	}

	out, err := New().
		SetInputStrings(in).
		SetOutputEndOfLine(LF).
		Filter("var1", true).
		Filter("var2", false).
		Replace("rep1", "replacement1").
		Replace("rep2", "replacement2").
		Replace("empty", "").
		Replace("multi", "multi_line_1\nmulti_line_2\n").
		Process()

	if err != nil {
		t.Error(err)
	}

	if err := compare(expected, out); err != nil {
		t.Error(err)
	}
}

func compare(a, b []string) error {
	if len(a) != len(b) {
		return fmt.Errorf("\nlengths are different: %d != %d\nexpected:\n%s\ngot:\n%s", len(a), len(b), strings.Join(a, ""), strings.Join(b, ""))
	}
	for i, s := range a {
		if s != b[i] {
			return fmt.Errorf("\nexpected:\n%s\ngot:\n%s", strings.Join(a, ""), strings.Join(b, ""))
		}
	}
	return nil
}
