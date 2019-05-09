package wrappers

import (
	"reflect"
	"sort"
)

// IArr2GArr :
func IArr2GArr(arr IArr) (garr []interface{}) {
	L := arr.Len()
	garr = make([]interface{}, L)
	for i := 0; i < L; i++ {
		garr[i] = arr.At(i)
	}
	return
}

// IArrEleIn :
func IArrEleIn(ele interface{}, arr IArr) bool {
	L := arr.Len()
	if L == 0 {
		return false
	}
	a0 := arr.At(0)
	tInput := reflect.TypeOf(ele)
	tArrEle := reflect.TypeOf(a0)
	PC(tInput != tArrEle, fEf("input element is <%v>, arr's element type is <%v>. Cannot compare!", tInput, tArrEle))

	for i := 0; i < L; i++ {
		a := arr.At(i)
		if ele == a {
			return true
		}
	}
	return false
}

// IArrSearchOne :
func IArrSearchOne(arr IArr, chk func(int, interface{}) (bool, interface{})) (ok bool, index int, rst interface{}) {
	L := arr.Len()
	for i := 0; i < L; i++ {
		a := arr.At(i)
		if yes, _ := chk(i, a); yes {
			ok, index, rst = true, i, a
			return
		}
	}
	return false, -1, nil
}

// IArrSearch :
func IArrSearch(arr IArr, chk func(int, interface{}) (bool, interface{})) (ok bool, indices []int, slice interface{}) {
	rst := GArr{}
	L := arr.Len()
	for i := 0; i < L; i++ {
		a := arr.At(i)
		if yes, _ := chk(i, a); yes {
			ok, indices, rst = true, append(indices, i), append(rst, a)
		}
	}
	return ok, indices, rst.Slice()
}

// IArrInsert :
func IArrInsert(arr IArr, it InsertType, chk func(int, interface{}) (bool, interface{})) (ok bool, indices []int, slice interface{}) {
	rst := GArr{}
	L := arr.Len()
	for i := 0; i < L; i++ {
		a := arr.At(i)
		switch it {
		case INSERT_BEFORE:
			if yes, item := chk(i, a); yes {
				ok, indices, rst = true, append(indices, i), append(rst, item)
			}
			rst = append(rst, a)
		case INSERT_AFTER:
			rst = append(rst, a)
			if yes, item := chk(i, a); yes {
				ok, indices, rst = true, append(indices, i), append(rst, item)
			}
		}
	}
	return ok, indices, rst.Slice()
}

// IArrRemove :
func IArrRemove(arr IArr, chk func(int, interface{}) (bool, interface{})) (ok bool, indices []int, dels, slice interface{}) {
	remain := GArr{}
	if ok, indices, dels = IArrSearch(arr, chk); ok {
		L := arr.Len()
		for i := 0; i < L; i++ {
			a := arr.At(i)
			if yes, _ := chk(i, a); yes {
				continue
			}
			remain = append(remain, a)
		}
	}
	return ok, indices, dels, remain.Slice()
}

// IArrReplace :
func IArrReplace(arr IArr, chk func(int, interface{}) (bool, interface{})) (ok bool, indices []int, old, slice interface{}) {
	rst := GArr{}
	if ok, indices, old = IArrSearch(arr, chk); ok {
		L := arr.Len()
		for i := 0; i < L; i++ {
			a := arr.At(i)
			if yes, item := chk(i, a); yes {
				rst = append(rst, item)
			} else {
				rst = append(rst, a)
			}
		}
	}
	return ok, indices, old, rst.Slice()
}

// IArrRmRep :
func IArrRmRep(arr IArr) (slice interface{}) {
	rst := GArr{}
	L := arr.Len()
OUTER:
	for i := 0; i < L; i++ {
		a := arr.At(i)
		for _, r := range rst {
			if a == r {
				continue OUTER
			}
		}
		rst = append(rst, a)
	}
	return rst.Slice()
}

// IArrFoldRep : ({"a", "b", "a"}, "[]")" => {[]"a", "b"}
//               ({"a", "b", "a"}, "[i]")" => {[2]"a", "b"}
func IArrFoldRep(arr IArr, style string) (slice interface{}) {
	switch arr.(type) {
	case Strs:
		mItemCnt := map[string]int{}
		L := arr.Len()
		for i := 0; i < L; i++ {
			a := arr.At(i).(string)
			if _, ok := mItemCnt[a]; ok {
				mItemCnt[a]++
			} else {
				mItemCnt[a] = 1
			}
		}

		rst := IArrRmRep(arr).([]string)
		L = len(rst)
		for i := 0; i < L; i++ {
			r := rst[i]
			n := mItemCnt[r]
			if n > 1 {
				switch style {
				case "[n]", "[i]":
					rst[i] = fSf("[%d]%v", n, r)
				case "[]", "":
					rst[i] = fSf("[]%v", r)
				}
			}
		}
		return rst

	default:
		panic("Now, only <Strs> can yield folded <[]string>")
	}
}

// IArrCtns :
func IArrCtns(arr IArr, others ...interface{}) bool {
	for _, oa := range others {
		if ok, _, _ := IArrSearch(arr, func(i int, a interface{}) (bool, interface{}) { return a == oa, "" }); !ok {
			return false
		}
	}
	return true
}

// IArrSeqCtns : ONLY apply to unique element array
func IArrSeqCtns(arr IArr, others ...interface{}) bool {
	ps := I32s{}
	for _, oa := range others {
		if ok, indices, _ := IArrSearch(arr, func(i int, a interface{}) (bool, interface{}) { return a == oa, "" }); ok {
			ps = append(ps, indices[0])
		} else {
			return false
		}
	}
	return sort.IntsAreSorted(ps)
}

// IArrIsSameEle :
func IArrIsSameEle(arr IArr) bool {
	L := arr.Len()
	for i := 0; i < L; i++ {
		if arr.At(i) != arr.At(0) {
			return false
		}
	}
	return true
}

// IArrIntersect :
func IArrIntersect(arr1, arr2 IArr) interface{} {
	L1, L2 := arr1.Len(), arr2.Len()
	rst := GArr{}
NEXT:
	for i := 0; i < L1; i++ {
		a1 := arr1.At(i)
		for j := 0; j < L2; j++ {
			if a1 == arr2.At(j) {
				rst = append(rst, a1)
				continue NEXT
			}
		}
	}
	return rst.Slice()
}

// IArrUnion :
func IArrUnion(arr1, arr2 IArr) interface{} {
	L1, L2 := arr1.Len(), arr2.Len()
	rst := make(GArr, L1, L1+L2)
	for i := 0; i < L1; i++ {
		rst[i] = arr1.At(i)
	}
NEXT:
	for j := 0; j < L2; j++ {
		a2 := arr2.At(j)
		for i := 0; i < L1; i++ {
			if a2 == arr1.At(i) {
				continue NEXT
			}
		}
		rst = append(rst, a2)
	}
	return rst.Slice()
}

// IArrStrJoinEx :
func IArrStrJoinEx(s1, s2 IArr, itemSep, strSep string) string {
	L1, L2 := s1.Len(), s2.Len()
	rmax, _ := Max(I32s{L1, L2})
	rmin, _ := Min(I32s{L1, L2}, "")
	Lmax, Lmin := rmax.(int), rmin.(int)

	rstArr := make([]string, Lmax)
	for i := 0; i < Lmin; i++ {
		rstArr[i] = fSf("%s%s%s", s1.At(i), itemSep, s2.At(i))
	}
	if L1 == L2 {
		return sJ(rstArr, strSep)
	}

	switch {
	case Lmax == L1:
		for i := L2; i < L1; i++ {
			rstArr[i] = s1.At(i).(string)
		}
	case Lmax == L2:
		for i := L1; i < L2; i++ {
			rstArr[i] = s2.At(i).(string)
		}
	}
	return sJ(rstArr, strSep)
}
