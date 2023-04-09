package eol

const (
	NONE = ""
	LF   = "\n"
	CRLF = "\r\n"
)

// Returns the EOL of s
func Get(s string) string {
	l := len(s)
	if l == 0 {
		return NONE
	}
	if l >= 2 && s[l-2] == '\r' && s[l-1] == '\n' {
		return CRLF
	}
	if l >= 1 && s[l-1] == '\n' {
		return LF
	}
	return NONE
}

// Appends EOL to s without checking if s already has an EOL
func Add(s string, eol string) string {
	if eol == NONE || !(eol == LF || eol == CRLF) {
		return s
	}
	return s + eol
}

// Changes the EOL of s but only if s has an EOL
func Change(s string, eol string) string {
	cur := Get(s)
	switch cur {
	case NONE:
		return s
	case LF:
		if eol == cur {
			return s
		}
		return s[:len(s)-1] + eol
	case CRLF:
		if eol == cur {
			return s
		}
		return s[:len(s)-2] + eol
	}
	return s
}

// Appends EOL to all elemnts in a without checking if elments already have an EOL
func AddAll(a []string, eol string) []string {
	out := make([]string, len(a))
	for i, s := range a {
		out[i] = Add(s, eol)
	}
	return out
}

// Changes the EOL of all elements in a but only if s has an EOL
func ChangeAll(a []string, eol string) []string {
	out := make([]string, len(a))
	for i, s := range a {
		out[i] = Change(s, eol)
	}
	return out
}
