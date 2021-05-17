package css

import (
	"log"
	"testing"
)

func TestRead(t *testing.T) {
	read(example01)
}

func read(buf string) {
	err := Read(buf)
	if err != nil {
		log.Fatal(err)
	}
}

const example01 = `
@IDENTIFIER (RULE);
@IDENTIFIER blah { blah }
.st0{fill:none;stroke:#000000;stroke-width:0.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:10;}
.st1{fill:none;stroke:#000000;stroke-width:0.25;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:10;}
.st2{fill:#304499;stroke:#FFFFFF;stroke-width:0.5;stroke-miterlimit:10;}
.st3{fill:#FFFFFF;stroke:#304499;stroke-width:0.3502;stroke-miterlimit:10;}
.st4{fill:#010101;}
.st5{font-family:'ArialMT';}
.st6{font-size:8px;}
`

const example01_tokens = `
Process exiting with code: 0
API server listening at: 127.0.0.1:38102
S (line: 1, column: 1): "\n"
ATKEYWORD (line: 2, column: 1): "@IDENTIFIE"...
S (line: 2, column: 12): " "
CHAR (line: 2, column: 13): "("
IDENT (line: 2, column: 14): "RULE"
CHAR (line: 2, column: 18): ")"
CHAR (line: 2, column: 19): ";"
S (line: 2, column: 20): "\n"
CHAR (line: 3, column: 1): "."
IDENT (line: 3, column: 2): "st0"
CHAR (line: 3, column: 5): "{"
IDENT (line: 3, column: 6): "fill"
CHAR (line: 3, column: 10): ":"
IDENT (line: 3, column: 11): "none"
CHAR (line: 3, column: 15): ";"
IDENT (line: 3, column: 16): "stroke"
CHAR (line: 3, column: 22): ":"
HASH (line: 3, column: 23): "#000000"
CHAR (line: 3, column: 30): ";"
IDENT (line: 3, column: 31): "stroke-wid"...
CHAR (line: 3, column: 43): ":"
NUMBER (line: 3, column: 44): "0.5"
CHAR (line: 3, column: 47): ";"
IDENT (line: 3, column: 48): "stroke-lin"...
CHAR (line: 3, column: 62): ":"
IDENT (line: 3, column: 63): "round"
CHAR (line: 3, column: 68): ";"
IDENT (line: 3, column: 69): "stroke-lin"...
CHAR (line: 3, column: 84): ":"
IDENT (line: 3, column: 85): "round"
CHAR (line: 3, column: 90): ";"
IDENT (line: 3, column: 91): "stroke-mit"...
CHAR (line: 3, column: 108): ":"
NUMBER (line: 3, column: 109): "10"
CHAR (line: 3, column: 111): ";"
CHAR (line: 3, column: 112): "}"
S (line: 3, column: 113): "\n"
CHAR (line: 4, column: 1): "."
IDENT (line: 4, column: 2): "st1"
CHAR (line: 4, column: 5): "{"
IDENT (line: 4, column: 6): "fill"
CHAR (line: 4, column: 10): ":"
IDENT (line: 4, column: 11): "none"
CHAR (line: 4, column: 15): ";"
IDENT (line: 4, column: 16): "stroke"
CHAR (line: 4, column: 22): ":"
HASH (line: 4, column: 23): "#000000"
CHAR (line: 4, column: 30): ";"
IDENT (line: 4, column: 31): "stroke-wid"...
CHAR (line: 4, column: 43): ":"
NUMBER (line: 4, column: 44): "0.25"
CHAR (line: 4, column: 48): ";"
IDENT (line: 4, column: 49): "stroke-lin"...
CHAR (line: 4, column: 63): ":"
IDENT (line: 4, column: 64): "round"
CHAR (line: 4, column: 69): ";"
IDENT (line: 4, column: 70): "stroke-lin"...
CHAR (line: 4, column: 85): ":"
IDENT (line: 4, column: 86): "round"
CHAR (line: 4, column: 91): ";"
IDENT (line: 4, column: 92): "stroke-mit"...
CHAR (line: 4, column: 109): ":"
NUMBER (line: 4, column: 110): "10"
CHAR (line: 4, column: 112): ";"
CHAR (line: 4, column: 113): "}"
S (line: 4, column: 114): "\n"
CHAR (line: 5, column: 1): "."
IDENT (line: 5, column: 2): "st2"
CHAR (line: 5, column: 5): "{"
IDENT (line: 5, column: 6): "fill"
CHAR (line: 5, column: 10): ":"
HASH (line: 5, column: 11): "#304499"
CHAR (line: 5, column: 18): ";"
IDENT (line: 5, column: 19): "stroke"
CHAR (line: 5, column: 25): ":"
HASH (line: 5, column: 26): "#FFFFFF"
CHAR (line: 5, column: 33): ";"
IDENT (line: 5, column: 34): "stroke-wid"...
CHAR (line: 5, column: 46): ":"
NUMBER (line: 5, column: 47): "0.5"
CHAR (line: 5, column: 50): ";"
IDENT (line: 5, column: 51): "stroke-mit"...
CHAR (line: 5, column: 68): ":"
NUMBER (line: 5, column: 69): "10"
CHAR (line: 5, column: 71): ";"
CHAR (line: 5, column: 72): "}"
S (line: 5, column: 73): "\n"
CHAR (line: 6, column: 1): "."
IDENT (line: 6, column: 2): "st3"
CHAR (line: 6, column: 5): "{"
IDENT (line: 6, column: 6): "fill"
CHAR (line: 6, column: 10): ":"
HASH (line: 6, column: 11): "#FFFFFF"
CHAR (line: 6, column: 18): ";"
IDENT (line: 6, column: 19): "stroke"
CHAR (line: 6, column: 25): ":"
HASH (line: 6, column: 26): "#304499"
CHAR (line: 6, column: 33): ";"
IDENT (line: 6, column: 34): "stroke-wid"...
CHAR (line: 6, column: 46): ":"
NUMBER (line: 6, column: 47): "0.3502"
CHAR (line: 6, column: 53): ";"
IDENT (line: 6, column: 54): "stroke-mit"...
CHAR (line: 6, column: 71): ":"
NUMBER (line: 6, column: 72): "10"
CHAR (line: 6, column: 74): ";"
CHAR (line: 6, column: 75): "}"
S (line: 6, column: 76): "\n"
CHAR (line: 7, column: 1): "."
IDENT (line: 7, column: 2): "st4"
CHAR (line: 7, column: 5): "{"
IDENT (line: 7, column: 6): "fill"
CHAR (line: 7, column: 10): ":"
HASH (line: 7, column: 11): "#010101"
CHAR (line: 7, column: 18): ";"
CHAR (line: 7, column: 19): "}"
S (line: 7, column: 20): "\n"
CHAR (line: 8, column: 1): "."
IDENT (line: 8, column: 2): "st5"
CHAR (line: 8, column: 5): "{"
IDENT (line: 8, column: 6): "font-famil"...
CHAR (line: 8, column: 17): ":"
STRING (line: 8, column: 18): "'ArialMT'"
CHAR (line: 8, column: 27): ";"
CHAR (line: 8, column: 28): "}"
S (line: 8, column: 29): "\n"
CHAR (line: 9, column: 1): "."
IDENT (line: 9, column: 2): "st6"
CHAR (line: 9, column: 5): "{"
IDENT (line: 9, column: 6): "font-size"
CHAR (line: 9, column: 15): ":"
DIMENSION (line: 9, column: 16): "8px"
CHAR (line: 9, column: 19): ";"
CHAR (line: 9, column: 20): "}"
S (line: 9, column: 21): "\n"
`
