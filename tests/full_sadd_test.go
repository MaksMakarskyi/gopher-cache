package tests

import (
	"testing"
)

var SaddTests = []TestCase{
	{
		Name:             "SADD_key_missing_creates_set",
		ExecuteBefore:    "",
		Message:          "*4\r\n$4\r\nSADD\r\n$14\r\nsadd_test_key1\r\n$5\r\napple\r\n$6\r\nbanana\r\n",
		ExpectedResponse: ":2\r\n",
	},
	{
		Name:             "SADD_key_exists_add_new_members",
		ExecuteBefore:    "*4\r\n$4\r\nSADD\r\n$14\r\nsadd_test_key2\r\n$5\r\napple\r\n$6\r\nbanana\r\n",
		Message:          "*4\r\n$4\r\nSADD\r\n$14\r\nsadd_test_key2\r\n$6\r\ncherry\r\n$5\r\nmango\r\n",
		ExpectedResponse: ":2\r\n",
	},
	{
		Name:             "SADD_key_exists_adding_duplicates",
		ExecuteBefore:    "*4\r\n$4\r\nSADD\r\n$14\r\nsadd_test_key3\r\n$5\r\napple\r\n$6\r\nbanana\r\n",
		Message:          "*4\r\n$4\r\nSADD\r\n$14\r\nsadd_test_key3\r\n$5\r\napple\r\n$6\r\nbanana\r\n",
		ExpectedResponse: ":0\r\n",
	},
	{
		Name:             "SADD_wrong_type_operation",
		ExecuteBefore:    "*4\r\n$5\r\nLPUSH\r\n$19\r\nsadd_wrong_type_key\r\n$5\r\nitem1\r\n$5\r\nitem2\r\n",
		Message:          "*3\r\n$4\r\nSADD\r\n$19\r\nsadd_wrong_type_key\r\n$5\r\napple\r\n",
		ExpectedResponse: "-WRONGTYPE Operation against a key holding the wrong kind of value\r\n",
	},
	{
		Name:             "SADD_wrong_number_of_arguments",
		ExecuteBefore:    "",
		Message:          "*2\r\n$4\r\nSADD\r\n$3\r\nfoo\r\n",
		ExpectedResponse: "-ERR wrong number of arguments for 'SADD' command\r\n",
	},
}

func TestSadd(t *testing.T) {
	for i, test := range SaddTests {
		RunTest(test, i, t)
	}
}
