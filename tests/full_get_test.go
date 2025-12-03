package tests

import (
	"testing"
)

var GetTests = []TestCase{
	{
		Name:             "GET_key_does_not_exist",
		ExecuteBefore:    "",
		Message:          "*2\r\n$3\r\nGET\r\n$26\r\nget_my_non_existing_string\r\n",
		ExpectedResponse: "$-1\r\n",
	},
	{
		Name:             "GET_key_does_exist",
		ExecuteBefore:    "*3\r\n$3\r\nSET\r\n$22\r\nget_my_existing_string\r\n$3\r\nbar\r\n",
		Message:          "*2\r\n$3\r\nGET\r\n$22\r\nget_my_existing_string\r\n",
		ExpectedResponse: "$3\r\nbar\r\n",
	},
	{
		Name:             "GET_wrong_number_of_args",
		ExecuteBefore:    "",
		Message:          "*4\r\n$3\r\nGET\r\n$3\r\nfoo\r\n$4\r\nfizz\r\n$4\r\nbazz\r\n",
		ExpectedResponse: "-ERR wrong number of arguments for 'GET' command\r\n",
	},
	{
		Name:             "GET_wrong_type_operation",
		ExecuteBefore:    "*4\r\n$5\r\nLPUSH\r\n$18\r\nget_wrong_type_key\r\n$5\r\nitem1\r\n$5\r\nitem2\r\n",
		Message:          "*2\r\n$3\r\nGET\r\n$18\r\nget_wrong_type_key\r\n",
		ExpectedResponse: "-WRONGTYPE Operation against a key holding the wrong kind of value\r\n",
	},
}

func TestGet(t *testing.T) {
	for i, test := range GetTests {
		RunTest(test, i, t)
	}
}
