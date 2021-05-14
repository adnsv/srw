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
.st0{fill:none;stroke:#000000;stroke-width:0.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:10;}
.st1{fill:none;stroke:#000000;stroke-width:0.25;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:10;}
.st2{fill:#304499;stroke:#FFFFFF;stroke-width:0.5;stroke-miterlimit:10;}
.st3{fill:#FFFFFF;stroke:#304499;stroke-width:0.3502;stroke-miterlimit:10;}
.st4{fill:#010101;}
.st5{font-family:'ArialMT';}
.st6{font-size:8px;}
`
