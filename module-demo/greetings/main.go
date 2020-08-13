package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"prefix-path-also-remote-url1.package.com/bantla/greetings/stringlib"
)

func main() {
	fmt.Println(stringlib.ReverseRunes("Hello Go! - !oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
