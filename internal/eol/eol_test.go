package eol

import "testing"

func TestEolGet(t *testing.T) {
	compare(t, NONE, Get(""))
	compare(t, NONE, Get("abc"))
	compare(t, LF, Get("\n"))
	compare(t, LF, Get("abc\n"))
	compare(t, CRLF, Get("\r\n"))
	compare(t, CRLF, Get("abc\r\n"))
}

func TestEolAdd(t *testing.T) {
	compare(t, "", Add("", NONE))
	compare(t, "\n", Add("", LF))
	compare(t, "\r\n", Add("", CRLF))
	compare(t, "", Add("", "xxx"))
}

func TestEolAddAllNONE(t *testing.T) {
	s := []string{
		"aaa",
		"bbb",
	}
	s = AddAll(s, NONE)
	compare(t, "aaa", s[0])
	compare(t, "bbb", s[1])
}

func TestEolAddAllToLF(t *testing.T) {
	s := []string{
		"aaa",
		"bbb",
	}
	s = AddAll(s, LF)
	compare(t, "aaa\n", s[0])
	compare(t, "bbb\n", s[1])
}

func TestEolAddAllToCRLF(t *testing.T) {
	s := []string{
		"aaa",
		"bbb",
	}
	s = AddAll(s, CRLF)
	compare(t, "aaa\r\n", s[0])
	compare(t, "bbb\r\n", s[1])
}

func TestEolChange(t *testing.T) {
	compare(t, "", Change("", NONE))
	compare(t, "", Change("\n", NONE))
	compare(t, "", Change("\r\n", NONE))

	compare(t, "abc", Change("abc", NONE))
	compare(t, "abc", Change("abc\n", NONE))
	compare(t, "abc", Change("abc\r\n", NONE))

	compare(t, "", Change("", LF))
	compare(t, "\n", Change("\n", LF))
	compare(t, "\n", Change("\r\n", LF))

	compare(t, "abc", Change("abc", LF))
	compare(t, "abc\n", Change("abc\n", LF))
	compare(t, "abc\n", Change("abc\r\n", LF))

	compare(t, "", Change("", CRLF))
	compare(t, "\r\n", Change("\n", CRLF))
	compare(t, "\r\n", Change("\r\n", CRLF))

	compare(t, "abc", Change("abc", CRLF))
	compare(t, "abc\r\n", Change("abc\n", CRLF))
	compare(t, "abc\r\n", Change("abc\r\n", CRLF))
}

func TestEolChangeAllToNONE(t *testing.T) {
	s := []string{
		"aaa\n",
		"bbb\r\n",
	}
	s = ChangeAll(s, NONE)
	compare(t, "aaa", s[0])
	compare(t, "bbb", s[1])
}

func TestEolChangeAllToLF(t *testing.T) {
	s := []string{
		"aaa\n",
		"bbb\r\n",
	}
	s = ChangeAll(s, LF)
	compare(t, "aaa\n", s[0])
	compare(t, "bbb\n", s[1])
}

func TestEolChangeAllToCRLF(t *testing.T) {
	s := []string{
		"aaa\n",
		"bbb\r\n",
	}
	s = ChangeAll(s, CRLF)
	compare(t, "aaa\r\n", s[0])
	compare(t, "bbb\r\n", s[1])
}

func compare(t *testing.T, want, got string) {
	if want != got {
		t.Errorf("wanted '%s' got '%s'", want, got)
	}
}
