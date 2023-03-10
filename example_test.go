package trpd_test

import (
	"time"

	"suah.dev/trpd"
)

func ExampleWat() {
	trpd.Wat()
	time.Sleep(3 * time.Second)
	trpd.Wat()
}
