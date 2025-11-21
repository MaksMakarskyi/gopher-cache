package cmds

type GopherCommand struct {
	Name       string
	Args       []string
	ResponseCh chan<- string
}

func NewGopherCommand(n string, a []string, r chan<- string) GopherCommand {
	return GopherCommand{
		Name:       n,
		Args:       a,
		ResponseCh: r,
	}
}
