package filter

import (
	"strings"
	"testing"
)

func TestEmpty(t *testing.T) {
	in := []string{}
	parser := New(in)
	out, err := parser.Process()
	if err != nil {
		t.Error(err)
	}
	if len(out) != 0 {
		t.Errorf("expected len 0 got %d", len(out))
	}
}

func TestNoConditions(t *testing.T) {
	in := []string{
		"a", "b", "c",
	}
	parser := New(in)
	out, err := parser.Process()
	if err != nil {
		t.Error(err)
	}
	if len(out) != 3 {
		t.Errorf("expected len 3 got %d", len(out))
	}
	s := strings.Join(out, "")
	if s != "abc" {
		t.Errorf("expected abc got %s", s)
	}
}

func TestErrorTooManyEndOfConditions(t *testing.T) {
	in := []string{
		"a",
		"//?flag",
		"b",
		"//-",
		"//-",
	}
	parser := New(in)
	_, err := parser.Process()
	if err == nil {
		t.Errorf("expected error 'end-of-condition without start-of-condition'")
	}
}

func TestErrorMissingEndOfCondition(t *testing.T) {
	in := []string{
		"a",
		"//?flag",
		"b",
	}
	parser := New(in)
	_, err := parser.Process()
	if err == nil {
		t.Errorf("expected error 'missing end-of-condition'")
	}
}

func TestSimpleCondition(t *testing.T) {
	in := []string{
		"      //?flagA",
		"a",
		" //-",
		"//!flagB",
		"b",
		"//-",
	}
	parser := New(in)
	out, err := parser.Process()
	if err != nil {
		t.Error(err)
	}
	s := strings.Join(out, "")
	if s != "b" {
		t.Errorf("expected b got %s", s)
	}

	parser.SetVar("flagA", true)
	out, err = parser.Process()
	if err != nil {
		t.Error(err)
	}
	s = strings.Join(out, "")
	if s != "ab" {
		t.Errorf("expected ab got %s", s)
	}

	parser.SetVar("flagA", false)
	parser.SetVar("flagB", true)
	out, err = parser.Process()
	if err != nil {
		t.Error(err)
	}
	s = strings.Join(out, "")
	if s != "" {
		t.Errorf("expected empty string got %s", s)
	}
}

func TestNestedCondition(t *testing.T) {
	in := []string{
		"//?flagOuter",
		"   //?flagA",
		"a",
		"   //-",
		"   //!flagA",
		"b",
		"   //-",
		"//-",
	}
	parser := New(in)
	out, err := parser.Process()
	if err != nil {
		t.Error(err)
	}
	s := strings.Join(out, "")
	if s != "" {
		t.Errorf("expected empty string got %s", s)
	}

	parser.SetVar("flagOuter", true)
	out, err = parser.Process()
	if err != nil {
		t.Error(err)
	}
	s = strings.Join(out, "")
	if s != "b" {
		t.Errorf("expected ab got %s", s)
	}

	parser.SetVar("flagA", true)
	out, err = parser.Process()
	if err != nil {
		t.Error(err)
	}
	s = strings.Join(out, "")
	if s != "a" {
		t.Errorf("expected empty string got %s", s)
	}
}
