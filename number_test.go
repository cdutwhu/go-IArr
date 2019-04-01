package wrappers

import "testing"

func Test2Str(t *testing.T) {
	a := C32('a')
	fPln(INum2Str(a))
	b := I32('b')
	fPln(INum2Str(b))
}

func TestNearest(t *testing.T) {
	b := I32(25)
	bb := I32s{231, 26, 6, 7, 11, 25}
	fPln(INumNearest(b, bb))
}
