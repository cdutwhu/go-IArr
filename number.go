package wrappers

import "math"

// INum2Str :
func INum2Str(a INum) string {
	return fSf("%d", a.V())
}

// INumInArr :
func INumInArr(a INum, arr IArr) bool {
	L := arr.Len()
	for i := 0; i < L; i++ {
		if a.V() == arr.At(i) {
			return true
		}
	}
	return false
}

// INumNearest :
func INumNearest(a INum, arr IArr) (interface{}, int) {
	minDis, iMin, aif64 := float64(MaxUint), -1, 0.0
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

		if dis := math.Abs(a.F64() - aif64); dis < minDis {
			minDis, iMin = dis, i
		}
	}
	return arr.At(iMin), iMin
}

// Max :
func Max(arr IArr) (interface{}, int) {
	aif64, max, iMax := 0.0, float64(MinInt), -1
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

		if i == 0 || aif64 > max {
			max, iMax = aif64, i
		}
	}
	return arr.At(iMax), iMax
}

// Min : check index in advance. if it is -1, it is invalid result
func Min(arr IArr, scope string) (interface{}, int) {
	aif64, min, iMin := 0.0, float64(MaxInt), -1
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

		switch scope {
		case "":
			if i == 0 || aif64 < min {
				min, iMin = aif64, i
			}
		case "pos", "Pos", "POS", "positive", "Positive", "POSITIVE":
			if aif64 < min && aif64 > 0 {
				min, iMin = aif64, i
			}
		case "non-neg", "Non-Neg", "NON-NEG", "non-negative", "Non-Negative", "NON-NEGATIVE":
			if aif64 < min && aif64 >= 0 {
				min, iMin = aif64, i
			}
		}
	}
	if iMin >= 0 {
		return arr.At(iMin), iMin
	}
	return nil, -1
}
