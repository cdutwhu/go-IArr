package wrappers

import (
	"reflect"
)

// GetMapKeys :
func GetMapKeys(m interface{}) interface{} {
	v := reflect.ValueOf(m)
	PC(v.Kind() != reflect.Map, fEf("NOT A MAP!"))
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

// GetMapKVs :
func GetMapKVs(m interface{}) (interface{}, interface{}) {
	v := reflect.ValueOf(m)
	PC(v.Kind() != reflect.Map, fEf("NOT A MAP!"))
	keys := v.MapKeys()
	if L := len(keys); L > 0 {
		kType := reflect.TypeOf(keys[0].Interface())
		kRst := reflect.MakeSlice(reflect.SliceOf(kType), L, L)
		vType := reflect.TypeOf(v.MapIndex(keys[0]).Interface())
		vRst := reflect.MakeSlice(reflect.SliceOf(vType), L, L)
		for i, k := range keys {
			kRst.Index(i).Set(reflect.ValueOf(k.Interface()))
			vRst.Index(i).Set(reflect.ValueOf(v.MapIndex(k).Interface()))
		}
		return kRst.Interface(), vRst.Interface()
	}
	return nil, nil
}
