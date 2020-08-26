package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"prefix-path-also-remote-url1.package.com/bantla/greetings/stringlib"
	"prefix-path-also-remote-url1.package.com/bantla/greetings/stringlib/nested"
)

func main() {
	fmt.Println(stringlib.ReverseRunes("Hello Go! - !oG ,olleH"))
	fmt.Println(nested.ReverseFn("Hello Go! - !oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
