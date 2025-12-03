package tests

import (
	"testing"
)

var HgetTests = []TestCase{
	{
		Name:             "HGET_key_exists_field_exists",
		ExecuteBefore:    "*4\r\n$4\r\nHSET\r\n$14\r\nhget_test_hash\r\n$5\r\nfield\r\n$3\r\nbar\r\n",
		Message:          "*3\r\n$4\r\nHGET\r\n$14\r\nhget_test_hash\r\n$5\r\nfield\r\n",
		ExpectedResponse: "$3\r\nbar\r\n",
	},
	{
		Name:             "HGET_key_exists_field_missing",
		ExecuteBefore:    "*4\r\n$4\r\nHSET\r\n$15\r\nhget_test_hash2\r\n$5\r\nfield\r\n$3\r\nbar\r\n",
		Message:          "*3\r\n$4\r\nHGET\r\n$15\r\nhget_test_hash2\r\n$7\r\nmissing\r\n",
		ExpectedResponse: "$-1\r\n",
	},
	{
		Name:             "HGET_key_missing",
		ExecuteBefore:    "",
		Message:          "*3\r\n$4\r\nHGET\r\n$16\r\nhget_missing_key\r\n$5\r\nfield\r\n",
		ExpectedResponse: "$-1\r\n",
	},
	{
		Name:             "HGET_wrong_type_operation",
		ExecuteBefore:    "*4\r\n$5\r\nLPUSH\r\n$19\r\nhget_wrong_type_key\r\n$5\r\nitem1\r\n$5\r\nitem2\r\n",
		Message:          "*3\r\n$4\r\nHGET\r\n$19\r\nhget_wrong_type_key\r\n$5\r\nfield\r\n",
		ExpectedResponse: "-WRONGTYPE Operation against a key holding the wrong kind of value\r\n",
	},
	{
		Name:             "HGET_wrong_number_of_arguments",
		ExecuteBefore:    "",
		Message:          "*4\r\n$4\r\nHGET\r\n$3\r\nfoo\r\n$5\r\nfield\r\n$3\r\nbar\r\n",
		ExpectedResponse: "-ERR wrong number of arguments for 'HGET' command\r\n",
	},
}

func TestHget(t *testing.T) {
	for i, test := range HgetTests {
		RunTest(test, i, t)
	}
}
