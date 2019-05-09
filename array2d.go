package wrappers

import "sort"

// Col :
func Col(A2D [][]int, iCol int) (col []int) {
	nRow := len(A2D)
	col = make([]int, nRow)
	for i := 0; i < nRow; i++ {
		col[i] = A2D[i][iCol]
	}
	return
}

// MustMatrix : if unregular 2-D array, panic
func MustMatrix(A2D [][]int) (nRow, nCol int) {
	nRow, nCol, nColP := len(A2D), 0, 0
	for i := 0; i < nRow; i++ {
		nCol = len(A2D[i])
		PC(nCol != nColP && nColP != 0, fEf("All 1D Array should have same items"))
		nColP = nCol
	}
	return
}

// Transpose : LIKE Matrix transpose
func Transpose(A2D [][]int) (A2DT [][]int) {
	_, nCol := MustMatrix(A2D)
	A2DT = make([][]int, nCol)
	for iCol := 0; iCol < nCol; iCol++ {
		A2DT[iCol] = Col(A2D, iCol)
	}
	return
}

// SortIntArr2D : sortType : "ASC" / "DESC"
func SortIntArr2D(A2D [][]int, sortType string) {
	_, nCol := MustMatrix(A2D)
	sort.SliceStable(A2D, func(i, j int) bool {
		for k := 0; k < nCol; k++ {
			if A2D[i][k] == A2D[j][k] {
				continue
			} else {
				switch sortType {
				case "asc", "ASC":
					return A2D[i][k] < A2D[j][k]
				case "desc", "DESC":
					return A2D[i][k] > A2D[j][k]
				}
			}
		}
		return false
	})
}
