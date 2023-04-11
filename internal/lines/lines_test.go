package lines

import (
	"strings"
	"testing"
)

func TestEmptyString(t *testing.T) {
	in := ""
	out := FromString(in)
	if len(out) != 0 {
		t.Errorf("expected 0 got %d", len(out))
	}
	in = "\n"
	out = FromString(in)
	if len(out) != 1 {
		t.Errorf("expected 1 got %d", len(out))
	}
	in = "\r\n"
	out = FromString(in)
	if len(out) != 1 {
		t.Errorf("expected 1 got %d", len(out))
	}
}

func TestSingleString(t *testing.T) {
	in := "a"
	out := FromString(in)
	if len(out) != 1 {
		t.Errorf("expected 1 got %d", len(out))
	}
	if out[0] != "a" {
		t.Errorf("expected 'a' got %s", out[0])
	}
	in = "a\n"
	out = FromString(in)
	if len(out) != 1 {
		t.Errorf("expected 1 got %d", len(out))
	}
	if out[0] != "a" {
		t.Errorf("expected 'a' got %s", out[0])
	}
	in = "a\r\n"
	out = FromString(in)
	if len(out) != 1 {
		t.Errorf("expected 1 got %d", len(out))
	}
	if out[0] != "a" {
		t.Errorf("expected 'a' got %s", out[0])
	}
}

func TestLineEndings(t *testing.T) {
	in := "a\nb\r\nc"
	out := FromString(in)
	if len(out) != 3 {
		t.Errorf("expected 3 got %d", len(out))
	}
	s := strings.Join(out, "")
	if s != "abc" {
		t.Errorf("expected 'abc' got %s", s)
	}
}
