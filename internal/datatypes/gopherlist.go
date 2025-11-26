package datatypes

type ListNode struct {
	Val  string
	Next *ListNode
	Prev *ListNode
}

type GopherList struct {
	Head *ListNode
	Tail *ListNode
	Len  int
}

func NewGopherList() *GopherList {
	return &GopherList{
		nil,
		nil,
		0,
	}
}

func (gl *GopherList) Lpush(args []string) {
	for _, item := range args {
		newNode := ListNode{
			item,
			gl.Head,
			nil,
		}

		if gl.Head == nil && gl.Tail == nil {
			gl.Head = &newNode
			gl.Tail = &newNode
		} else {
			gl.Head.Prev = &newNode
			gl.Head = &newNode
		}

		gl.Len += 1
	}
}

func (gl *GopherList) Rpush(args []string) {
	for _, item := range args {
		newNode := ListNode{
			item,
			nil,
			gl.Tail,
		}

		if gl.Head == nil && gl.Tail == nil {
			gl.Head = &newNode
			gl.Tail = &newNode
		} else {
			gl.Tail.Next = &newNode
			gl.Tail = &newNode
		}

		gl.Len += 1
	}
}

func (gl *GopherList) Lpop(count int) {
	for range count {
		if gl.Head == nil && gl.Tail == nil {
			break
		} else if gl.Head == gl.Tail {
			gl.Head = nil
			gl.Tail = nil
			gl.Len -= 1
			break
		}

		gl.Head = gl.Head.Next
		gl.Head.Prev = nil
		gl.Len -= 1
	}
}

func (gl *GopherList) Rpop(count int) {
	for range count {
		if gl.Head == nil && gl.Tail == nil {
			break
		} else if gl.Head == gl.Tail {
			gl.Head = nil
			gl.Tail = nil
			gl.Len -= 1
			break
		}

		gl.Tail = gl.Tail.Prev
		gl.Tail.Next = nil
		gl.Len -= 1
	}
}

func (gl *GopherList) Llen() int {
	return gl.Len
}
