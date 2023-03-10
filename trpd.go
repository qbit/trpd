package trpd

import "fmt"

// Count is the number of times we have called Wat.
var Count = 0

// Wat prints "HERE" and adds a trailing digit for every call.
// This is useful for debugging!
func Wat() {
	if Count == 0 {
		fmt.Println("HERE")
	} else {
		fmt.Printf("HERE%d\n", Count)
	}

	Count = Count + 1
}
