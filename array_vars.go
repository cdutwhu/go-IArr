package wrappers

type (
	InsertType int
)

const (
	IT_BEFORE InsertType = 1
	IT_AFTER  InsertType = 2
)

type IArr interface {
	Len() int
	At(int) interface{}
	Swap(int, int)
	Less(int, int) bool
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
func (arr GArr) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
func (arr GArr) Less(i, j int) bool {
	// PC(FunSortLess == nil, fEf("Set func(interface{}, interface{}) bool before using sort"))
	return FunSortLess(arr[i], arr[j])
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
	// PC(FunSortLess == nil, fEf("Set func(interface{}, interface{}) bool before using sort"))
	return FunSortLess(arr[i], arr[j])
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
	// PC(FunSortLess == nil, fEf("Set func(interface{}, interface{}) bool before using sort"))
	return FunSortLess(arr[i], arr[j])
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
	// PC(FunSortLess == nil, fEf("Set func(interface{}, interface{}) bool before using sort"))
	return FunSortLess(arr[i], arr[j])
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
	// PC(FunSortLess == nil, fEf("Set func(interface{}, interface{}) bool before using sort"))
	return FunSortLess(arr[i], arr[j])
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
	// PC(FunSortLess == nil, fEf("Set func(interface{}, interface{}) bool before using sort"))
	return FunSortLess(arr[i], arr[j])
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
	// PC(FunSortLess == nil, fEf("Set func(interface{}, interface{}) bool before using sort"))
	return FunSortLess(arr[i], arr[j])
}
