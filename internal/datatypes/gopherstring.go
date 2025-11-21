package datatypes

type GopherString struct {
	Entry string
}

func NewGopherString(s string) *GopherString {
	return &GopherString{Entry: s}
}
