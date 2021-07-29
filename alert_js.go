//+build js

package ebitentest

import "syscall/js"

func alert(msg string) {
	js.Global().Call("alert", msg)
}
