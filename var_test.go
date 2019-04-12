package wrappers

import "testing"

func TestConst(t *testing.T) {
	fPln(QSingle, QDouble)
	fPln(BRound, BBox, BCurly, BAngle)
	fPln(ASCII, UTF8)
}
