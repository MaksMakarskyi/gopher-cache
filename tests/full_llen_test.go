package tests

import (
	"testing"
)

var LlenTests = []TestCase{
	{
		Name:             "LLEN_key_missing_returns_zero",
		ExecuteBefore:    "",
		Message:          "*2\r\n$4\r\nLLEN\r\n$11\r\nllen_testk1\r\n",
		ExpectedResponse: ":0\r\n",
	},
	{
		Name:             "LLEN_key_exists_returns_length",
		ExecuteBefore:    "*4\r\n$5\r\nRPUSH\r\n$11\r\nllen_testk2\r\n$5\r\napple\r\n$6\r\nbanana\r\n",
		Message:          "*2\r\n$4\r\nLLEN\r\n$11\r\nllen_testk2\r\n",
		ExpectedResponse: ":2\r\n",
	},
	{
		Name:             "LLEN_wrong_type_operation",
		ExecuteBefore:    "*4\r\n$4\r\nSADD\r\n$15\r\nllen_wrong_type\r\n$5\r\nitem1\r\n$5\r\nitem2\r\n",
		Message:          "*2\r\n$4\r\nLLEN\r\n$15\r\nllen_wrong_type\r\n",
		ExpectedResponse: "-WRONGTYPE Operation against a key holding the wrong kind of value\r\n",
	},
	{
		Name:             "LLEN_wrong_number_of_arguments",
		ExecuteBefore:    "",
		Message:          "*3\r\n$4\r\nLLEN\r\n$3\r\nfoo\r\n$3\r\nbar\r\n",
		ExpectedResponse: "-ERR wrong number of arguments for 'LLEN' command\r\n",
	},
}

func TestLlen(t *testing.T) {
	for i, test := range LlenTests {
		RunTest(test, i, t)
	}
}
