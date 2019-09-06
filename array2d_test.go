package wrappers

import "testing"

func TestSortIntArr2D(t *testing.T) {

	A2D := [][]int{
		[]int{3, 1, 2, 4},
		[]int{1, 2, 1, 1},
		[]int{2, 3, 1, 4},
		[]int{1, 3, 2, 0},
		[]int{2, 1, 1, 2},
		[]int{2, 1, 1, 1},
		[]int{1, 2, 3, 2},
	}

	col := Col(A2D, 0)
	fPln(col)

	fPln(Transpose(A2D))

	SortIntArr2D(A2D, "asc")
	fPln(A2D)

	/***************************************/

	fPln(SlcD2ToD1([][]string{{"abc"}, {"ghi"}, {"xxx"}}))
	fPln(SlcD2ToD1([][]string{{"abc"}, {"ghi"}, {"xxx"}, {}}))
	fPln(SlcD2ToD1([][]int{{1}, {2}, {3}}))
	fPln(SlcD2ToD1([][]int{{1}, {2}, {3, 4}}))

	if a, ok := SlcD2ToD1([][]int{{1}, {2}, {3}, {4}}); ok {
		b := a.([]int)
		fPln(b)
	}
}
