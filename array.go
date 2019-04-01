package wrappers

import "sort"

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
