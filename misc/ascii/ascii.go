package ascii

func IsAlpha(b byte) bool {
	return ('A' <= b && b <= 'Z') || ('a' <= b && b <= 'z')
}

func DecDigit(b byte) byte {
	return b - '0'
}

func HexDigit(b byte) byte {
	b -= '0'
	if b < 10 {
		return b
	}
	b -= ('A' - '0')
	if b < 6 {
		return b
	}
	return b - ('a' - 'A')
}

func IsDecDigit(b byte) bool {
	return DecDigit(b) < 10
}

func IsHexDigit(b byte) bool {
	return HexDigit(b) < 16
}
