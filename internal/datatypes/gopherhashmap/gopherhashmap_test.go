package gopherhashmap

import (
	"testing"
)

func CompareMaps(expected, got map[string]string) bool {
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

func CompareSlices(expected, got []string) bool {
	if len(expected) != len(got) {
		return false
	}

	for i := range len(expected) {
		if expected[i] != got[i] {
			return false
		}
	}

	return true
}

// HSET
type TestCaseHset struct {
	Name       string
	Args       []string
	InitialMap map[string]string
	FinalMap   map[string]string
	ShouldFail bool
}

var HsetTests = []TestCaseHset{
	{
		"empty_map_1",
		[]string{"foo", "bar"},
		make(map[string]string, 0),
		map[string]string{"foo": "bar"},
		false,
	},
	{
		"empty_map_2",
		[]string{"foo", "bar", "fizz", "bazz"},
		make(map[string]string, 0),
		map[string]string{"foo": "bar", "fizz": "bazz"},
		false,
	},
	{
		"random_map_1",
		[]string{"fizz", "bazz"},
		map[string]string{"foo": "bar"},
		map[string]string{"foo": "bar", "fizz": "bazz"},
		false,
	},
	{
		"random_map_2",
		[]string{"python", "true", "javascript", "false"},
		map[string]string{"foo": "bar", "fizz": "bazz"},
		map[string]string{"foo": "bar", "fizz": "bazz", "python": "true", "javascript": "false"},
		false,
	},
	{
		"invalid_args_1",
		[]string{"python", "true", "javascript"},
		map[string]string{"foo": "bar", "fizz": "bazz"},
		map[string]string{"foo": "bar", "fizz": "bazz"},
		true,
	},
	{
		"invalid_args_2",
		[]string{"foo"},
		map[string]string{"foo": "bar", "fizz": "bazz"},
		map[string]string{"foo": "bar", "fizz": "bazz"},
		true,
	},
	{
		"empty_input",
		[]string{},
		map[string]string{"foo": "bar", "fizz": "bazz"},
		map[string]string{"foo": "bar", "fizz": "bazz"},
		false,
	},
}

func TestHset(t *testing.T) {
	for i, test := range HsetTests {
		t.Run(test.Name, func(t *testing.T) {
			gmap := NewGopherMap()
			gmap.Data = test.InitialMap
			err := gmap.Hset(test.Args)

			if test.ShouldFail && err == nil {
				t.Errorf("#%d: Expected error", i)

			} else if !test.ShouldFail && err != nil {
				t.Errorf("#%d: Unexpected error: %s", i, err.Error())
			} else if !CompareMaps(test.FinalMap, gmap.Data) {
				t.Errorf("#%d: Expected: %#v, Got: %#v", i, test.FinalMap, gmap.Data)
			}
		})
	}
}

// HGET
type TestCaseHget struct {
	Name       string
	InitialMap map[string]string
	Key        string
	Expected   string
}

var HgetTests = []TestCaseHget{
	{
		"empty_map",
		make(map[string]string, 0),
		"foo",
		"",
	},
	{
		"existing_key_1",
		map[string]string{"foo": "bar", "fizz": "bazz"},
		"foo",
		"bar",
	},
	{
		"existing_key_2",
		map[string]string{"foo": "bar", "fizz": "bazz"},
		"fizz",
		"bazz",
	},
	{
		"missing_key_1",
		map[string]string{"foo": "bar", "fizz": "bazz"},
		"go",
		"",
	},
	{
		"missing_key_2",
		map[string]string{"foo": "bar", "fizz": "bazz"},
		"fOo",
		"",
	},
}

func TestHget(t *testing.T) {
	for i, test := range HgetTests {
		t.Run(test.Name, func(t *testing.T) {
			gmap := NewGopherMap()
			gmap.Data = test.InitialMap
			got := gmap.Hget(test.Key)

			if got != test.Expected {
				t.Errorf("#%d: Expected: %s, Got: %s", i, test.Expected, got)
			}
		})
	}
}

// HMGET
type TestCaseHmget struct {
	Name       string
	InitialMap map[string]string
	Keys       []string
	Expected   []string
}

var HmgetTests = []TestCaseHmget{
	{
		"empty_map",
		make(map[string]string, 0),
		[]string{"foo", "bar"},
		[]string{"", ""},
	},
	{
		"missing_keys_1",
		map[string]string{"foo": "bar", "fizz": "bazz"},
		[]string{"foo", "bar"},
		[]string{"bar", ""},
	},
	{
		"missing_keys_2",
		map[string]string{"foo": "bar", "fizz": "bazz"},
		[]string{"foO", "bar"},
		[]string{"", ""},
	},
	{
		"empty_input",
		map[string]string{"foo": "bar", "fizz": "bazz"},
		make([]string, 0),
		make([]string, 0),
	},
}

func TestHmget(t *testing.T) {
	for i, test := range HmgetTests {
		t.Run(test.Name, func(t *testing.T) {
			gmap := NewGopherMap()
			gmap.Data = test.InitialMap
			got := gmap.Hmget(test.Keys)

			if !CompareSlices(test.Expected, got) {
				t.Errorf("#%d: Expected: %s, Got: %s", i, test.Expected, got)
			}
		})
	}
}
