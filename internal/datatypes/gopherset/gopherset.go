package gopherset

type GopherSet struct {
	Data map[string]bool
}

func NewGopherSet() *GopherSet {
	return &GopherSet{
		make(map[string]bool),
	}
}

func (gs *GopherSet) Sadd(args []string) int {
	count := 0
	for _, item := range args {
		if val, ok := gs.Data[item]; !ok || !val {
			gs.Data[item] = true
			count += 1
		}
	}

	return count
}

func (gs *GopherSet) Srem(args []string) int {
	count := 0
	for _, item := range args {
		if val, ok := gs.Data[item]; ok && val {
			delete(gs.Data, item)
			count += 1
		}
	}

	return count
}

func (gs *GopherSet) Sismember(s string) int {
	// Return 1 if memmber exist, otherwise return 0

	val, ok := gs.Data[s]
	if val && ok {
		return 1
	}

	return 0
}

func (gs *GopherSet) Scard() int {
	return len(gs.Data)
}
