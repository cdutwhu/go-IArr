package wrappers

type (
	QFlag int // QFlag : Flag for Quotes, single or double
	BFlag int // BFlag : Flag for Brackets
	EnCo  int // EnCo : string encoded type
)

const (
	ALL     = MaxInt
	TORANGE = 2048
	LAST    = ALL

	QSingle QFlag = 1 // QSingle : single quotes   ''
	QDouble QFlag = 2 // QDouble : double quotes   ""

	BRound  BFlag = 1 // BRound : round brackets   ()
	BBox    BFlag = 2 // BBox : box brackets       []
	BSquare BFlag = 2 // BSquare : square brackets []
	BCurly  BFlag = 3 // BCurly : curly brackets   {}
	BAngle  BFlag = 4 // BAngle : angle brackets   <>

	ASCII EnCo = 1
	UTF8  EnCo = 2
)

var (
	EnCoDesc      = map[EnCo]string{ASCII: "ASCII", UTF8: "UTF-8"}
	EnCoType EnCo = UTF8
)
