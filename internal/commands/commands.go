package command

type GopherCommand struct {
	Name       string
	Args       []string
	ResponseCh chan string
}
