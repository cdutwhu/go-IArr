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
	r, i := INumNearest(b, bb)
	fPln(r.(int), i)
}

func TestMin(t *testing.T) {
	bb := I32s{-100, 231, 266, 6, 7, 11, 1, 25, -2}
	if i, p := Min(bb, ""); p >= 0 {
		fPln(i.(int), p)
	}
	max, p := Max(bb)
	fPln(max.(int), p)
}
