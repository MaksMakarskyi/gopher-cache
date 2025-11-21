package cmdparser

import (
	"testing"
)

// type TestCase struct {
// 	Input        string
// 	ExpectedName string
// 	ExpectedArgs []string
// 	ShouldFail   bool
// }

// var tests = []TestCase{
// 	{
// 		Input:        "*3\r\n$3\r\nSET\r\n$3\r\nfoo\r\n$3\r\nbar\r\n",
// 		ExpectedName: "SET",
// 		ExpectedArgs: []string{"foo", "bar"},
// 		ShouldFail:   false,
// 	},
// 	{
// 		Input:        "*3\r\n$3\r\nGET\r\n$3\r\nfoo\r\n",
// 		ExpectedName: "GET",
// 		ExpectedArgs: []string{"foo"},
// 		ShouldFail:   false,
// 	},
// }

// func TestGopherCommandParser(t *testing.T) {
// 	for i, test := range tests {
// 		parser := NewGopherCommandParser(test.Input)
// 		name, args, err := parser.Parse()

// 		if test.ShouldFail {
// 			if err == nil {
// 				t.Errorf("#%d: Expected error for input '%s', but got none", i, test.Input)
// 			}
// 			continue
// 		}

// 		if err != nil {
// 			t.Errorf("#%d: Unexpected error for input '%s': %v", i, test.Input, err)
// 			continue
// 		}

// 		if name != test.ExpectedName {
// 			t.Errorf("#%d: Name mismatch. Got '%s', want '%s'", i, name, test.ExpectedName)
// 		}

// 		if !reflect.DeepEqual(args, test.ExpectedArgs) {
// 			t.Errorf("#%d: Args mismatch. Got %v, want %v", i, args, test.ExpectedArgs)
// 		}
// 	}
// }

type ParseBulkStringTestCase struct {
	Name           string
	Input          string
	ExpectedOutput string
	ShouldFail     bool
	ExpectedCursor int
}

var parsBulkStringTests = []ParseBulkStringTestCase{
	{
		Name:           "invalid_prefix",
		Input:          "*3\r\nGET\r\n",
		ExpectedOutput: "",
		ShouldFail:     true,
		ExpectedCursor: 0,
	},
	{
		Name:           "invalid_prefix_2",
		Input:          "$3\rGET\r\n",
		ExpectedOutput: "",
		ShouldFail:     true,
		ExpectedCursor: 1,
	},
	{
		Name:           "invalid_sufix",
		Input:          "$3\r\nGET\r",
		ExpectedOutput: "",
		ShouldFail:     true,
		ExpectedCursor: 4,
	},
	{
		Name:           "invalid_length",
		Input:          "$sdsf\r\nGET\r\n",
		ExpectedOutput: "",
		ShouldFail:     true,
		ExpectedCursor: 1,
	},
	{
		Name:           "mismatched_length_text_shorter",
		Input:          "$4\r\nGET\r\n",
		ExpectedOutput: "",
		ShouldFail:     true,
		ExpectedCursor: 4,
	},
	{
		Name:           "mismatched_length_text_longer",
		Input:          "$4\r\nGETGET\r\n",
		ExpectedOutput: "",
		ShouldFail:     true,
		ExpectedCursor: 4,
	},
	{
		Name:           "no_crlf",
		Input:          "$4\rGETGET\n",
		ExpectedOutput: "",
		ShouldFail:     true,
		ExpectedCursor: 1,
	},
	{
		Name:           "short_input",
		Input:          "$100\r\nGET\r\n",
		ExpectedOutput: "",
		ShouldFail:     true,
		ExpectedCursor: 6,
	},
	{
		Name:           "simple",
		Input:          "$3\r\nGET\r\n",
		ExpectedOutput: "GET",
		ShouldFail:     false,
		ExpectedCursor: 9,
	},
	{
		Name:           "null_string",
		Input:          "$-1\r\n",
		ExpectedOutput: "",
		ShouldFail:     false,
		ExpectedCursor: 5,
	},
	{
		Name:           "longer",
		Input:          "$6\r\nfoobar\r\n",
		ExpectedOutput: "foobar",
		ShouldFail:     false,
		ExpectedCursor: 12,
	},
	{
		Name:           "multidigit_length",
		Input:          "$12\r\nfoobarfoobar\r\n",
		ExpectedOutput: "foobarfoobar",
		ShouldFail:     false,
		ExpectedCursor: 19,
	},
	{
		Name:           "binary_safety",
		Input:          "$4\r\nA\r\nB\r\n",
		ExpectedOutput: "A\r\nB",
		ShouldFail:     false,
		ExpectedCursor: 10,
	},
	{
		Name:           "empty_string",
		Input:          "$0\r\n\r\n",
		ExpectedOutput: "",
		ShouldFail:     false,
		ExpectedCursor: 6,
	},
}

func TestGopherCommandParser(t *testing.T) {
	parser := NewGopherCommandParser()

	for i, tc := range parsBulkStringTests {
		t.Run(tc.Name, func(t *testing.T) {
			c := 0
			strValue, err := parser.ParseBulkString(&c, tc.Input)

			if tc.ShouldFail {
				if err == nil {
					t.Errorf("#%d: Expected error for input '%s', but got none", i, tc.Input)
				}
			} else if err != nil {
				t.Errorf("#%d: Unexpected error for input '%s': %v", i, tc.Input, err)
			}

			if strValue != tc.ExpectedOutput {
				t.Errorf("#%d: Output mismatch. Got '%s', want '%s'", i, strValue, tc.ExpectedOutput)
			}

			if c != tc.ExpectedCursor {
				t.Errorf("#%d: Cursor mismatch.\nGot:  %d\nWant: %d", i, c, tc.ExpectedCursor)
			}
		})
	}
}
