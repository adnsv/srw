package multiline

import "strings"

// SplitAtEOLs splits a string at EOLs and removes white space at the end of
// each line
func SplitAtEOLs(s string) []string {
	ret := []string{}
	o, i, n := 0, 0, len(s)
	if n == 0 {
		return ret
	}

	for i < n {
		switch s[i] {
		case '\r':
			ret = append(ret, trimRHS(s[o:i]))
			i++
			if i < n && s[i] == '\n' {
				i++
			}
			o = i
		case '\n':
			ret = append(ret, trimRHS(s[o:i]))
			i++
			o = i
		default:
			i++
		}
	}
	return append(ret, trimRHS(s[o:i]))
}

func TrimTop(ss []string) []string {
	for len(ss) > 0 && len(ss[0]) == 0 {
		ss = ss[1:]
	}
	return ss
}

func TrimBottom(ss []string) []string {
	for len(ss) > 0 && len(ss[len(ss)-1]) == 0 {
		ss = ss[:len(ss)-1]
	}
	return ss
}

func TrimTopAndBottom(ss []string) []string {
	return TrimTop(TrimBottom(ss))
}

func Reindent(ss []string, prefix, indent string) {
	const tabs = 0
	const unknown = -1
	oind := unknown

	for _, s := range ss {
		if len(s) == 0 {
			continue
		}
		if s[0] == '\t' {
			oind = tabs
			break
		}
		n := countPrefixByte(s, ' ')
		if n == 0 {
			continue
		} else if oind == unknown {
			oind = n
		} else if n >= oind {
			// keep oind
		} else if n == 1 {
			oind = n
			break
		} else {
			oind = n
		}
	}

	if oind == tabs {
		common := -1
		for _, s := range ss {
			if len(s) == 0 {
				continue
			}
			n := countPrefixByte(s, '\t')
			if common < 0 {
				common = n
			} else if n < common {
				common = n
			}
		}
		for i, s := range ss {
			if len(s) == 0 {
				continue
			}
			n := countPrefixByte(s, '\t')
			ss[i] = prefix + strings.Repeat(indent, n-common) + s[n:]
		}
	} else if oind > 0 {
		common := -1
		for _, s := range ss {
			if len(s) == 0 {
				continue
			}
			n := countPrefixByte(s, ' ')
			if common < 0 {
				common = n
			} else if n < common {
				common = n
			}
		}
		for i, s := range ss {
			if len(s) == 0 {
				continue
			}
			n := countPrefixByte(s, ' ')
			l := (n - common) / oind
			n = common + l*oind
			ss[i] = prefix + strings.Repeat(indent, l) + s[n:]
		}
	} else {
		for i, s := range ss {
			if len(s) == 0 {
				continue
			}
			ss[i] = prefix + s
		}
	}
}

func countPrefixByte(s string, b byte) (n int) {
	for n < len(s) && s[n] == b {
		n++
	}
	return
}

func trimRHS(s string) string {
	n := len(s)
	for ; n > 0 && (s[n-1] == ' ' || s[n-1] == '\t' || s[n-1] == '\v' || s[n-1] == '\f'); n-- {
	}
	return s[:n]
}
