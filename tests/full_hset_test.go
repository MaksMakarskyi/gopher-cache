package tests

import (
	"testing"
)

var HsetTests = []TestCase{
	{
		Name:             "HSET_key_missing_creates_hash",
		ExecuteBefore:    "",
		Message:          "*4\r\n$4\r\nHSET\r\n$14\r\nhset_test_key1\r\n$5\r\nfield\r\n$3\r\nbar\r\n",
		ExpectedResponse: ":1\r\n",
	},
	{
		Name:             "HSET_key_exists_new_field",
		ExecuteBefore:    "*4\r\n$4\r\nHSET\r\n$14\r\nhset_test_key2\r\n$5\r\nfield\r\n$3\r\nbar\r\n",
		Message:          "*4\r\n$4\r\nHSET\r\n$14\r\nhset_test_key2\r\n$6\r\nfield2\r\n$5\r\nvalue\r\n",
		ExpectedResponse: ":1\r\n",
	},
	{
		Name:             "HSET_key_exists_overwrite_field",
		ExecuteBefore:    "*4\r\n$4\r\nHSET\r\n$14\r\nhset_test_key3\r\n$5\r\nfield\r\n$3\r\nbar\r\n",
		Message:          "*4\r\n$4\r\nHSET\r\n$14\r\nhset_test_key3\r\n$5\r\nfield\r\n$3\r\nbaz\r\n",
		ExpectedResponse: ":0\r\n",
	},
	{
		Name:             "HSET_wrong_type_operation",
		ExecuteBefore:    "*4\r\n$5\r\nLPUSH\r\n$19\r\nhset_wrong_type_key\r\n$5\r\nitem1\r\n$5\r\nitem2\r\n",
		Message:          "*4\r\n$4\r\nHSET\r\n$19\r\nhset_wrong_type_key\r\n$5\r\nfield\r\n$3\r\nbar\r\n",
		ExpectedResponse: "-WRONGTYPE Operation against a key holding the wrong kind of value\r\n",
	},
	{
		Name:             "HSET_wrong_number_of_arguments",
		ExecuteBefore:    "",
		Message:          "*5\r\n$4\r\nHSET\r\n$3\r\nfoo\r\n$5\r\nfield\r\n$3\r\nbar\r\n$3\r\nbaz\r\n",
		ExpectedResponse: "-ERR wrong number of arguments for 'HSET' command\r\n",
	},
}

func TestHset(t *testing.T) {
	for i, test := range HsetTests {
		RunTest(test, i, t)
	}
}
