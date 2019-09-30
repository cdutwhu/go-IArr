package wrappers

import (
	"fmt"
	"strconv"
	"strings"

	u "github.com/cdutwhu/go-util"
)

var (
	sJ       = strings.Join
	sSpl     = strings.Split
	sCtns    = strings.Contains
	sCnt     = strings.Count
	sFF      = strings.FieldsFunc
	sRplc    = strings.Replace
	sRplcAll = strings.ReplaceAll
	sc2Int   = strconv.ParseInt
	sc2Uint  = strconv.ParseUint
	sc2Float = strconv.ParseFloat
	fPln     = fmt.Println
	fSf      = fmt.Sprintf
	fEf      = fmt.Errorf
	pc       = u.PanicOnCondition
	IF       = u.IF
	must     = u.Must
	mAssign  = u.MatchAssign
)

const (
	BLANK = " \t\n\r"
)
