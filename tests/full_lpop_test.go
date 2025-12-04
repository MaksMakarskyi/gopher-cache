package tests

import (
	"testing"
)

var LpopTests = []TestCase{
	{
		Name:             "LPOP_key_missing_returns_nil",
		ExecuteBefore:    "",
		Message:          "*2\r\n$4\r\nLPOP\r\n$11\r\nlpop_testk1\r\n",
		ExpectedResponse: "$-1\r\n",
	},
	{
		Name:             "LPOP_key_exists_pops_first_element",
		ExecuteBefore:    "*4\r\n$5\r\nLPUSH\r\n$11\r\nlpop_testk2\r\n$5\r\napple\r\n$6\r\nbanana\r\n",
		Message:          "*2\r\n$4\r\nLPOP\r\n$11\r\nlpop_testk2\r\n",
		ExpectedResponse: "$6\r\nbanana\r\n",
	},
	{
		Name:             "LPOP_count_pops_multiple_elements",
		ExecuteBefore:    "*5\r\n$5\r\nLPUSH\r\n$15\r\nlpop_count_key1\r\n$5\r\napple\r\n$6\r\nbanana\r\n$6\r\ncherry\r\n",
		Message:          "*3\r\n$4\r\nLPOP\r\n$15\r\nlpop_count_key1\r\n$1\r\n2\r\n",
		ExpectedResponse: "*2\r\n$6\r\ncherry\r\n$6\r\nbanana\r\n",
	},
	{
		Name:             "LPOP_count_greater_than_length",
		ExecuteBefore:    "*4\r\n$5\r\nLPUSH\r\n$15\r\nlpop_count_key2\r\n$5\r\napple\r\n$6\r\nbanana\r\n",
		Message:          "*3\r\n$4\r\nLPOP\r\n$15\r\nlpop_count_key2\r\n$1\r\n5\r\n",
		ExpectedResponse: "*2\r\n$6\r\nbanana\r\n$5\r\napple\r\n",
	},
	{
		Name:             "LPOP_invalid_integer",
		ExecuteBefore:    "*4\r\n$5\r\nLPUSH\r\n$11\r\nlpop_testk3\r\n$5\r\napple\r\n$6\r\nbanana\r\n",
		Message:          "*3\r\n$4\r\nLPOP\r\n$11\r\nlpop_testk3\r\n$3\r\n56A\r\n",
		ExpectedResponse: "-ERR value is not an integer or out of range\r\n",
	},
	{
		Name:             "LPOP_wrong_type_operation",
		ExecuteBefore:    "*4\r\n$4\r\nSADD\r\n$15\r\nlpop_wrong_type\r\n$5\r\nitem1\r\n$5\r\nitem2\r\n",
		Message:          "*2\r\n$4\r\nLPOP\r\n$15\r\nlpop_wrong_type\r\n",
		ExpectedResponse: "-WRONGTYPE Operation against a key holding the wrong kind of value\r\n",
	},
	{
		Name:             "LPOP_wrong_number_of_arguments",
		ExecuteBefore:    "",
		Message:          "*4\r\n$4\r\nLPOP\r\n$3\r\nfoo\r\n$3\r\nbar\r\n$3\r\nbar\r\n",
		ExpectedResponse: "-ERR wrong number of arguments for 'LPOP' command\r\n",
	},
}

func TestLpop(t *testing.T) {
	for i, test := range LpopTests {
		RunTest(test, i, t)
	}
}
