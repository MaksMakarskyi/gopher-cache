package tests

import (
	"testing"
)

var SetTests = []TestCase{
	{
		Name:             "SET_key_does_not_exist",
		ExecuteBefore:    "",
		Message:          "*3\r\n$3\r\nSET\r\n$13\r\nset_my_string\r\n$3\r\nbar\r\n",
		ExpectedResponse: "+OK\r\n",
	},
	{
		Name:             "SET_key_does_exist",
		ExecuteBefore:    "",
		Message:          "*3\r\n$3\r\nSET\r\n$26\r\nset_my_non_existing_string\r\n$4\r\nbazz\r\n",
		ExpectedResponse: "+OK\r\n",
	},
	{
		Name:             "SET_wrong_number_of_args",
		ExecuteBefore:    "",
		Message:          "*4\r\n$3\r\nSET\r\n$13\r\nset_my_string\r\n$4\r\nfizz\r\n$4\r\nbazz\r\n",
		ExpectedResponse: "-ERR wrong number of arguments for 'SET' command\r\n",
	},
}

func TestSet(t *testing.T) {
	for i, test := range SetTests {
		RunTest(test, i, t)
	}
}
