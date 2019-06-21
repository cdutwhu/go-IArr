package wrappers

import (
	"regexp"
	"strconv"
	"unicode"
	"unicode/utf8"
)

// ASCIIToOri :
func ASCIIToOri(astr string) string {
	start, end, base := "^", "$", 36
	r, _ := regexp.Compile(`\` + start + `[a-z0-9]?[a-z0-9]?[a-z0-9]?[a-z0-9]\` + end)
	for codedarea := r.FindString(astr); codedarea != ""; codedarea = r.FindString(astr) {
		code := codedarea[1 : len(codedarea)-1]
		utf8rune, _ := strconv.ParseInt(code, base, 64)
		astr = sRplc(astr, codedarea, string(utf8rune), 1)
	}
	return astr
}

// UTF8ToASCII : Convert UTF8 string to ascii string
func UTF8ToASCII(str string) (ascii bool, astr string) {
	start, end, base := "^", "$", 36
	I, L, mPosUTF8Len := 0, len(str), map[int]int{}
	for i := 0; i < L; i++ {
		if str[i] > unicode.MaxASCII {
			if str[i]&0xC0 == 0xC0 {
				I = i
				mPosUTF8Len[I] = 0
			}
			if str[i]&0x80 == 0x80 {
				mPosUTF8Len[I]++
			}
		}
	}
	if len(mPosUTF8Len) == 0 {
		return true, str
	}

	mPosRuneStr, mPosUTF8, mCharUniB36 := map[int]string{}, map[int]string{}, map[string]string{}
	for k, v := range mPosUTF8Len {
		r, _ := utf8.DecodeRune([]byte(str[k : k+v]))
		mPosRuneStr[k] = strconv.FormatInt((int64)(r), base)
		mPosUTF8[k] = string(r)
	}
	for k, v := range mPosRuneStr {
		runeRplc := start + v + end
		mCharUniB36[mPosUTF8[k]] = runeRplc
	}
	astr = str
	for k, v := range mCharUniB36 {
		astr = sRplcAll(astr, k, v)
	}
	return false, astr
}
