package svg_test

import (
	"log"
	"os"

	"github.com/adnsv/srw/svg"
	"github.com/adnsv/srw/xml"
)

const sampleSVG = `<?xml version="1.0" encoding="utf-8"?>
<!-- Generator: Adobe Illustrator 15.0.0, SVG Export Plug-In . SVG Version: 6.00 Build 0) -->
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1 Tiny//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11-tiny.dtd">
<svg version="1.1" baseProfile="tiny" id="Layer_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px" width="25.2px" height="25.2px" viewBox="0 0 25.2 25.2" xml:space="preserve">
	<path fill="#231F20" d="M12.601,0.9c6.451,0,11.7,5.249,11.7,11.7c0,6.451-5.249,11.701-11.7,11.701S0.9,19.051,0.9,12.6
	C0.9,6.149,6.149,0.9,12.601,0.9 M12.601,0C5.642,0,0,5.641,0,12.6c0,6.958,5.642,12.6,12.601,12.6s12.6-5.642,12.6-12.6
	C25.2,5.641,19.56,0,12.601,0L12.601,0z" />
	<polygon points="19.801,11.16 19.801,14.04 9.9,14.04 9.9,16.2 4.5,12.6 9.9,9 9.9,11.16 " />
</svg>`

func Example_readwrite() {
	s, err := svg.FromString(sampleSVG)
	if err != nil {
		log.Fatal(err)
	}

	svg.Write(os.Stdout, s, xml.WriterConfig{Indent: xml.Indent2Spaces})
}
