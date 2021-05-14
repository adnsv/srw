package xml

func ScrambleContent(s string) RawString {
	i, o, n := 0, 0, len(s)
	if n <= 0 {
		return ""
	}
	r := ""
	for i < n {
		c := s[i]
		i++
		switch c {
		case '&':
			r += s[o : i-1]
			r += "&amp;"
			o = i
		case '<':
			r += s[o : i-1]
			r += "&lt;"
			o = i
		case '>':
			r += s[o : i-1]
			r += "&gt;"
			o = i
		default:
			if c < ' ' {
				r += s[o : i-1]
				switch c {
				case '\r', '\n':
					r += "\n"
					if c == '\r' && i < n && s[i] == 'n' {
						i++
					}
				case '\t', '\v', '\f':
					r += string(c)
				default:
					buf := [5]byte{'&', '#', '0', '0', ';'}
					buf[2] += c / 10
					buf[3] += c % 10
					r += string(buf[:])
				}
				o = i
			}
		}
	}
	r += s[o:n]
	return RawString(r)
}

func ScrambleAttr(s string) RawString {
	// our regular attribute uses quotes, so we allow apostrophes to
	// remain unscrambled
	i, o, n := 0, 0, len(s)
	if n <= 0 {
		return ""
	}
	r := ""
	for i < n {
		c := s[i]
		i++
		switch c {
		case '&':
			r += s[o : i-1]
			r += "&amp;"
			o = i
		case '<':
			r += s[o : i-1]
			r += "&lt;"
			o = i
		case '>':
			r += s[o : i-1]
			r += "&gt;"
			o = i
		case '"':
			r += s[o : i-1]
			r += "&quot;"
			o = i
		default:
			if c < ' ' {
				r += s[o : i-1]
				switch c {
				case '\r', '\n':
					r += "\n"
					if c == '\r' && i < n && s[i] == 'n' {
						i++
					}
				case '\t', '\v', '\f':
					r += string(c)
				default:
					buf := [5]byte{'&', '#', '0', '0', ';'}
					buf[2] += c / 10
					buf[3] += c % 10
					r += string(buf[:])
				}
				o = i
			}
		}
	}
	r += s[o:n]
	return RawString(r)
}
