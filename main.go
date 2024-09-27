//go:build js && wasm
// +build js,wasm

package main

import (
	"syscall/js"
)

/* func Compute(a int, b bool) int {
	if b {
		return a * 2
	}
	return a - 1
} */

func compute(this js.Value, p []js.Value) interface{} {
	a := p[0].Int()
	b := p[1].Bool()

	var c int

	if b {
		c = a * 2
	} else {
		c = a - 1
	}

	return js.ValueOf(c)
}

func main() {
	// expose functions to JavaScript
	js.Global().Set("compute", js.FuncOf(compute))

	// keep Go program running
	select {}
}
