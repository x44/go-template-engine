package temple

import (
	"fmt"
	"path/filepath"
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
		InputStrings(in).
		EndOfLine(LF).
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

	exp := []string{
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
		InputStrings(in).
		EndOfLine(LF).
		Filter("var1", true).
		Filter("var2", false).
		Replace("rep1", "replacement1").
		Replace("rep2", "replacement2").
		Replace("empty", "").
		Replace("multi", "multi_line_1\nmulti_line_2").
		Process()

	if err != nil {
		t.Error(err)
	}

	if err := compare(exp, out); err != nil {
		t.Error(err)
	}
}

func TestProcessWalkerIdentity(t *testing.T) {
	root, _, _ := createTmp(t)

	walker := NewWalker(root)
	temple := New()
	err := temple.ProcessWalker(walker)

	exp := []string{
		"line_with_crlf\r\n",
		"line_with_lf\n",
		"line_no_eol",
	}

	if err != nil {
		t.Error(err)
	}

	// check the test files with an identity Temple instance
	out, err := New().
		File(filepath.Join(root, "file1.txt")).
		Process()
	if err != nil {
		t.Error(err)
	}
	if err := compare(exp, out); err != nil {
		t.Error(err)
	}

	out, err = New().
		File(filepath.Join(root, "dir1", "dir2", "file2.txt")).
		Process()
	if err != nil {
		t.Error(err)
	}
	if err := compare(exp, out); err != nil {
		t.Error(err)
	}

	deleteTmp()
}

func TestProcessWalkerEol(t *testing.T) {
	root, _, _ := createTmp(t)

	walker := NewWalker(root)
	temple := New().
		EndOfLine(LF)
	err := temple.ProcessWalker(walker)

	exp := []string{
		"line_with_crlf\n",
		"line_with_lf\n",
		"line_no_eol",
	}

	if err != nil {
		t.Error(err)
	}

	// check the test files with an identity Temple instance
	out, err := New().
		File(filepath.Join(root, "file1.txt")).
		Process()
	if err != nil {
		t.Error(err)
	}
	if err := compare(exp, out); err != nil {
		t.Error(err)
	}

	out, err = New().
		File(filepath.Join(root, "dir1", "dir2", "file2.txt")).
		Process()
	if err != nil {
		t.Error(err)
	}
	if err := compare(exp, out); err != nil {
		t.Error(err)
	}

	deleteTmp()
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
