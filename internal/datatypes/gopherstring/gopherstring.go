package gopherstring

type GopherString struct {
	Data string
}

func NewGopherString(s string) *GopherString {
	return &GopherString{Data: s}
}

func (gs *GopherString) Get() string {
	return gs.Data
}
