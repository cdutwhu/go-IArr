package wrappers

import (
	"reflect"
)

// GetMapKeys :
func GetMapKeys(m interface{}) interface{} {
	v := reflect.ValueOf(m)
	PC(v.Kind() != reflect.Map, fEf("not a map!"))
	keys := v.MapKeys()
	if L := len(keys); L > 0 {
		kType := reflect.TypeOf(keys[0].Interface())
		rst := reflect.MakeSlice(reflect.SliceOf(kType), L, L)
		for i, k := range keys {
			rst.Index(i).Set(reflect.ValueOf(k.Interface()))
		}
		return rst.Interface()
	}
	return nil
}
