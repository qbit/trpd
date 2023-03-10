package trpd

import "testing"

func TestCount(t *testing.T) {
	Wat()
	Wat()
	Wat()
	Wat()
	if Count != 4 {
		t.Error("Count not incrementing!")
	}
}
