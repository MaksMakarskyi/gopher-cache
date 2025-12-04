package tests

import (
	"testing"
)

var RpopTests = []TestCase{
	{
		Name:             "RPOP_key_missing_returns_nil",
		ExecuteBefore:    "",
		Message:          "*2\r\n$4\r\nRPOP\r\n$11\r\nrpop_testk1\r\n",
		ExpectedResponse: "$-1\r\n",
	},
	{
		Name:             "RPOP_key_exists_pops_last_element",
		ExecuteBefore:    "*4\r\n$5\r\nRPUSH\r\n$11\r\nrpop_testk2\r\n$5\r\napple\r\n$6\r\nbanana\r\n",
		Message:          "*2\r\n$4\r\nRPOP\r\n$11\r\nrpop_testk2\r\n",
		ExpectedResponse: "$6\r\nbanana\r\n",
	},
	{
		Name:             "RPOP_count_pops_multiple_elements",
		ExecuteBefore:    "*5\r\n$5\r\nRPUSH\r\n$15\r\nrpop_count_key1\r\n$5\r\napple\r\n$6\r\nbanana\r\n$6\r\ncherry\r\n",
		Message:          "*3\r\n$4\r\nRPOP\r\n$15\r\nrpop_count_key1\r\n$1\r\n2\r\n",
		ExpectedResponse: "*2\r\n$6\r\ncherry\r\n$6\r\nbanana\r\n",
	},
	{
		Name:             "RPOP_count_greater_than_length",
		ExecuteBefore:    "*4\r\n$5\r\nRPUSH\r\n$15\r\nrpop_count_key2\r\n$5\r\napple\r\n$6\r\nbanana\r\n",
		Message:          "*3\r\n$4\r\nRPOP\r\n$15\r\nrpop_count_key2\r\n$1\r\n4\r\n",
		ExpectedResponse: "*2\r\n$6\r\nbanana\r\n$5\r\napple\r\n",
	},
	{
		Name:             "RPOP_invalid_integer",
		ExecuteBefore:    "*4\r\n$5\r\nRPUSH\r\n$11\r\nrpop_testk3\r\n$5\r\napple\r\n$6\r\nbanana\r\n",
		Message:          "*3\r\n$4\r\nRPOP\r\n$11\r\nrpop_testk3\r\n$3\r\n56A\r\n",
		ExpectedResponse: "-ERR value is not an integer or out of range\r\n",
	},
	{
		Name:             "RPOP_wrong_type_operation",
		ExecuteBefore:    "*4\r\n$4\r\nSADD\r\n$15\r\nrpop_wrong_type\r\n$5\r\nitem1\r\n$5\r\nitem2\r\n",
		Message:          "*2\r\n$4\r\nRPOP\r\n$15\r\nrpop_wrong_type\r\n",
		ExpectedResponse: "-WRONGTYPE Operation against a key holding the wrong kind of value\r\n",
	},
	{
		Name:             "RPOP_wrong_number_of_arguments",
		ExecuteBefore:    "",
		Message:          "*4\r\n$4\r\nRPOP\r\n$3\r\nfoo\r\n$3\r\nbar\r\n$3\r\nbar\r\n",
		ExpectedResponse: "-ERR wrong number of arguments for 'RPOP' command\r\n",
	},
}

func TestRpop(t *testing.T) {
	for i, test := range RpopTests {
		RunTest(test, i, t)
	}
}
