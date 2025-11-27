package gopherlist

import (
	"testing"
)

func CreateInitialList(args []string) *GopherList {
	l := NewGopherList()

	if len(args) > 0 {
		var curNode *ListNode
		for i := range len(args) {
			nextNode := ListNode{args[i], nil, nil}

			if i == 0 {
				l.Head = &nextNode
				curNode = &nextNode
			} else {
				curNode.Next = &nextNode
				nextNode.Prev = curNode
				curNode = curNode.Next
			}

			l.Len += 1
		}
		l.Tail = curNode
	}

	return l
}

func CheckList(list *GopherList, target []string) bool {
	if len(target) == 0 && list.Head == nil && list.Tail == nil && list.Len == 0 {
		return true
	}

	curNode := list.Head
	for i := range len(target) {
		if curNode == nil || curNode.Val != target[i] {
			return false
		}
		curNode = curNode.Next
	}

	curNode = list.Tail
	for i := range len(target) {
		if curNode == nil || curNode.Val != target[len(target)-i-1] {
			return false
		}
		curNode = curNode.Prev
	}

	return list.Len == len(target)
}

// *PUSH
type TestCasePush struct {
	Name        string
	InitialList []string
	Args        []string
	FinalList   []string
}

// LPUSH
var LpushTests = []TestCasePush{
	{
		"empty_array_1",
		make([]string, 0),
		[]string{"foo", "bar", "buz"},
		[]string{"buz", "bar", "foo"},
	},
	{
		"empty_array_2",
		make([]string, 0),
		[]string{"foo", "bar"},
		[]string{"bar", "foo"},
	},
	{
		"empty_array_3",
		make([]string, 0),
		[]string{"bar"},
		[]string{"bar"},
	},
	{
		"list_of_one_1",
		[]string{"foo"},
		[]string{"bar", "fiz"},
		[]string{"fiz", "bar", "foo"},
	},
	{
		"list_of_one_2",
		[]string{"foo"},
		[]string{"bar"},
		[]string{"bar", "foo"},
	},
	{
		"list_of_two",
		[]string{"foo", "bar"},
		[]string{"buz", "fiz"},
		[]string{"fiz", "buz", "foo", "bar"},
	},
}

func TestLpush(t *testing.T) {
	for i, test := range LpushTests {
		t.Run(test.Name, func(t *testing.T) {
			list := CreateInitialList(test.InitialList)
			list.Lpush(test.Args)

			if !CheckList(list, test.FinalList) {
				t.Errorf("#%d: Invalid final list", i)
			}
		})
	}
}

// RPUSH
var RpushTests = []TestCasePush{
	{
		"empty_array_1",
		make([]string, 0),
		[]string{"foo", "bar", "buz"},
		[]string{"foo", "bar", "buz"},
	},
	{
		"empty_array_2",
		make([]string, 0),
		[]string{"foo", "bar"},
		[]string{"foo", "bar"},
	},
	{
		"empty_array_3",
		make([]string, 0),
		[]string{"bar"},
		[]string{"bar"},
	},
	{
		"list_of_one_1",
		[]string{"foo"},
		[]string{"bar", "fiz"},
		[]string{"foo", "bar", "fiz"},
	},
	{
		"list_of_one_2",
		[]string{"foo"},
		[]string{"bar"},
		[]string{"foo", "bar"},
	},
	{
		"list_of_two",
		[]string{"foo", "bar"},
		[]string{"buz", "fiz"},
		[]string{"foo", "bar", "buz", "fiz"},
	},
}

func TestRpush(t *testing.T) {
	for i, test := range RpushTests {
		t.Run(test.Name, func(t *testing.T) {
			list := CreateInitialList(test.InitialList)
			list.Rpush(test.Args)

			if !CheckList(list, test.FinalList) {
				t.Errorf("#%d: Invalid final list", i)
			}
		})
	}
}

// *POP
type TestCasePop struct {
	Name        string
	InitialList []string
	Count       int
	FinalList   []string
}

// LPOP
var LpopTests = []TestCasePop{
	{
		"empty_list_1",
		make([]string, 0),
		1,
		make([]string, 0),
	},
	{
		"empty_list_2",
		make([]string, 0),
		15,
		make([]string, 0),
	},
	{
		"list_of_one_1",
		[]string{"foo"},
		1,
		make([]string, 0),
	},
	{
		"list_of_one_2",
		[]string{"foo"},
		7,
		make([]string, 0),
	},
	{
		"list_of_two_1",
		[]string{"foo", "bar"},
		1,
		[]string{"bar"},
	},
	{
		"list_of_two_2",
		[]string{"foo", "bar"},
		2,
		make([]string, 0),
	},
	{
		"random_list_1",
		[]string{"foo", "bar", "fizz", "bazz", "golang", "python", "yavascript"},
		4,
		[]string{"golang", "python", "yavascript"},
	},
	{
		"random_list_2",
		[]string{"foo", "bar", "fizz", "bazz", "golang", "python", "yavascript"},
		15,
		make([]string, 0),
	},
	{
		"zero_count",
		[]string{"foo", "bar", "fizz", "bazz", "golang", "python", "yavascript"},
		0,
		[]string{"foo", "bar", "fizz", "bazz", "golang", "python", "yavascript"},
	},
}

func TestLpop(t *testing.T) {
	for i, test := range LpopTests {
		t.Run(test.Name, func(t *testing.T) {
			list := CreateInitialList(test.InitialList)
			list.Lpop(test.Count)

			if !CheckList(list, test.FinalList) {
				t.Errorf("#%d: Invalid final list", i)
			}
		})
	}
}

// RPOP
var RpopTests = []TestCasePop{
	{
		"empty_list_1",
		make([]string, 0),
		1,
		make([]string, 0),
	},
	{
		"empty_list_2",
		make([]string, 0),
		15,
		make([]string, 0),
	},
	{
		"list_of_one_1",
		[]string{"foo"},
		1,
		make([]string, 0),
	},
	{
		"list_of_one_2",
		[]string{"foo"},
		7,
		make([]string, 0),
	},
	{
		"list_of_two_1",
		[]string{"foo", "bar"},
		1,
		[]string{"foo"},
	},
	{
		"list_of_two_2",
		[]string{"foo", "bar"},
		2,
		make([]string, 0),
	},
	{
		"random_list_1",
		[]string{"foo", "bar", "fizz", "bazz", "golang", "python", "yavascript"},
		4,
		[]string{"foo", "bar", "fizz"},
	},
	{
		"random_list_2",
		[]string{"foo", "bar", "fizz", "bazz", "golang", "python", "yavascript"},
		15,
		make([]string, 0),
	},
	{
		"zero_count",
		[]string{"foo", "bar", "fizz", "bazz", "golang", "python", "yavascript"},
		0,
		[]string{"foo", "bar", "fizz", "bazz", "golang", "python", "yavascript"},
	},
}

func TestRpop(t *testing.T) {
	for i, test := range RpopTests {
		t.Run(test.Name, func(t *testing.T) {
			list := CreateInitialList(test.InitialList)
			list.Rpop(test.Count)

			if !CheckList(list, test.FinalList) {
				t.Errorf("#%d: Invalid final list", i)
			}
		})
	}
}

// Llen
type TestCaseLlen struct {
	Name        string
	InitialList []string
	TargetLen   int
}

var LlenTests = []TestCaseLlen{
	{
		"empty_list",
		make([]string, 0),
		0,
	},
	{
		"list_of_one",
		[]string{"fiz"},
		1,
	},
	{
		"list_of_two",
		[]string{"fiz", "baz"},
		2,
	},
	{
		"random_list_1",
		[]string{"foo", "bar", "fizz", "bazz", "golang", "python", "yavascript"},
		7,
	},
	{
		"random_list_2",
		[]string{"foo", "bar", "fizz", "bazz", "golang", "python", "yavascript", "foo", "bar", "fizz", "bazz", "golang", "python", "yavascript"},
		14,
	},
}

func TestLlen(t *testing.T) {
	for i, test := range LlenTests {
		t.Run(test.Name, func(t *testing.T) {
			list := CreateInitialList(test.InitialList)
			length := list.Llen()

			if length != test.TargetLen {
				t.Errorf("#%d: Length mismatch, expected: %d, got: %d", i, test.TargetLen, length)
			}
		})
	}
}
