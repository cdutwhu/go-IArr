package wrappers

import (
	"reflect"
	"sort"
)

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
		pc(nCol != nColP && nColP != 0, fEf("All 1D Array should have same items"))
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

// SlcD2ToD1 :
func SlcD2ToD1(slc2d interface{}) (interface{}, bool) {
	v2d := reflect.ValueOf(slc2d)
	pc(v2d.Kind() != reflect.Slice, fEf("NOT A SLICE!"))
	l2 := v2d.Len()
	if l2 == 0 {
		return nil, false
	}
	slc2 := make([]interface{}, l2)
	for i := 0; i < l2; i++ {
		slc2[i] = v2d.Index(i).Interface()
	}
	var sRst reflect.Value
	for i, ele := range slc2 {
		v1d := reflect.ValueOf(ele)
		pc(v1d.Kind() != reflect.Slice, fEf("NOT A SLICE!"))
		if i == 0 {
			sType := reflect.TypeOf(v1d.Index(0).Interface())
			sRst = reflect.MakeSlice(reflect.SliceOf(sType), l2, l2)
		}
		if v1d.Len() != 1 {
			return nil, false
		}
		sRst.Index(i).Set(reflect.ValueOf(v1d.Index(0).Interface()))
	}
	return sRst.Interface(), true
}
