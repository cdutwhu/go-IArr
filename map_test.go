package wrappers

import "testing"

func TestMap(t *testing.T) {
	m := map[string]string{"a": "b", "c": "d", "e": "f"}
	fPln(GetMapKeys(m).([]string))
	m1 := map[int]string{1: "b", 2: "d", 3: "f"}
	fPln(GetMapKeys(m1).([]int))
	fPln(GetMapKeys(EnCoDesc))
}
