package gopherobject

type GopherObjectType int

const (
	GopherString GopherObjectType = iota
	GopherHashmap
	GopherList
	GopherSet
)

var TypeStringMap = map[GopherObjectType]string{
	GopherString:  "GopherString",
	GopherHashmap: "GopherHashmap",
	GopherList:    "GopherList",
	GopherSet:     "GopherSet",
}

func DtypeToString(t GopherObjectType) string {
	return TypeStringMap[t]
}
