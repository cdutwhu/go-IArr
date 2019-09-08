package wrappers

import (
	"strings"
	"unicode/utf8"

	"github.com/google/uuid"
)

// Str is UTF-8 string 'class'
type Str string

// V : get string value
func (s Str) V() string {
	return string(s)
}

// SetEnC : default is UTF8, if ASCII is suitable, switch to ASCII
func (s Str) SetEnC() string {
	EnCoType = UTF8
	EnCoType = IF(s.L() == len(s.V()), ASCII, EnCoType).(EnCo)
	return EnCoDesc[EnCoType]
}

// L : get string characters' count                          $
func (s Str) L() int {
	if EnCoType == ASCII {
		return len(s.V())
	}
	return utf8.RuneCountInString(s.V()) // maybe faster than "len([]rune(s))"
}

// S : like [a:b], a included, b excluded                    $
func (s Str) S(from, to int, strlen ...int) Str {
	L := 0
	if len(strlen) > 0 {
		L = strlen[0]
	} else {
		L = s.L()
	}

	d := ALL - to
	to = IF(d < TORANGE, L-d, to).(int)
	// to = IF(to < L, to, L).(int)               ** DO NOT ignore parameter error **
	// from = IF(from >= 0, from, 0).(int)        ** DO NOT ignore parameter error **

	PC(from < 0 || from > L, fEf("Invalid parameters <from>"))
	PC(to < 0 || to > L, fEf("Invalid parameters <to>"))
	PC(from > to, fEf("Invalid parameters, <from > to>"))
	if from == to {
		return Str("")
	}

	if EnCoType == ASCII {
		return s[from:to]
	}

	i, startByteIdx := 0, 0
	for j := range s {
		if i == from {
			startByteIdx = j
		}
		if i == to {
			return s[startByteIdx:j]
		}
		i++
	}
	return s[startByteIdx:]

	// return Str([]rune(s.V())[from:to]) //            ** slow **
}

// STo :
func (s Str) STo(to string) Str {
	if p := s.Idx(to); p >= 0 {
		return s.S(0, p)
	}
	return s
}

// SegRep : replace a section of a string                    $
func (s Str) SegRep(from, to int, seg string) Str {
	L := s.L()
	left, right := s.S(0, from, L), s.S(to, ALL, L)
	return Str(left.V() + seg + right.V())
}

// Replace : replace all old pieces to new pieces
func (s Str) Replace(old, new string) Str {
	return Str(sRplc(s.V(), old, new, -1))
}

// C : the p position's character                            $
func (s Str) C(p int) rune {
	d, L := LAST-p, s.L()
	p = IF(d < TORANGE, L-d-1, p).(int)
	//p = IF(p < l, p, l-1).(int)                         // ** DO NOT ignore parameter error **
	PC(p < 0 || p >= L, fEf("Invalid parameters <p>"))

	if EnCoType == ASCII {
		return rune(s[p])
	}

	return []rune(s.S(p, p+1, L))[0]
}

// Idx : strings.Index                                       $
func (s Str) Idx(sub string) int {
	if pByte := strings.Index(s.V(), sub); pByte >= 0 {
		return s[:pByte].L()
	}
	return -1
}

// LIdx : strings.LastIndex                                  $
func (s Str) LIdx(sub string) int {
	if pByte := strings.LastIndex(s.V(), sub); pByte >= 0 {
		return s[:pByte].L()
	}
	return -1
}

// HP : HasPrefix                                            $
func (s Str) HP(prefix string) bool {
	return strings.HasPrefix(s.V(), prefix)
}

// HS : HasSuffix                                            $
func (s Str) HS(suffix string) bool {
	return strings.HasSuffix(s.V(), suffix)
}

// T : strings.Trim                                          $
func (s Str) T(cutset string) Str {
	return Str(strings.Trim(s.V(), cutset))
}

// TL : strings.TrimLeft                                     $
func (s Str) TL(cutset string) Str {
	return Str(strings.TrimLeft(s.V(), cutset))
}

// TR : strings.TrimRight                                    $
func (s Str) TR(cutset string) Str {
	return Str(strings.TrimRight(s.V(), cutset))
}

// ToInt64 :                                                 $
func (s Str) ToInt64() int64 {
	return Must(sc2Int(s.V(), 10, 64)).(int64)
}

// ToInt :                                                   $
func (s Str) ToInt() int {
	return int(s.ToInt64())
}

// ToLower :
func (s Str) ToLower() Str {
	return Str(strings.ToLower(s.V()))
}

// ToUpper :
func (s Str) ToUpper() Str {
	return Str(strings.ToUpper(s.V()))
}

// ******************************************************************************************************* //

// DefValue : if is empty, assign it with input string, otherwise keep its value             $
func (s Str) DefValue(def string) Str {
	return IF(s.V() == "", Str(def), s).(Str)
}

// Repeat : e.g. "ABC"(2) => "ABCABC"                                                        $
func (s Str) Repeat(n int, sep string) Str {
	strs := []string{}
	for i := 0; i < n; i++ {
		strs = append(strs, s.V())
	}
	return Str(strings.Join(strs, sep))
}

// Contains :
func (s Str) Contains(substr string) bool {
	return sCtns(s.V(), substr)
}

// HasAny : e.g. "ABC"('A', 'M') => true                                                     $
func (s Str) HasAny(cks ...rune) bool {
	for _, c := range s {
		for _, ck := range cks {
			if c == ck {
				return true
			}
		}
	}
	return false
}

// IsMadeFrom : e.g. "ABC"('C','B','A','D') => true                                          $
func (s Str) IsMadeFrom(chars ...rune) bool {
NEXT:
	for _, c := range s {
		for _, ck := range chars {
			if c == ck {
				continue NEXT
			}
		}
		return false
	}
	return true
}

// InArr : check if at least one same value exists in string array                           $
func (s Str) InArr(arr ...string) (bool, int) {
	for i, a := range arr {
		if s.V() == a {
			return true, i
		}
	}
	return false, -1
}

// InMapSIKeys : check if at least one same value exists in string-key map                   $
func (s Str) InMapSIKeys(m map[string]int) (bool, int) {
	for k, v := range m {
		if s.V() == k {
			return true, v
		}
	}
	return false, -1
}

// InMapSSValues : check if at least a same value exists in string-value map                 $
func (s Str) InMapSSValues(m map[string]string) (bool, string) {
	for k, v := range m {
		if s.V() == v {
			return true, k
		}
	}
	return false, ""
}

// BeCoveredInMapSIKeys : check if at least one map(string)key value can cover the calling string  &
func (s Str) BeCoveredInMapSIKeys(m map[string]int) (bool, int) {
	for k, v := range m {
		if Str(k).Contains(s.V()) {
			return true, v
		}
	}
	return false, -1
}

// CoverAnyKeyInMapSI :                                                                            &
func (s Str) CoverAnyKeyInMapSI(m map[string]int) (bool, string, int) {
	for k, v := range m {
		if s.Contains(k) {
			return true, k, v
		}
	}
	return false, "", -1
}

// MkBrackets : e.g. "ABC"(BRound) => "(ABC)"                                                      &
func (s Str) MkBrackets(f BFlag) Str {
	bracketL := MatchAssign(f, BRound, BBox, BSquare, BCurly, BAngle, "(", "[", "[", "{", "<").(string)
	bracketR := MatchAssign(f, BRound, BBox, BSquare, BCurly, BAngle, ")", "]", "]", "}", ">").(string)
	return IF(s.HP(bracketL) && s.HS(bracketR), s, Str(bracketL+s.V()+bracketR)).(Str)
}

// RmBrackets : e.g. "(ABC)" => "ABC"                                                              &
func (s Str) RmBrackets(bfs ...BFlag) Str {
	if bfs[0] == BAll {
		bfs = []BFlag{BRound, BBox, BSquare, BCurly, BAngle}
	}
	for _, f := range bfs {
		bracketL := MatchAssign(f, BRound, BBox, BSquare, BCurly, BAngle, "(", "[", "[", "{", "<").(string)
		bracketR := MatchAssign(f, BRound, BBox, BSquare, BCurly, BAngle, ")", "]", "]", "}", ">").(string)
		if s.HP(bracketL) && s.HS(bracketR) {
			return s.S(1, ALL-1)
		}
	}
	return s

	// bracketL := MatchAssign(f, BRound, BBox, BSquare, BCurly, BAngle, "(", "[", "[", "{", "<").(string)
	// bracketR := MatchAssign(f, BRound, BBox, BSquare, BCurly, BAngle, ")", "]", "]", "}", ">").(string)
	// if s.HP(bracketL) && s.HS(bracketR) {
	// 	return s.S(1, ALL-1)
	// }
	// return s
}

// QuotePairCount :
func (s Str) QuotePairCount(f QFlag) int {
	quote := MatchAssign(f, QSingle, QDouble, '\'', '"').(rune)
	return sCnt(s.V(), string(quote)) / 2
}

// QuotesPos : index from 1                                                                        &
func (s Str) QuotesPos(f QFlag, index int) (str Str, left, right int) {
	quote := MatchAssign(f, QSingle, QDouble, '\'', '"').(rune)
	cnt, left, right := 0, -1, -1
	i := 0
	for _, c := range s {
		if left != -1 && right != -1 {
			break
		}
		if c == quote {
			cnt++
		}
		if (cnt+1)/2 == index {
			left = IF(cnt%2 == 1 && left == -1, i, left).(int)
			right = IF(cnt%2 == 0 && right == -1, i, right).(int)
		}
		i++
	}
	if left != -1 && right != -1 {
		return s.S(left, right+1), left, right
	}
	return "", -1, -1
}

// BracketsPos : level from 1, index from 1, if index > count, get the last one                    &
func (s Str) BracketsPos(f BFlag, level, index int) (str Str, left, right int) {
	bracketL := MatchAssign(f, BRound, BBox, BSquare, BCurly, BAngle, '(', '[', '[', '{', '<').(rune)
	bracketR := MatchAssign(f, BRound, BBox, BSquare, BCurly, BAngle, ')', ']', ']', '}', '>').(rune)
	curLevel, curIndex := 0, 0

	found := false
	i := 0
	for _, c := range s {
		curLevel = IF(c == bracketL, curLevel+1, curLevel).(int)
		curLevel = IF(c == bracketR, curLevel-1, curLevel).(int)
		if curLevel == level && c == bracketL {
			left = i
		}
		if curLevel == level-1 && c == bracketR {
			right = i
			curIndex++
			if curIndex == index {
				found = true
				break
			}
		}
		i++
	}

	if !found { // ** not break, come to here, not found **
		return "", -1, -1
	}

	return s.S(left, right+1), left, right
}

// BracketPairCount : only count top level                                                         &
func (s Str) BracketPairCount(f BFlag) (count int) {
	bracketL := MatchAssign(f, BRound, BBox, BSquare, BCurly, BAngle, '(', '[', '[', '{', '<').(rune)
	bracketR := MatchAssign(f, BRound, BBox, BSquare, BCurly, BAngle, ')', ']', ']', '}', '>').(rune)
	level, inflag := 0, false
	for _, c := range s {
		if c == bracketL {
			level++
		}
		if c == bracketR {
			level--
			if level == 0 {
				inflag = false
			}
		}
		if level == 1 {
			if !inflag {
				count++
				inflag = true
			}
		}
	}
	return count
}

// BracketDepth :                                                                                  &
func (s Str) BracketDepth(f BFlag, pos int) int {
	bracketL := MatchAssign(f, BRound, BBox, BSquare, BCurly, BAngle, '(', '[', '[', '{', '<').(rune)
	bracketR := MatchAssign(f, BRound, BBox, BSquare, BCurly, BAngle, ')', ']', ']', '}', '>').(rune)
	level, found := 0, false
	i := 0
	for _, c := range s {
		if c == bracketL {
			level++
		}
		if i == pos {
			found = true
			break
		}
		if c == bracketR {
			level--
		}
		i++
	}
	return IF(found, level, -1).(int)
}

// MkQuotes : e.g. "ABC"(QSingle) => "'ABC'"                                                       &
func (s Str) MkQuotes(f QFlag) Str {
	quote := MatchAssign(f, QSingle, QDouble, "'", "\"").(string)
	return IF(s.HP(quote) && s.HS(quote), s, Str(quote+s.V()+quote)).(Str)
}

// RmQuotes : Remove single or double Quotes from a string. If no quotations, do nothing           &
func (s Str) RmQuotes(qfs ...QFlag) Str {
	if qfs[0] == QAll {
		qfs = []QFlag{QSingle, QDouble}
	}
	for _, f := range qfs {
		quote := MatchAssign(f, QSingle, QDouble, "'", "\"").(string)
		if s.HP(quote) && s.HS(quote) {
			return s.S(1, ALL-1)
		}
	}
	return s

	// quote := MatchAssign(f, QSingle, QDouble, "'", "\"").(string)
	// if s.HP(quote) && s.HS(quote) {
	// 	return s.S(1, ALL-1)
	// }
	// return s
}

// MkPrefix : e.g. "ABC"("abc") => "abcABC"                                                        &
func (s Str) MkPrefix(prefix string) Str {
	return IF(!s.HP(prefix), Str(prefix+s.V()), s).(Str)
}

// RmPrefix : e.g. "abcABC"("abc") => "ABC"                                                        &
func (s Str) RmPrefix(prefix string) Str {
	if s.HP(prefix) {
		return s.S(Str(prefix).L(), ALL)
	}
	return s
}

// MkSuffix : e.g. "ABC"("abc") => "ABCabc"                                                        &
func (s Str) MkSuffix(suffix string) Str {
	return IF(!s.HS(suffix), Str(s.V()+suffix), s).(Str)
}

// RmSuffix : e.g. "ABCabc"("abc") => "ABC"                                                        &
func (s Str) RmSuffix(suffix string) Str {
	if s.HS(suffix) {
		return s.S(0, ALL-Str(suffix).L())
	}
	return s
}

// RmTailFromLast : e.g. "AB.CD.EF"(".") => "AB.CD"                                                &
func (s Str) RmTailFromLast(mark string) Str {
	if i := s.LIdx(mark); i >= 0 {
		return s.S(0, i)
	}
	return s
}

// RmHeadToLast : e.g. "AB.CD.EF"(".") => "EF"                                                     &
func (s Str) RmHeadToLast(mark string) Str {
	mL := Str(mark).L()
	if i := s.LIdx(mark); i >= 0 {
		if i < s.L()-mL {
			return s.S(i+mL, ALL)
		}
		return ""
	}
	return s
}

// RmBlankBefore :                                                                                 &
func (s Str) RmBlankBefore(strs ...string) Str {
	whole := s.V()
	for _, str := range strs {
		str0, str1 := " "+str, "\t"+str
	NEXT:
		sWhole := Str(whole)
		if p := sWhole.Idx(str0); p >= 0 {
			whole = sWhole.SegRep(p, p+1, "").V()
			goto NEXT
		}
		if p := sWhole.Idx(str1); p >= 0 {
			whole = sWhole.SegRep(p, p+1, "").V()
			goto NEXT
		}
	}
	return Str(whole)
}

// RmBlankAfter :                                                                                  &
func (s Str) RmBlankAfter(strs ...string) Str {
	whole := s.V()
	for _, str := range strs {
		str0, str1 := str+" ", str+"\t"
		len0, len1 := Str(str0).L(), Str(str1).L()
	NEXT:
		sWhole := Str(whole)
		if p := sWhole.Idx(str0); p >= 0 {
			whole = sWhole.SegRep(p+len0-1, p+len0, "").V()
			goto NEXT
		}
		if p := sWhole.Idx(str1); p >= 0 {
			whole = sWhole.SegRep(p+len1-1, p+len1, "").V()
			goto NEXT
		}
	}
	return Str(whole)
}

// RmBlankNear :                                                                                   &
func (s Str) RmBlankNear(strs ...string) Str {
	s0 := s.RmBlankBefore(strs...)
	return Str(s0).RmBlankAfter(strs...)
}

// RmBlankNBefore : n start from 1                                                                 &
func (s Str) RmBlankNBefore(n int, str string) Str {
	// whole, left, right, strs := s.V(), "", "", []string{}
	// for i := 0; i < n; i++ {
	// 	if p := sI(whole, str); p >= 0 {
	// 		left, right = whole[:p+1], whole[p+1:]
	// 		left, whole = Str(left).RmBlankBefore(str), right
	// 		strs = append(strs, left)
	// 		if i == n-1 {
	// 			strs = append(strs, right)
	// 		}
	// 	} else {
	// 		if right != "" {
	// 			strs = append(strs, right)
	// 		}
	// 		break
	// 	}
	// }
	// return sJ(strs, "")

	segs, strs := sSpl(s.V(), str), []string{}
	for i, seg := range segs {
		if i < n {
			seg = Str(seg).TR(BLANK).V()
		}
		strs = append(strs, seg)
	}
	return Str(sJ(strs, str))
}

// RmBlankNAfter : n start from 1                                                                  $
func (s Str) RmBlankNAfter(n int, str string) Str {
	strs := []string{}
	for i, seg := range sSpl(s.V(), str) {
		if i >= 1 && i <= n {
			seg = Str(seg).TL(BLANK).V()
		}
		strs = append(strs, seg)
	}
	return Str(sJ(strs, str))
}

// RmBlankNNear : n start from 1                                                                   $
func (s Str) RmBlankNNear(n int, str string) Str {
	// s0 := s.RmBlankNBefore(n, str)
	// return Str(s0).RmBlankNAfter(n, str)

	segs, strs := sSpl(s.V(), str), []string{}
	for i, seg := range segs {
		if i == 0 && i != n {
			seg = Str(seg).TR(BLANK).V()
		} else if i == n {
			seg = Str(seg).TL(BLANK).V()
		} else if i >= 1 && i < n {
			seg = Str(seg).T(BLANK).V()
		}
		strs = append(strs, seg)
	}
	return Str(sJ(strs, str))
}

// TrimInternal :                                                                                  $
func (s Str) TrimInternal(cutset rune, nkeep int) (r Str) {
	pos, lens, strs := []int{}, []int{}, []string{}
	L := s.L()
	p := 0
	for _, c := range s {
		if p < L-1 {
			if c != cutset && s.C(p+1) == cutset {
				pos = append(pos, p+1)
			}
		}
		p++
	}
NEXT:
	for _, p := range pos {
		ncs := 0
		for _, c := range s.S(p, ALL, L) {
			if c != cutset {
				// lens = append(lens, MinOf(ncs, nkeep))
				min, _ := Min(I32s{ncs, nkeep}, "")
				lens = append(lens, min.(int))
				continue NEXT
			}
			ncs++
		}
	}
	for _, str := range sFF(s.V(), func(c rune) bool { return c == cutset }) {
		strs = append(strs, str)
	}
	// ***
	cntL, cntR := 0, 0
	p = 0
	for _, c := range s {
		if c != cutset {
			cntL = p
			break
		}
		p++
	}
	for p := L - 1; p >= 0; p-- {
		if s.C(p) != cutset {
			cntR = L - p - 1
			break
		}
	}

	r += Str(cutset).Repeat(cntL, "")
	for i, str := range strs {
		r += Str(str)
		if i < len(strs)-1 {
			r += Str(cutset).Repeat(lens[i], "")
		}
	}
	r += Str(cutset).Repeat(cntR, "")
	return r
}

// TrimInternalEachLine :                                                                          ?
func (s Str) TrimInternalEachLine(cutset rune, nkeep int) (r Str) {
	strs := []string{}
	lns := sFF(s.V(), func(c rune) bool { return c == '\n' })
	for iln, ln := range lns {
		strs = append(strs, Str(ln).TrimInternal(cutset, nkeep).V())
		if iln < len(lns)-1 {
			strs = append(strs, "\n")
		}
	}
	str := IF(s.C(LAST) == '\n', sJ(strs, "")+"\n", sJ(strs, "")).(string)
	return Str(str)
}

// TrimAllInternal :                                                                               &
func (s Str) TrimAllInternal(cutset string) (r Str) {
	cuts := []rune(cutset)
	for _, c := range cuts {
		r = s.TrimInternal(c, 0)
		s = Str(r)
	}
	return
}

// TrimAllLMR :                                                                                    &
func (s Str) TrimAllLMR(cutset string) Str {
	return s.T(cutset).TrimAllInternal(cutset)
}

// KeyValueMap :                                                                                   &
func (s Str) KeyValueMap(delimiter, assign, terminator rune) (r map[string]string) {
	r = make(map[string]string)
	sAssign := string(assign)
	str := s.RmBlankNear(sAssign)
	Sstr := Str(str)
	if pt := Sstr.Idx(string(terminator)); pt > 0 {
		str = Sstr.S(0, pt)
	}
	for _, kv := range sFF(str.V(), func(c rune) bool { return c == delimiter }) {
		if Str(kv).Contains(sAssign) {
			kvpair := sSpl(kv, sAssign)
			kvpair[0] = Str(kvpair[0]).T(BLANK).V()
			r[kvpair[0]] = Str(kvpair[1]).T(BLANK).RmQuotes(QDouble).V()
		}
	}
	return
}

// KeyValuePair : (if assign mark cannot be found, k is empty, v is original string)               &
func (s Str) KeyValuePair(assign, termsK, termsV string, rmQuotes, trimBlank bool) (k, v Str) {
	str, sAssign := Str(s.RmBlankNNear(1, assign)), Str(assign)
	Lassign := sAssign.L()
	if p := str.Idx(assign); p >= 0 {
		Kstr, Vstr := str.S(0, p), str.S(p+Lassign, ALL)
		k, v = Kstr, Vstr

		lk, lv := I32s{}, I32s{} //[]int{}, []int{}
		for _, tk := range []rune(termsK) {
			lk = append(lk, Kstr.LIdx(string(tk)))
		}
		for _, tv := range []rune(termsV) {
			lv = append(lv, Vstr.Idx(string(tv)))
		}

		if pk, i := Min(lk, "Non-Neg"); i >= 0 && pk.(int) >= 0 {
			k = str.S(pk.(int)+1, p)
		}
		if pv, i := Min(lv, "Non-Neg"); i >= 0 && pv.(int) >= 0 {
			v = str.S(p+Lassign, p+Lassign+pv.(int))
		}
	} else {
		return "", s
	}
	if trimBlank {
		k, v = k.T(BLANK), v.T(BLANK)
	}
	if rmQuotes {
		k, v = k.RmQuotes(QDouble), v.RmQuotes(QDouble)
	}
	return
}

// Indices :                                                                                       &
func (s Str) Indices(aim string) (posList []int) {
	str := s.V()
	la := Str(aim).L()
AGAIN:
	if p := Str(str).Idx(aim); p >= 0 {
		pos := p
		if len(posList) > 0 {
			pos = p + posList[len(posList)-1] + la
		}
		posList = append(posList, pos)
		str = Str(str).S(p+la, ALL).V()
		goto AGAIN
	}
	return
}

// IdxAnyInRange :
func (s Str) IdxAnyInRange(prefix string, strs []string, suffix string) int {
	ps := []int{}
	for _, str := range strs {
		aim := prefix + str + suffix
		if p := s.Idx(aim); p >= 0 {
			ps = append(ps, p)
		}
	}
	if p, _ := Min(I32s(ps), ""); p != nil {
		return p.(int)
	}
	return -1
}

// Indices2StrsIgnore :
func (s Str) Indices2StrsIgnore(first, second, ignore string) (starts, ends []int) {
AGAIN:
	if ok, start, end := s.Search2StrsIgnore(first, second, ignore); ok {
		starts, ends = append(starts, start), append(ends, end)
		s = s.SegRep(start, start+1, "*")
		goto AGAIN
	}
	return
}

// Search2StrsIgnore :
func (s Str) Search2StrsIgnore(first, second, ignore string) (bool, int, int) {
	fL := Str(first).L()
	for _, fPos := range s.Indices(first) {
		st := s.S(fPos+fL, ALL)
		if st.TL(ignore).HP(second) {
			return true, fPos, fPos + fL + st.Idx(second)
		}
	}
	return false, -1, -1
}

// SearchStrsIgnore : last param is ignored runes string
func (s Str) SearchStrsIgnore(aims ...string) (ok bool, start, end int) {
	PC(len(aims) < 3, fEf("At least 3 params, the last is ignored runes string"))

	start, end = -1, -1
	nCheck, prevScd := len(aims)-2, -1
	ignored := aims[nCheck+1]

	for _, p := range s.Indices(aims[0]) {
		ok = false
		curS, iCheck := s.S(p, ALL), 0

		for j := 0; j < nCheck; j++ {
			first, second := aims[j], aims[j+1]
			if ok, fst, scd := curS.Search2StrsIgnore(first, second, ignored); ok {
				if j == 0 || (j != 0 && fst == prevScd) {
					iCheck++
					prevScd = scd
					start = IF(j == 0, p+fst, start).(int)
					end = IF(j == nCheck-1, p+scd, end).(int)
				}
			}
		}
		if iCheck == nCheck {
			ok = true
			break
		}
	}

	start, end = IF(ok, start, -1).(int), IF(ok, end, -1).(int)
	return
}

// SearchAny2StrsIgnore :
func (s Str) SearchAny2StrsIgnore(fstArr, scdArr []string, ignore string) (bool, int, int) {
	OK, starts, ends := false, I32s{}, I32s{}
	for _, fst := range fstArr {
		for _, scd := range scdArr {
			if ok, start, end := s.Search2StrsIgnore(fst, scd, ignore); ok {
				OK, starts, ends = true, append(starts, start), append(ends, end)
			}
		}
	}
	smin, _ := Min(starts, "")
	emin, _ := Min(ends, "")
	if OK {
		return OK, smin.(int), emin.(int)
	}
	return OK, -1, -1
}

// // IsXMLSegSimple :
// func (s Str) IsXMLSegSimple() bool {
// 	c := s.BracketPairCount(BAngle)
// 	if c == 0 {
// 		return false
// 	}
// 	tagsstr, _, _ := s.BracketsPos(BAngle, 1, 1)
// 	tagestr, _, _ := s.BracketsPos(BAngle, 1, c)
// 	tagsStr, tageStr := Str(tagsstr), Str(tagestr)
// 	tage := tageStr.S(2, LAST-1)
// 	tags := tagsStr.S(1, 1+tage.L())
// 	return tags == tage &&
// 		(tagsStr.C(tags.L()+1) == ' ' || tagsStr.C(tags.L()+1) == '>') &&
// 		tagsStr.C(0) == '<' &&
// 		tageStr.S(0, 2).V() == "</"
// }

// IsUUID :
func (s Str) IsUUID() bool {
	_, e := uuid.Parse(s.V())
	return e == nil
}

// FieldsSeqCtn :
func (s Str) FieldsSeqCtn(str, sep string) bool {
	sArr0, sArr1 := sSpl(s.V(), sep), sSpl(str, sep)
	gArr1 := IArr2GArr(Strs(sArr1))
	return IArrSeqCtns(Strs(sArr0), gArr1...)
}

// SplitEx :
func (s Str) SplitEx(strSep, itemSep, r1AEType, r2AEType string) (r1, r2 interface{}) {

	end := false
	ss1, ss2 := []string{}, []string{}
	for _, item := range sSpl(s.V(), strSep) {
		coms := sSpl(item, itemSep)
		switch len(coms) {
		case 2:
			ss1, ss2 = append(ss1, coms[0]), append(ss2, coms[1])
			PC(end, fEf("error format string to SplitEx"))
		case 1:
			ss1 = append(ss1, coms[0])
			end = true
		}
	}

	switch r1AEType {
	case "string":
		r1 = ss1
	case "int":
		{
			a1 := make([]int, len(ss1))
			for i, s1 := range ss1 {
				a1[i] = Str(s1).ToInt()
			}
			r1 = a1
		}
	default:
		panic("not implemented " + r1AEType + " in StrsSplitEx r1AEType")
	}

	switch r2AEType {
	case "string":
		r2 = ss2
	case "int":
		{
			a2 := make([]int, len(ss2))
			for i, s2 := range ss2 {
				a2[i] = Str(s2).ToInt()
			}
			r2 = a2
		}
	default:
		panic("not implemented " + r2AEType + " in StrsSplitEx r2AEType")
	}

	return
}
