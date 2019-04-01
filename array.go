package wrappers

import "sort"

// IArrSearch :
func IArrSearch(arr IArr, chk func(int, interface{}) (bool, interface{})) (ok bool, indices []int, rst GArr) {
	L := arr.Len()
	for i := 0; i < L; i++ {
		a := arr.At(i)
		if yes, _ := chk(i, a); yes {
			ok, indices, rst = true, append(indices, i), append(rst, a)
		}
	}
	return
}

// IArrInsert :
func IArrInsert(arr IArr, it InsertType, chk func(int, interface{}) (bool, interface{})) (ok bool, indices []int, rst GArr) {
	L := arr.Len()
	for i := 0; i < L; i++ {
		a := arr.At(i)
		switch it {
		case IT_BEFORE:
			if yes, item := chk(i, a); yes {
				ok, indices, rst = true, append(indices, i), append(rst, item)
			}
			rst = append(rst, a)
		case IT_AFTER:
			rst = append(rst, a)
			if yes, item := chk(i, a); yes {
				ok, indices, rst = true, append(indices, i), append(rst, item)
			}
		}
	}
	return
}

// IArrRemove :
func IArrRemove(arr IArr, chk func(int, interface{}) (bool, interface{})) (ok bool, indices []int, dels, remain GArr) {
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
	return
}

// IArrReplace :
func IArrReplace(arr IArr, chk func(int, interface{}) (bool, interface{})) (ok bool, indices []int, old, rst GArr) {
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
	return
}

// IArrRmRep :
func IArrRmRep(arr IArr) (rst GArr) {
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
	return
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

// // AllAreIdentical :
// func (ga GArr) AllAreIdentical() bool {
// 	if ga.L() > 1 {
// 		for _, a := range ga {
// 			if ga[0] != a {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }

// // InterSec :
// func (ga GArr) InterSec(items ...interface{}) (r []interface{}) {
// NEXT:
// 	for _, g := range ga {
// 		for _, item := range items {
// 			if g == item {
// 				r = append(r, g)
// 				continue NEXT
// 			}
// 		}
// 	}
// 	return
// }

// // Union :
// func (ga GArr) Union(items ...interface{}) (r []interface{}) {
// 	for _, g := range ga {
// 		r = append(r, g)
// 	}
// 	rOri := make([]interface{}, len(r))
// 	copy(rOri, r)
// NEXT:
// 	for _, item := range items {
// 		for _, rItem := range rOri {
// 			if item == rItem {
// 				continue NEXT
// 			}
// 		}
// 		r = append(r, item)
// 	}
// 	return
// }

// /**********************************************************/

// // ToStrs :
// func (ga GArr) ToStrs() (r []string) {
// 	for _, a := range ga {
// 		r = append(r, a.(string))
// 	}
// 	return
// }

// // ToInts :
// func (ga GArr) ToInts() (r []int) {
// 	for _, a := range ga {
// 		r = append(r, a.(int))
// 	}
// 	return
// }

// // ToInt64s :
// func (ga GArr) ToInt64s() (r []int64) {
// 	for _, a := range ga {
// 		r = append(r, a.(int64))
// 	}
// 	return
// }
