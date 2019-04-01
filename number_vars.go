package wrappers

const (
	MaxUint = ^uint(0)          // max uint
	MinUint = 0                 // min uint
	MaxInt  = int(MaxUint >> 1) // max int
	MinInt  = -MaxInt - 1       // min int
)

type INum interface {
	V() interface{}
	DF() float64
}

// **********************************************************************

// C32 is rune 'class'
type C32 rune

func (a C32) V() interface{} {
	return a
}
func (a C32) DF() float64 {
	return float64(a)
}

// **********************************************************************

// I32 is int 'class'
type I32 int

func (a I32) V() interface{} {
	return a
}
func (a I32) DF() float64 {
	return float64(a)
}

// **********************************************************************

// I64 is int64 'class'
type I64 int64

func (a I64) V() interface{} {
	return a
}
func (a I64) DF() float64 {
	return float64(a)
}

// **********************************************************************

// F32 is float32 'class'
type F32 float32

func (a F32) V() interface{} {
	return a
}
func (a F32) DF() float64 {
	return float64(a)
}

// **********************************************************************

// F64 is float64 'class'
type F64 float64

func (a F64) V() interface{} {
	return a
}
func (a F64) DF() float64 {
	return float64(a)
}
