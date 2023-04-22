package replacer

import (
	"testing"

	"github.com/x44/go-template-engine/internal/lines"
	"github.com/x44/go-template-engine/internal/testutil"
)

func TestEmpty(t *testing.T) {
	in := ""
	replacer := New(lines.FromString(in))
	replacer.Replace("var", "abc")
	out := replacer.GetOutput()
	if len(out) != 0 {
		t.Errorf("expected 0 got %d", len(out))
	}
}

func TestSimpleNoVar(t *testing.T) {
	in := "abc"
	replacer := New(lines.FromString(in))
	replacer.Replace("var", "abc")
	out := replacer.GetOutput()
	if len(out) != 1 {
		t.Errorf("expected 1 got %d", len(out))
	}
	if out[0] != "abc" {
		t.Errorf("expected abc got %s", out[0])
	}
}

func TestSingleLineWithSingleLineVarAll(t *testing.T) {
	in := "abc${var}def${var}"
	replacer := New(lines.FromString(in))
	replacer.Replace("var", "xyz")
	out := replacer.GetOutput()
	if len(out) != 1 {
		t.Errorf("expected 1 got %d", len(out))
	}
	if out[0] != "abcxyzdefxyz" {
		t.Errorf("expected abcxyzdefxyz got %s", out[0])
	}
}

func TestSingleLineWithSingleLineVarFirst(t *testing.T) {
	in := "abc${var}def${var}"
	replacer := New(lines.FromString(in))
	replacer.ReplaceFirst("var", "xyz")
	out := replacer.GetOutput()
	if len(out) != 1 {
		t.Errorf("expected 1 got %d", len(out))
	}
	if out[0] != "abcxyzdef${var}" {
		t.Errorf("expected abcxyzdef${var} got %s", out[0])
	}
}

func TestMultiLineWithSingleLineVarAll(t *testing.T) {
	in := []string{
		"abc${var}def${var}",
		"abc${var}",
	}
	replacer := New(in)
	replacer.Replace("var", "xyz")
	out := replacer.GetOutput()
	if len(out) != 2 {
		t.Errorf("expected 2 got %d", len(out))
	}
	if out[0] != "abcxyzdefxyz" {
		t.Errorf("expected abcxyzdefxyz got %s", out[0])
	}
	if out[1] != "abcxyz" {
		t.Errorf("expected abcxyz got %s", out[1])
	}
}

func TestMultiLineWithSingleLineVarFirst(t *testing.T) {
	in := []string{
		"abc${var}def${var}",
		"abc${var}",
	}
	replacer := New(in)
	replacer.ReplaceFirst("var", "xyz")
	out := replacer.GetOutput()
	if len(out) != 2 {
		t.Errorf("expected 2 got %d", len(out))
	}
	if out[0] != "abcxyzdef${var}" {
		t.Errorf("expected abcxyzdef${var} got %s", out[0])
	}
	if out[1] != "abc${var}" {
		t.Errorf("expected abc${var} got %s", out[1])
	}
}

func TestMultiLineWithEmptyVar(t *testing.T) {
	in := []string{
		"abc${var}",
		"${var}",
		"${var}abc",
	}
	replacer := New(in)
	replacer.Replace("var", "")
	out := replacer.GetOutput()
	if len(out) != 2 {
		t.Errorf("expected 2 got %d", len(out))
	}
	if out[0] != "abc" {
		t.Errorf("expected abc got %s", out[0])
	}
	if out[1] != "abc" {
		t.Errorf("expected abc got %s", out[1])
	}
}

func TestMultiLineWithMultiLineVar(t *testing.T) {
	in := []string{
		"${var}",
	}
	with := "first\nsecond\r\nthird"
	replacer := New(in)
	replacer.Replace("var", with)
	out := replacer.GetOutput()
	exp := []string{
		"first\n",
		"second\r\n",
		"third",
	}
	if len(out) != len(exp) {
		t.Errorf("expected %d got %d", len(exp), len(out))
	}
	n := len(exp)
	if len(out) < n {
		n = len(out)
	}
	for i := 0; i < n; i++ {
		if out[i] != exp[i] {
			t.Errorf("\nexpected %s\ngot      %s", testutil.ReplaceControlChars(exp[i]), testutil.ReplaceControlChars(out[i]))
		}
	}
}

func TestMultiLineWithMultiLineIndentVar(t *testing.T) {
	in := []string{
		"\t${var}after",
	}
	with := "first\nsecond\r\nthird"
	replacer := New(in)
	replacer.Replace("var", with)
	out := replacer.GetOutput()
	exp := []string{
		"\tfirst\n",
		"\tsecond\r\n",
		"\tthirdafter",
	}
	if len(out) != len(exp) {
		t.Errorf("expected %d got %d", len(exp), len(out))
	}
	n := len(exp)
	if len(out) < n {
		n = len(out)
	}
	for i := 0; i < n; i++ {
		if out[i] != exp[i] {
			t.Errorf("\nexpected %s\ngot      %s", testutil.ReplaceControlChars(exp[i]), testutil.ReplaceControlChars(out[i]))
		}
	}
}

func TestMultiLineWithMultiLineIndentVar2(t *testing.T) {
	in := []string{
		"\tabc${var}after",
	}
	with := "first\nsecond\r\nthird"
	replacer := New(in)
	replacer.Replace("var", with)
	out := replacer.GetOutput()
	exp := []string{
		"\tabcfirst\n",
		"\tsecond\r\n",
		"\tthirdafter",
	}
	if len(out) != len(exp) {
		t.Errorf("expected %d got %d", len(exp), len(out))
	}
	n := len(exp)
	if len(out) < n {
		n = len(out)
	}
	for i := 0; i < n; i++ {
		if out[i] != exp[i] {
			t.Errorf("\nexpected %s\ngot      %s", testutil.ReplaceControlChars(exp[i]), testutil.ReplaceControlChars(out[i]))
		}
	}
}

func TestMultiLineWithMultiLineIndentVarAndSingleLineVar(t *testing.T) {
	in := []string{
		"\tabc${var}${var2}",
	}
	with := "first\nsecond\r\nthird"
	with2 := "one"
	replacer := New(in)
	replacer.Replace("var", with)
	replacer.Replace("var2", with2)
	out := replacer.GetOutput()
	exp := []string{
		"\tabcfirst\n",
		"\tsecond\r\n",
		"\tthirdone",
	}
	if len(out) != len(exp) {
		t.Errorf("expected %d got %d", len(exp), len(out))
	}
	n := len(exp)
	if len(out) < n {
		n = len(out)
	}
	for i := 0; i < n; i++ {
		if out[i] != exp[i] {
			t.Errorf("\nexpected %s\ngot      %s", testutil.ReplaceControlChars(exp[i]), testutil.ReplaceControlChars(out[i]))
		}
	}
}

func TestMultiLineWithManyMultiLineIndentVars(t *testing.T) {
	in := []string{
		"\tabc${var}${var2}",
	}
	with := "first\nsecond\r\nthird"
	with2 := "one\ntwo\r\n"
	replacer := New(in)
	replacer.Replace("var", with)
	replacer.Replace("var2", with2)
	out := replacer.GetOutput()
	exp := []string{
		"\tabcfirst\n",
		"\tsecond\r\n",
		"\tthirdone\n",
		"\ttwo\r\n",
		"\t",
	}
	if len(out) != len(exp) {
		t.Errorf("expected %d got %d", len(exp), len(out))
	}
	n := len(exp)
	if len(out) < n {
		n = len(out)
	}
	for i := 0; i < n; i++ {
		if out[i] != exp[i] {
			t.Errorf("\nexpected %s\ngot      %s", testutil.ReplaceControlChars(exp[i]), testutil.ReplaceControlChars(out[i]))
		}
	}
}

func TestMultiLineWithManyMultiLineIndentVarsAndLeadingNewline(t *testing.T) {
	in := []string{
		"\tabc${var}${var2}",
	}
	with := "\nfirst\nsecond\r\nthird"
	with2 := "\none\ntwo\r\n"
	replacer := New(in)
	replacer.Replace("var", with)
	replacer.Replace("var2", with2)
	out := replacer.GetOutput()
	exp := []string{
		"\tabc\n",
		"\tfirst\n",
		"\tsecond\r\n",
		"\tthird\n",
		"\tone\n",
		"\ttwo\r\n",
		"\t",
	}
	if len(out) != len(exp) {
		t.Errorf("expected %d got %d", len(exp), len(out))
	}
	n := len(exp)
	if len(out) < n {
		n = len(out)
	}
	for i := 0; i < n; i++ {
		if out[i] != exp[i] {
			t.Errorf("\nexpected %s\ngot      %s", testutil.ReplaceControlChars(exp[i]), testutil.ReplaceControlChars(out[i]))
		}
	}
}
