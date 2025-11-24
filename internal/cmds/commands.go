package cmds

type GopherCommand struct {
	Name       string
	Args       []any
	ResponseCh chan<- string
}

func NewGopherCommand(n string, a []any, r chan<- string) GopherCommand {
	return GopherCommand{
		Name:       n,
		Args:       a,
		ResponseCh: r,
	}
}
