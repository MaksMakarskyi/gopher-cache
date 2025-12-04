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

func CompareSlices(l1 []string, l2 []string) bool {
	if len(l1) != len(l2) {
		return false
	}

	for i := range l1 {
		if l1[i] != l2[i] {
			return false
		}
	}

	return true
}

// *PUSH
type TestCasePush struct {
	Name          string
	InitialList   []string
	Args          []string
	FinalList     []string
	ExpectedCount int
}

// LPUSH
var LpushTests = []TestCasePush{
	{
		"empty_array_1",
		make([]string, 0),
		[]string{"foo", "bar", "buz"},
		[]string{"buz", "bar", "foo"},
		3,
	},
	{
		"empty_array_2",
		make([]string, 0),
		[]string{"foo", "bar"},
		[]string{"bar", "foo"},
		2,
	},
	{
		"empty_array_3",
		make([]string, 0),
		[]string{"bar"},
		[]string{"bar"},
		1,
	},
	{
		"list_of_one_1",
		[]string{"foo"},
		[]string{"bar", "fiz"},
		[]string{"fiz", "bar", "foo"},
		3,
	},
	{
		"list_of_one_2",
		[]string{"foo"},
		[]string{"bar"},
		[]string{"bar", "foo"},
		2,
	},
	{
		"list_of_two",
		[]string{"foo", "bar"},
		[]string{"buz", "fiz"},
		[]string{"fiz", "buz", "foo", "bar"},
		4,
	},
}

func TestLpush(t *testing.T) {
	for i, test := range LpushTests {
		t.Run(test.Name, func(t *testing.T) {
			list := CreateInitialList(test.InitialList)
			count := list.Lpush(test.Args)

			if !CheckList(list, test.FinalList) {
				t.Errorf("#%d: Invalid final list", i)
			}

			if count != test.ExpectedCount {
				t.Errorf("#%d: Count Expected: %d, Got: %d", i, test.ExpectedCount, count)
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
		3,
	},
	{
		"empty_array_2",
		make([]string, 0),
		[]string{"foo", "bar"},
		[]string{"foo", "bar"},
		2,
	},
	{
		"empty_array_3",
		make([]string, 0),
		[]string{"bar"},
		[]string{"bar"},
		1,
	},
	{
		"list_of_one_1",
		[]string{"foo"},
		[]string{"bar", "fiz"},
		[]string{"foo", "bar", "fiz"},
		3,
	},
	{
		"list_of_one_2",
		[]string{"foo"},
		[]string{"bar"},
		[]string{"foo", "bar"},
		2,
	},
	{
		"list_of_two",
		[]string{"foo", "bar"},
		[]string{"buz", "fiz"},
		[]string{"foo", "bar", "buz", "fiz"},
		4,
	},
}

func TestRpush(t *testing.T) {
	for i, test := range RpushTests {
		t.Run(test.Name, func(t *testing.T) {
			list := CreateInitialList(test.InitialList)
			count := list.Rpush(test.Args)

			if !CheckList(list, test.FinalList) {
				t.Errorf("#%d: Invalid final list", i)
			}

			if count != test.ExpectedCount {
				t.Errorf("#%d: Count Expected: %d, Got: %d", i, test.ExpectedCount, count)
			}
		})
	}
}

// *POP
type TestCasePop struct {
	Name             string
	InitialList      []string
	Count            int
	FinalList        []string
	ExpectedResponse []string
}

// LPOP
var LpopTests = []TestCasePop{
	{
		"empty_list_1",
		make([]string, 0),
		1,
		make([]string, 0),
		make([]string, 0),
	},
	{
		"empty_list_2",
		make([]string, 0),
		15,
		make([]string, 0),
		make([]string, 0),
	},
	{
		"list_of_one_1",
		[]string{"foo"},
		1,
		make([]string, 0),
		[]string{"foo"},
	},
	{
		"list_of_one_2",
		[]string{"foo"},
		7,
		make([]string, 0),
		[]string{"foo"},
	},
	{
		"list_of_two_1",
		[]string{"foo", "bar"},
		1,
		[]string{"bar"},
		[]string{"foo"},
	},
	{
		"list_of_two_2",
		[]string{"foo", "bar"},
		2,
		make([]string, 0),
		[]string{"foo", "bar"},
	},
	{
		"random_list_1",
		[]string{"foo", "bar", "fizz", "bazz", "golang", "python", "yavascript"},
		4,
		[]string{"golang", "python", "yavascript"},
		[]string{"foo", "bar", "fizz", "bazz"},
	},
	{
		"random_list_2",
		[]string{"foo", "bar", "fizz", "bazz", "golang", "python", "yavascript"},
		15,
		make([]string, 0),
		[]string{"foo", "bar", "fizz", "bazz", "golang", "python", "yavascript"},
	},
	{
		"zero_count",
		[]string{"foo", "bar", "fizz", "bazz", "golang", "python", "yavascript"},
		0,
		[]string{"foo", "bar", "fizz", "bazz", "golang", "python", "yavascript"},
		make([]string, 0),
	},
}

func TestLpop(t *testing.T) {
	for i, test := range LpopTests {
		t.Run(test.Name, func(t *testing.T) {
			list := CreateInitialList(test.InitialList)
			resp := list.Lpop(test.Count)

			if !CheckList(list, test.FinalList) {
				t.Errorf("#%d: Invalid final list", i)
			}

			if !CompareSlices(resp, test.ExpectedResponse) {
				t.Errorf("#%d: Invalid response list", i)
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
		make([]string, 0),
	},
	{
		"empty_list_2",
		make([]string, 0),
		15,
		make([]string, 0),
		make([]string, 0),
	},
	{
		"list_of_one_1",
		[]string{"foo"},
		1,
		make([]string, 0),
		[]string{"foo"},
	},
	{
		"list_of_one_2",
		[]string{"foo"},
		7,
		make([]string, 0),
		[]string{"foo"},
	},
	{
		"list_of_two_1",
		[]string{"foo", "bar"},
		1,
		[]string{"foo"},
		[]string{"bar"},
	},
	{
		"list_of_two_2",
		[]string{"foo", "bar"},
		2,
		make([]string, 0),
		[]string{"bar", "foo"},
	},
	{
		"random_list_1",
		[]string{"foo", "bar", "fizz", "bazz", "golang", "python", "yavascript"},
		4,
		[]string{"foo", "bar", "fizz"},
		[]string{"yavascript", "python", "golang", "bazz"},
	},
	{
		"random_list_2",
		[]string{"foo", "bar", "fizz", "bazz", "golang", "python", "yavascript"},
		15,
		make([]string, 0),
		[]string{"yavascript", "python", "golang", "bazz", "fizz", "bar", "foo"},
	},
	{
		"zero_count",
		[]string{"foo", "bar", "fizz", "bazz", "golang", "python", "yavascript"},
		0,
		[]string{"foo", "bar", "fizz", "bazz", "golang", "python", "yavascript"},
		make([]string, 0),
	},
}

func TestRpop(t *testing.T) {
	for i, test := range RpopTests {
		t.Run(test.Name, func(t *testing.T) {
			list := CreateInitialList(test.InitialList)
			resp := list.Rpop(test.Count)

			if !CheckList(list, test.FinalList) {
				t.Errorf("#%d: Invalid final list", i)
			}

			if !CompareSlices(resp, test.ExpectedResponse) {
				t.Errorf("#%d: Invalid response list", i)
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
