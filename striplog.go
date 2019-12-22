package main

import (
	"bytes"
	"io"
	"os"
	"regexp"
)

// ANSI matching regex magic from the internet
const ansi = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[-a-zA-Z\\d\\/#&.:=?%@~_]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PR-TZcf-ntqry=><~]))"

var re = regexp.MustCompile(ansi)

func main() {
	// just buffer it all, we're lazy. No rune readers here.
	var buf bytes.Buffer
	io.Copy(&buf, os.Stdin)
	out := re.ReplaceAll(buf.Bytes(), []byte{})
	io.Copy(os.Stdout, bytes.NewReader(out))
}
