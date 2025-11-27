package gopherstring

import "testing"

// GET
type TestCaseGet struct {
	Name          string
	InitialString string
	Expected      string
}

var GetTests = []TestCaseGet{
	{
		"get_1",
		"fizz",
		"fizz",
	},
	{
		"get_2",
		"fizzbazz",
		"fizzbazz",
	},
	{
		"get_3",
		"",
		"",
	},
}

func TestGet(t *testing.T) {
	for i, test := range GetTests {
		t.Run(test.Name, func(t *testing.T) {
			gstring := NewGopherString(test.InitialString)
			got := gstring.Get()

			if got != test.Expected {
				t.Errorf("#%d: Expected: %s, Got: %s", i, test.Expected, got)
			}
		})
	}
}
