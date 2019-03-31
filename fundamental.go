package wrappers

func IType2Str(a IType) string {
	return fSf("%d", a.V())
}
