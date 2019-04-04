package wrappers

import (
	"sort"
	"testing"
)

func TestSearch(t *testing.T) {
	// arr := I32s{1, 2, 3, 4, 5}
	// arr := Strs{"a", "b", "c", "d", "b"}
	arr := F64s{1, 2, 3, 4, 5.5, 6, 7}
	ok1, index, rst1 := IArrSearchOne(arr, func(i int, a interface{}) (bool, interface{}) { return a == 5.5 || a == "b", "junk" })
	fPln(ok1, index, rst1.(float64))
	ok, indices, rst := IArrSearch(arr, func(i int, a interface{}) (bool, interface{}) { return i == 0 || i == 2 || a == "b", "junk" })
	fPln(ok, indices, rst.([]float64))
}

func TestInsert(t *testing.T) {
	// arr := I32s{1, 2, 3, 4, 5}
	arr := Strs{"a", "b", "c", "d"}
	fPln(IArrInsert(arr, INSERT_BEFORE, func(i int, a interface{}) (bool, interface{}) { return i == 0 || i == 1 || a == "d", a.(string) + "1" }))
	fPln(IArrInsert(arr, INSERT_AFTER, func(i int, a interface{}) (bool, interface{}) { return i == 0 || i == 1 || a == "d", a.(string) + "2" }))
}

func TestRemove(t *testing.T) {
	arr := I32s{1, 2, 3, 4, 5}
	// arr := Strs{"a", "b", "c", "d"}
	fPln(IArrRemove(arr, func(i int, a interface{}) (bool, interface{}) { return a == 4 || i == 0, "Junk" }))
}

func TestReplace(t *testing.T) {
	arr := I32s{1, 2, 3, 4, 5}
	// arr := Strs{"a", "b", "c", "d"}
	fPln(IArrReplace(arr, func(i int, a interface{}) (bool, interface{}) { return a == 4 || i == 0, a.(int) + 100 }))
}

func TestRmRep(t *testing.T) {
	arr := Strs{"abc", "::", "abc", "de", "de"}
	// arr := I32s{ 2, 4, 5, 6, 3, 2, 5, 2, 4 }
	fPln(IArrRmRep(arr))
}

func TestContain(t *testing.T) {
	arr := Strs{"::", "abc", "de", "mn", " "}
	fPln(IArrCtns(arr, "abc", "mn", "de", "::", " "))
}

func TestSeqContain(t *testing.T) {
	arr := Strs{"::", "abc", "de", "mn", ""}
	fPln(IArrSeqCtns(arr, "::", "de", "mn"))
}

func TestOrder(t *testing.T) {
	arr := Strs{":::", "abcd", "zt", "de", "mn", "A", ""}
	FunSortLess = func(left, right interface{}) bool { return len(left.(string)) < len(right.(string)) }
	sort.Sort(arr)
	fPln(arr)
	arr1 := I32s{2, 4, 5, 6, 3, 2, 5, 2, 4}
	FunSortLess = func(left, right interface{}) bool { return left.(int) < right.(int) }
	sort.Sort(arr1)
	fPln(arr1)
}

func TestSameEle(t *testing.T) {
	arr := Strs{"abc", "abc", "abc"}
	fPln(IArrIsSameEle(arr))
}

func TestInterSecANDUnion(t *testing.T) {
	arr1 := Strs{"::", "abcd", "def", "mn", "A", "C"}
	arr2 := Strs{"abcd", "def", "::", "A", "B", "D"}
	r := IArrIntersect(arr1, arr2)
	fPln(r.([]string))
	r = IArrUnion(arr1, arr2)
	fPln(r.([]string))
}
