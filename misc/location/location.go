package location

import (
	"strings"
	"unicode/utf8"
)

// Calculate returns 1-based [line:col] pair from byte offset into the buf
func Calculate(buf string, byteoffset int) (line, col int) {
	buf = strings.TrimPrefix(buf, "\xef\xbb\xbf") // skip bom

	col = 1

	cur, end := 0, len(buf)
	linestart := cur

	if byteoffset > end {
		byteoffset = end
	}

	for cur < byteoffset {
		c := buf[cur]
		cur++
		if c == '\n' {
			line++
			linestart = cur
		} else if c == '\r' {
			if cur < byteoffset && buf[cur] == '\n' {
				cur++
			}
			line++
			linestart = cur
		}
	}
	col = 1 + utf8.RuneCountInString(buf[linestart:byteoffset])

	return
}
