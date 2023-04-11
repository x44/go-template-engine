package replacer

import (
	"testing"

	"github.com/x44/go-template-engine/internal/lines"
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
	if len(out) != 3 {
		t.Errorf("expected 3 got %d", len(out))
	}
	if out[0] != "first" {
		t.Errorf("expected first got %s", out[0])
	}
	if out[1] != "second" {
		t.Errorf("expected second got %s", out[1])
	}
	if out[2] != "third" {
		t.Errorf("expected third got %s", out[2])
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
	if len(out) != 4 {
		t.Errorf("expected 4 got %d", len(out))
	}
	if out[0] != "\tfirst" {
		t.Errorf("expected \\tfirst got %s", out[0])
	}
	if out[1] != "\tsecond" {
		t.Errorf("expected \\tsecond got %s", out[1])
	}
	if out[2] != "\tthird" {
		t.Errorf("expected \\tthird got %s", out[2])
	}
	if out[3] != "after" {
		t.Errorf("expected after got %s", out[3])
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
	if len(out) != 5 {
		t.Errorf("expected 5 got %d", len(out))
	}
	if out[0] != "\tabc" {
		t.Errorf("expected \\tabc got %s", out[1])
	}
	if out[1] != "\tfirst" {
		t.Errorf("expected \\tfirst got %s", out[1])
	}
	if out[2] != "\tsecond" {
		t.Errorf("expected \\tsecond got %s", out[2])
	}
	if out[3] != "\tthird" {
		t.Errorf("expected \\tthird got %s", out[3])
	}
	if out[4] != "after" {
		t.Errorf("expected after got %s", out[4])
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
	if len(out) != 5 {
		t.Errorf("expected 5 got %d", len(out))
	}
	if out[0] != "\tabc" {
		t.Errorf("expected \\tabc got %s", out[1])
	}
	if out[1] != "\tfirst" {
		t.Errorf("expected \\tfirst got %s", out[1])
	}
	if out[2] != "\tsecond" {
		t.Errorf("expected \\tsecond got %s", out[2])
	}
	if out[3] != "\tthird" {
		t.Errorf("expected \\tthird got %s", out[3])
	}
	if out[4] != "one" {
		t.Errorf("expected one got %s", out[4])
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
	if len(out) != 6 {
		t.Errorf("expected 6 got %d", len(out))
	}
	if out[0] != "\tabc" {
		t.Errorf("expected \\tabc got %s", out[1])
	}
	if out[1] != "\tfirst" {
		t.Errorf("expected \\tfirst got %s", out[1])
	}
	if out[2] != "\tsecond" {
		t.Errorf("expected \\tsecond got %s", out[2])
	}
	if out[3] != "\tthird" {
		t.Errorf("expected \\tthird got %s", out[3])
	}
	if out[4] != "one" {
		t.Errorf("expected one got %s", out[4])
	}
	if out[5] != "two" {
		t.Errorf("expected two got %s", out[5])
	}
}
