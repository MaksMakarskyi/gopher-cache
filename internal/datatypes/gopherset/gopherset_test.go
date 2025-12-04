package gopherset

import "testing"

func CompareSets(expected, got map[string]bool) bool {
	if len(expected) != len(got) {
		return false
	}

	for key, value := range expected {
		if got[key] != value {
			return false
		}
	}

	return true
}

// SADD
type TestCaseSadd struct {
	Name          string
	InitialSet    map[string]bool
	Args          []string
	FinalSet      map[string]bool
	ExpectedCount int
}

var SaddTests = []TestCaseSadd{
	{
		"empty_set_1",
		make(map[string]bool, 0),
		[]string{"foo"},
		map[string]bool{"foo": true},
		1,
	},
	{
		"empty_set_2",
		make(map[string]bool, 0),
		[]string{"foo", "bar"},
		map[string]bool{"foo": true, "bar": true},
		2,
	},
	{
		"random_set_1",
		map[string]bool{"foo": true, "bar": true},
		[]string{"foo", "bar"},
		map[string]bool{"foo": true, "bar": true},
		0,
	},
	{
		"random_set_2",
		map[string]bool{"foo": true, "fizz": true},
		[]string{"foo", "bar"},
		map[string]bool{"foo": true, "fizz": true, "bar": true},
		1,
	},
	{
		"random_set_3",
		map[string]bool{"bazz": true, "fizz": true},
		[]string{"foo", "bar"},
		map[string]bool{"bazz": true, "fizz": true, "foo": true, "bar": true},
		2,
	},
	{
		"empty_input_1",
		map[string]bool{"bazz": true, "fizz": true},
		[]string{},
		map[string]bool{"bazz": true, "fizz": true},
		0,
	},
	{
		"empty_input_2",
		make(map[string]bool, 0),
		[]string{},
		make(map[string]bool, 0),
		0,
	},
}

func TestSadd(t *testing.T) {
	for i, test := range SaddTests {
		t.Run(test.Name, func(t *testing.T) {
			set := NewGopherSet()
			set.Data = test.InitialSet
			count := set.Sadd(test.Args)

			if !CompareSets(test.FinalSet, set.Data) {
				t.Errorf("#%d: Set Expected: %#v, Got: %#v", i, test.FinalSet, set.Data)
			}

			if count != test.ExpectedCount {
				t.Errorf("#%d: Count Expected: %d, Got: %d", i, test.ExpectedCount, count)
			}
		})
	}
}

// SREM
type TestCaseSrem struct {
	Name          string
	InitialSet    map[string]bool
	Args          []string
	FinalSet      map[string]bool
	ExpectedCount int
}

var SremTests = []TestCaseSrem{
	{
		"empty_set_1",
		make(map[string]bool, 0),
		[]string{"foo"},
		make(map[string]bool, 0),
		0,
	},
	{
		"empty_set_2",
		make(map[string]bool, 0),
		[]string{"foo", "bar"},
		make(map[string]bool, 0),
		0,
	},
	{
		"duplicated_key",
		map[string]bool{"foo": true, "bar": true},
		[]string{"foo", "foo"},
		map[string]bool{"bar": true},
		1,
	},
	{
		"missing_keys",
		map[string]bool{"foo": true, "bar": true},
		[]string{"fizz", "bazz"},
		map[string]bool{"foo": true, "bar": true},
		0,
	},
	{
		"more_args_than_items",
		map[string]bool{"foo": true, "bar": true},
		[]string{"fizz", "bazz", "foo", "bar"},
		make(map[string]bool, 0),
		2,
	},
	{
		"empty_input",
		map[string]bool{"foo": true, "bar": true},
		make([]string, 0),
		map[string]bool{"foo": true, "bar": true},
		0,
	},
}

func TestSrem(t *testing.T) {
	for i, test := range SremTests {
		t.Run(test.Name, func(t *testing.T) {
			set := NewGopherSet()
			set.Data = test.InitialSet
			count := set.Srem(test.Args)

			if !CompareSets(test.FinalSet, set.Data) {
				t.Errorf("#%d: Set Expected: %#v, Got: %#v", i, test.FinalSet, set.Data)
			}

			if count != test.ExpectedCount {
				t.Errorf("#%d: Count Expected: %d, Got: %d", i, test.ExpectedCount, count)
			}
		})
	}
}

// SISMEMBER
type TestCaseSismember struct {
	Name       string
	InitialSet map[string]bool
	Key        string
	Expected   int
}

var SismemberTests = []TestCaseSismember{
	{
		"empty_set",
		make(map[string]bool, 0),
		"foo",
		0,
	},
	{
		"missing_key",
		map[string]bool{"foo": true, "bar": true},
		"fizz",
		0,
	},
	{
		"wrong_case_key",
		map[string]bool{"foo": true, "bar": true},
		"fOo",
		0,
	},
	{
		"existing_key_1",
		map[string]bool{"foo": true, "bar": true},
		"bar",
		1,
	},
	{
		"existing_key_2",
		map[string]bool{"foo": true, "bar": true},
		"foo",
		1,
	},
}

func TestSismember(t *testing.T) {
	for i, test := range SismemberTests {
		t.Run(test.Name, func(t *testing.T) {
			set := NewGopherSet()
			set.Data = test.InitialSet
			got := set.Sismember(test.Key)

			if got != test.Expected {
				t.Errorf("#%d: Expected: %d, Got: %d", i, test.Expected, got)
			}
		})
	}
}

// SCARD
type TestCaseScard struct {
	Name       string
	InitialSet map[string]bool
	Expected   int
}

var ScardTests = []TestCaseScard{
	{
		"empty_set",
		make(map[string]bool, 0),
		0,
	},
	{
		"set_of_one",
		map[string]bool{"foo": true},
		1,
	},
	{
		"random_set_1",
		map[string]bool{"foo": true, "bar": true, "fizz": true},
		3,
	},
	{
		"random_set_2",
		map[string]bool{"foo": true, "bar": true, "fizz": true, "python": true, "javascript": true},
		5,
	},
}

func TestScard(t *testing.T) {
	for i, test := range ScardTests {
		t.Run(test.Name, func(t *testing.T) {
			set := NewGopherSet()
			set.Data = test.InitialSet
			got := set.Scard()

			if got != test.Expected {
				t.Errorf("#%d: Expected: %d, Got: %d", i, test.Expected, got)
			}
		})
	}
}
