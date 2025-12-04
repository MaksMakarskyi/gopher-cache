package tests

import (
	"testing"
)

var SremTests = []TestCase{
	{
		Name:             "SREM_key_missing_returns_zero",
		ExecuteBefore:    "",
		Message:          "*4\r\n$4\r\nSREM\r\n$14\r\nsrem_test_key1\r\n$5\r\napple\r\n$6\r\nbanana\r\n",
		ExpectedResponse: ":0\r\n",
	},
	{
		Name:             "SREM_key_exists_remove_existing_members",
		ExecuteBefore:    "*4\r\n$4\r\nSADD\r\n$14\r\nsrem_test_key2\r\n$5\r\napple\r\n$6\r\nbanana\r\n",
		Message:          "*4\r\n$4\r\nSREM\r\n$14\r\nsrem_test_key2\r\n$6\r\nbanana\r\n$5\r\napple\r\n",
		ExpectedResponse: ":2\r\n",
	},
	{
		Name:             "SREM_key_exists_remove_non_existing_members",
		ExecuteBefore:    "*4\r\n$4\r\nSADD\r\n$14\r\nsrem_test_key3\r\n$5\r\napple\r\n$6\r\nbanana\r\n",
		Message:          "*4\r\n$4\r\nSREM\r\n$14\r\nsrem_test_key3\r\n$6\r\norange\r\n$5\r\npeach\r\n",
		ExpectedResponse: ":0\r\n",
	},
	{
		Name:             "SREM_key_exists_remove_mixed_existing_and_missing",
		ExecuteBefore:    "*4\r\n$4\r\nSADD\r\n$14\r\nsrem_test_key4\r\n$5\r\napple\r\n$6\r\nbanana\r\n",
		Message:          "*5\r\n$4\r\nSREM\r\n$14\r\nsrem_test_key4\r\n$6\r\nbanana\r\n$5\r\nmango\r\n$5\r\napple\r\n",
		ExpectedResponse: ":2\r\n",
	},
	{
		Name:             "SREM_wrong_type_operation",
		ExecuteBefore:    "*4\r\n$5\r\nLPUSH\r\n$19\r\nsrem_wrong_type_key\r\n$5\r\nitem1\r\n$5\r\nitem2\r\n",
		Message:          "*3\r\n$4\r\nSREM\r\n$19\r\nsrem_wrong_type_key\r\n$5\r\napple\r\n",
		ExpectedResponse: "-WRONGTYPE Operation against a key holding the wrong kind of value\r\n",
	},
	{
		Name:             "SREM_wrong_number_of_arguments",
		ExecuteBefore:    "",
		Message:          "*2\r\n$4\r\nSREM\r\n$3\r\nfoo\r\n",
		ExpectedResponse: "-ERR wrong number of arguments for 'SREM' command\r\n",
	},
}

func TestSrem(t *testing.T) {
	for i, test := range SremTests {
		RunTest(test, i, t)
	}
}
