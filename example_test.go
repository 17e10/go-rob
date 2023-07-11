package rob_test

import (
	"github.com/17e10/go-rob"
)

func Example_crash() {
	s := "hello"
	b := rob.Bytes(s)
	b[0] = 'H' // crash: DO NOT change
}
