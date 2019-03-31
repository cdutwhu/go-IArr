package wrappers

import "testing"

func Test2Str(t *testing.T) {
	a := C32('a')
	fPln(IType2Str(a))
	b := I32('b')
	fPln(IType2Str(b))
}
