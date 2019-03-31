package wrappers

const (
	MaxUint = ^uint(0)          // max uint
	MinUint = 0                 // min uint
	MaxInt  = int(MaxUint >> 1) // max int
	MinInt  = -MaxInt - 1       // min int
)

type IType interface {
	V() interface{}
}

// C32 is rune 'class'
type C32 rune

func (a C32) V() interface{} {
	return a
}

// I32 is int 'class'
type I32 int

func (a I32) V() interface{} {
	return a
}
