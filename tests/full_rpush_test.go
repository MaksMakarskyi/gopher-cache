package tests

import (
	"testing"
)

var RpushTests = []TestCase{
	{
		Name:             "RPUSH_key_missing_creates_list",
		ExecuteBefore:    "",
		Message:          "*4\r\n$5\r\nRPUSH\r\n$15\r\nrpush_test_key1\r\n$5\r\napple\r\n$6\r\nbanana\r\n",
		ExpectedResponse: ":2\r\n",
	},
	{
		Name:             "RPUSH_key_exists_appends_values",
		ExecuteBefore:    "*4\r\n$5\r\nRPUSH\r\n$15\r\nrpush_test_key2\r\n$5\r\napple\r\n$6\r\nbanana\r\n",
		Message:          "*4\r\n$5\r\nRPUSH\r\n$15\r\nrpush_test_key2\r\n$6\r\ncherry\r\n$5\r\nmango\r\n",
		ExpectedResponse: ":4\r\n",
	},
	{
		Name:             "RPUSH_key_exists_adds_no_new_elements_when_empty_input",
		ExecuteBefore:    "*4\r\n$5\r\nRPUSH\r\n$15\r\nrpush_test_key3\r\n$5\r\napple\r\n$6\r\nbanana\r\n",
		Message:          "*2\r\n$5\r\nRPUSH\r\n$15\r\nrpush_test_key3\r\n",
		ExpectedResponse: "-ERR wrong number of arguments for 'RPUSH' command\r\n",
	},
	{
		Name:             "RPUSH_wrong_type_operation",
		ExecuteBefore:    "*4\r\n$4\r\nSADD\r\n$20\r\nrpush_wrong_type_key\r\n$5\r\napple\r\n$6\r\nbanana\r\n",
		Message:          "*3\r\n$5\r\nRPUSH\r\n$20\r\nrpush_wrong_type_key\r\n$5\r\napple\r\n",
		ExpectedResponse: "-WRONGTYPE Operation against a key holding the wrong kind of value\r\n",
	},
	{
		Name:             "RPUSH_wrong_number_of_arguments",
		ExecuteBefore:    "",
		Message:          "*2\r\n$5\r\nRPUSH\r\n$3\r\nfoo\r\n",
		ExpectedResponse: "-ERR wrong number of arguments for 'RPUSH' command\r\n",
	},
}

func TestRpush(t *testing.T) {
	for i, test := range RpushTests {
		RunTest(test, i, t)
	}
}
