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
	sc2Float = strconv.ParseFloat

	fPln = fmt.Println
	fSf  = fmt.Sprintf
	fEf  = fmt.Errorf

	PC         = u.PanicOnCondition
	IF         = u.IF
	Must       = u.Must
	CaseAssign = u.CaseAssign
)

const (
	BLANK = " \t\n\r"
)
