package wrappers

import (
	"reflect"
)

type (
	InsertType int
)

const (
	INSERT_BEFORE, INSERT_AFTER InsertType = 1, 2
)

type IArr interface {
	Len() int
	Swap(int, int)
	Less(int, int) bool
	At(int) interface{}
	Set(int, interface{})
}

var FunSortLess func(interface{}, interface{}) bool

// ********************************************************************** interface{}

type GArr []interface{}

func (arr GArr) Len() int {
	return len(arr)
}
func (arr GArr) At(i int) interface{} {
	return arr[i]
}
func (arr GArr) Set(i int, v interface{}) {
	pc(i < 0 || i >= arr.Len(), fEf("Index out of range"))
	arr[i] = v
}
func (arr GArr) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
func (arr GArr) Less(i, j int) bool {
	return FunSortLess(arr[i], arr[j])
}
func (arr GArr) Slice() interface{} {
	if L := arr.Len(); L > 0 {
		eleType := reflect.TypeOf(arr.At(0))
		rst := reflect.MakeSlice(reflect.SliceOf(eleType), L, L)
		for i := range arr {
			rst.Index(i).Set(reflect.ValueOf(arr[i]))
		}
		return rst.Interface()
	}
	return nil
}

// ********************************************************************** string

type Strs []string

func (arr Strs) Len() int {
	return len(arr)
}
func (arr Strs) At(i int) interface{} {
	return arr[i]
}
func (arr Strs) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
func (arr Strs) Less(i, j int) bool {
	return FunSortLess(arr[i], arr[j])
}
func (arr Strs) Set(i int, v interface{}) {
	pc(i < 0 || i >= arr.Len(), fEf("Index out of range"))
	arr[i] = v.(string)
}

// ********************************************************************** int

type I32s []int

func (arr I32s) Len() int {
	return len(arr)
}
func (arr I32s) At(i int) interface{} {
	return arr[i]
}
func (arr I32s) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
func (arr I32s) Less(i, j int) bool {
	return FunSortLess(arr[i], arr[j])
}
func (arr I32s) Set(i int, v interface{}) {
	pc(i < 0 || i >= arr.Len(), fEf("Index out of range"))
	arr[i] = v.(int)
}

// ********************************************************************** int64

type I64s []int64

func (arr I64s) Len() int {
	return len(arr)
}
func (arr I64s) At(i int) interface{} {
	return arr[i]
}
func (arr I64s) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
func (arr I64s) Less(i, j int) bool {
	return FunSortLess(arr[i], arr[j])
}
func (arr I64s) Set(i int, v interface{}) {
	pc(i < 0 || i >= arr.Len(), fEf("Index out of range"))
	arr[i] = v.(int64)
}

// ********************************************************************** float

type F32s []float32

func (arr F32s) Len() int {
	return len(arr)
}
func (arr F32s) At(i int) interface{} {
	return arr[i]
}
func (arr F32s) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
func (arr F32s) Less(i, j int) bool {
	return FunSortLess(arr[i], arr[j])
}
func (arr F32s) Set(i int, v interface{}) {
	pc(i < 0 || i >= arr.Len(), fEf("Index out of range"))
	arr[i] = v.(float32)
}

// ********************************************************************** float64

type F64s []float64

func (arr F64s) Len() int {
	return len(arr)
}
func (arr F64s) At(i int) interface{} {
	return arr[i]
}
func (arr F64s) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
func (arr F64s) Less(i, j int) bool {
	return FunSortLess(arr[i], arr[j])
}
func (arr F64s) Set(i int, v interface{}) {
	pc(i < 0 || i >= arr.Len(), fEf("Index out of range"))
	arr[i] = v.(float64)
}

// ********************************************************************** rune

type C32s []rune

func (arr C32s) Len() int {
	return len(arr)
}
func (arr C32s) At(i int) interface{} {
	return arr[i]
}
func (arr C32s) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
func (arr C32s) Less(i, j int) bool {
	return FunSortLess(arr[i], arr[j])
}
func (arr C32s) Set(i int, v interface{}) {
	pc(i < 0 || i >= arr.Len(), fEf("Index out of range"))
	arr[i] = v.(rune)
}

// ********************************************************************** byte

type C8s []byte

func (arr C8s) Len() int {
	return len(arr)
}
func (arr C8s) At(i int) interface{} {
	return arr[i]
}
func (arr C8s) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
func (arr C8s) Less(i, j int) bool {
	return FunSortLess(arr[i], arr[j])
}
func (arr C8s) Set(i int, v interface{}) {
	pc(i < 0 || i >= arr.Len(), fEf("Index out of range"))
	arr[i] = v.(byte)
}
