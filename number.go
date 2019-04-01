package wrappers

import "math"

func INum2Str(a INum) string {
	return fSf("%d", a.V())
}

func INumInArr(a INum, arr IArr) bool {
	L := arr.Len()
	for i := 0; i < L; i++ {
		if a.V() == arr.At(i) {
			return true
		}
	}
	return false
}

func INumNearest(a INum, arr IArr) (interface{}, int) {
	minDis, minIdx, aif64 := float64(MaxUint), -1, 0.0
	L := arr.Len()
	for i := 0; i < L; i++ {
		ai := arr.At(i)
		switch ai.(type) {
		case rune:
			aif64 = float64(ai.(rune))
		case int:
			aif64 = float64(ai.(int))
		case int64:
			aif64 = float64(ai.(int64))
		case float32:
			aif64 = float64(ai.(float32))
		case float64:
			aif64 = ai.(float64)
		default:
			panic("invalid number type in array")
		}

		if dis := math.Abs(a.DF() - aif64); dis < minDis {
			minDis, minIdx = dis, i
		}
	}
	return arr.At(minIdx), minIdx
}
