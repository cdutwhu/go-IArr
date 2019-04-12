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
)

const (
	ASCII EnCo = iota
	UTF8
)

const (
	QSingle QFlag = 1 << iota // QSingle : single quotes   ''
	QDouble                   // QDouble : double quotes   ""
	QAll
)

const (
	BRound BFlag = 1 << iota // BRound : round brackets   ()
	BBox                     // BBox : box brackets       []
	BCurly                   // BCurly : curly brackets   {}
	BAngle                   // BAngle : angle brackets   <>
	BAll
	BSquare BFlag = BBox // BSquare : square brackets []
)

var (
	EnCoDesc      = map[EnCo]string{ASCII: "ASCII", UTF8: "UTF-8"}
	EnCoType EnCo = UTF8
)
