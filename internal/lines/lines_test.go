package lines

import (
	"strings"
	"testing"

	"github.com/x44/go-template-engine/internal/testutil"
)

func TestLines(t *testing.T) {
	stringsToTest := []string{
		"",
		"a",
		"\n",
		"\r\n",
		"a\n",
		"a\r\n",
		"a\n\n",
		"a\r\n\n",
		"a\n\r\n",
		"a\r\n\r\n",
		"\n\n",
		"\n\r\n\r\n",
		"a\nb",
		"a\r\nb",
		"a\nb\n",
		"a\r\nb\r\n",
		"\na\n\nb",
		"\na\n\nb\n",
		"\na\n\nb\r\n",
		"\na\n\r\nb",
		"\na\n\r\nb\n",
		"\na\n\r\nb\r\n",
	}
	for _, in := range stringsToTest {
		a := FromString(in)
		out := strings.Join(a, "")
		if out != in {
			t.Errorf("\nexpected %s\ngot      %s", testutil.ReplaceControlChars(in), testutil.ReplaceControlChars(out))
		}
	}
}

// func TestEmptyString(t *testing.T) {
// 	in := ""
// 	out := FromString(in)
// 	if len(out) != 0 {
// 		t.Errorf("expected 0 got %d", len(out))
// 	}
// 	in = "\n"
// 	out = FromString(in)
// 	if len(out) != 1 {
// 		t.Errorf("expected 1 got %d", len(out))
// 	}
// 	in = "\r\n"
// 	out = FromString(in)
// 	if len(out) != 1 {
// 		t.Errorf("expected 1 got %d", len(out))
// 	}
// }

// func TestSingleString(t *testing.T) {
// 	in := "a"
// 	out := FromString(in)
// 	if len(out) != 1 {
// 		t.Errorf("expected 1 got %d", len(out))
// 	}
// 	if out[0] != "a" {
// 		t.Errorf("expected 'a' got %s", out[0])
// 	}
// 	in = "a\n"
// 	out = FromString(in)
// 	if len(out) != 1 {
// 		t.Errorf("expected 1 got %d", len(out))
// 	}
// 	if out[0] != "a" {
// 		t.Errorf("expected 'a' got %s", out[0])
// 	}
// 	in = "a\r\n"
// 	out = FromString(in)
// 	if len(out) != 1 {
// 		t.Errorf("expected 1 got %d", len(out))
// 	}
// 	if out[0] != "a" {
// 		t.Errorf("expected 'a' got %s", out[0])
// 	}
// }

// func TestLineEndings(t *testing.T) {
// 	in := "a\nb\r\nc"
// 	out := FromString(in)
// 	if len(out) != 3 {
// 		t.Errorf("expected 3 got %d", len(out))
// 	}
// 	s := strings.Join(out, "")
// 	if s != "abc" {
// 		t.Errorf("expected 'abc' got %s", s)
// 	}
// }
