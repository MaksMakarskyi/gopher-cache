package tests

import (
	"testing"
)

var ScardTests = []TestCase{
	{
		Name:             "SCARD_key_missing_returns_zero",
		ExecuteBefore:    "",
		Message:          "*2\r\n$5\r\nSCARD\r\n$12\r\nscard_testk1\r\n",
		ExpectedResponse: ":0\r\n",
	},
	{
		Name:             "SCARD_key_exists_returns_size",
		ExecuteBefore:    "*5\r\n$4\r\nSADD\r\n$12\r\nscard_testk2\r\n$5\r\napple\r\n$6\r\nbanana\r\n$6\r\ncherry\r\n",
		Message:          "*2\r\n$5\r\nSCARD\r\n$12\r\nscard_testk2\r\n",
		ExpectedResponse: ":3\r\n",
	},
	{
		Name:             "SCARD_wrong_type_operation",
		ExecuteBefore:    "*4\r\n$5\r\nLPUSH\r\n$16\r\nscard_wrong_type\r\n$5\r\nitem1\r\n$5\r\nitem2\r\n",
		Message:          "*2\r\n$5\r\nSCARD\r\n$16\r\nscard_wrong_type\r\n",
		ExpectedResponse: "-WRONGTYPE Operation against a key holding the wrong kind of value\r\n",
	},
	{
		Name:             "SCARD_wrong_number_of_arguments",
		ExecuteBefore:    "",
		Message:          "*3\r\n$5\r\nSCARD\r\n$3\r\nfoo\r\n$3\r\nbar\r\n",
		ExpectedResponse: "-ERR wrong number of arguments for 'SCARD' command\r\n",
	},
}

func TestScard(t *testing.T) {
	for i, test := range ScardTests {
		RunTest(test, i, t)
	}
}
