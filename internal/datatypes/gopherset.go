package datatypes

type GopherSet struct {
	Data map[string]bool
}

func NewGopherSet() *GopherSet {
	return &GopherSet{
		make(map[string]bool),
	}
}

func (gs *GopherSet) Sadd(args []string) {
	for _, item := range args {
		gs.Data[item] = true
	}
}

func (gs *GopherSet) Srem(args []string) {
	for _, item := range args {
		delete(gs.Data, item)
	}
}

func (gs *GopherSet) Sismember(s string) bool {
	val, ok := gs.Data[s]
	return val && ok
}

func (gs *GopherSet) Scard() int {
	return len(gs.Data)
}
