package tests

import (
	"testing"
)

var HmgetTests = []TestCase{
	{
		Name:             "HMGET_key_exists_some_fields_exist",
		ExecuteBefore:    "*4\r\n$4\r\nHSET\r\n$16\r\nhmget_test_hash1\r\n$5\r\nfield\r\n$3\r\nbar\r\n",
		Message:          "*4\r\n$5\r\nHMGET\r\n$16\r\nhmget_test_hash1\r\n$5\r\nfield\r\n$7\r\nmissing\r\n",
		ExpectedResponse: "*2\r\n$3\r\nbar\r\n$-1\r\n",
	},
	{
		Name:             "HMGET_key_exists_no_fields_exist",
		ExecuteBefore:    "*4\r\n$4\r\nHSET\r\n$16\r\nhmget_test_hash2\r\n$4\r\nname\r\n$3\r\nbob\r\n",
		Message:          "*4\r\n$5\r\nHMGET\r\n$16\r\nhmget_test_hash2\r\n$3\r\nage\r\n$6\r\nsalary\r\n",
		ExpectedResponse: "*2\r\n$-1\r\n$-1\r\n",
	},
	{
		Name:             "HMGET_key_missing",
		ExecuteBefore:    "",
		Message:          "*4\r\n$5\r\nHMGET\r\n$17\r\nhmget_missing_key\r\n$3\r\none\r\n$5\r\ntwooo\r\n",
		ExpectedResponse: "*2\r\n$-1\r\n$-1\r\n",
	},
	{
		Name:             "HMGET_wrong_type_operation",
		ExecuteBefore:    "*4\r\n$5\r\nLPUSH\r\n$20\r\nhmget_wrong_type_key\r\n$5\r\nitem1\r\n$5\r\nitem2\r\n",
		Message:          "*4\r\n$5\r\nHMGET\r\n$20\r\nhmget_wrong_type_key\r\n$5\r\nfield\r\n$3\r\nbar\r\n",
		ExpectedResponse: "-WRONGTYPE Operation against a key holding the wrong kind of value\r\n",
	},
	{
		Name:             "HMGET_wrong_number_of_arguments",
		ExecuteBefore:    "",
		Message:          "*2\r\n$5\r\nHMGET\r\n$3\r\nfoo\r\n",
		ExpectedResponse: "-ERR wrong number of arguments for 'HMGET' command\r\n",
	},
}

func TestHmget(t *testing.T) {
	for i, test := range HmgetTests {
		RunTest(test, i, t)
	}
}
