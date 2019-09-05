package wrappers

import "testing"

func TestMap(t *testing.T) {
	m := map[string]string{"a": "b", "c": "d", "e": "f"}
	fPln(GetMapKeys(m).([]string))
	m1 := map[int]string{1: "B", 2: "D", 3: "F"}
	fPln(GetMapKeys(m1).([]int))
	fPln(GetMapKeys(EnCoDesc))

	k, v := GetMapKVs(m)
	K, V := k.([]string), v.([]string)
	fPln(K)
	fPln(V)

	k, v = GetMapKVs(m1)
	K1, V1 := k.([]int), v.([]string)
	fPln(K1)
	fPln(V1)
}
