package tests

import (
	"testing"
)

var SismemberTests = []TestCase{
	{
		Name:             "SISMEMBER_key_missing_returns_zero",
		ExecuteBefore:    "",
		Message:          "*3\r\n$9\r\nSISMEMBER\r\n$14\r\nsismember_key1\r\n$5\r\napple\r\n",
		ExpectedResponse: ":0\r\n",
	},
	{
		Name:             "SISMEMBER_key_exists_member_exists",
		ExecuteBefore:    "*4\r\n$4\r\nSADD\r\n$14\r\nsismember_key2\r\n$5\r\napple\r\n$6\r\nbanana\r\n",
		Message:          "*3\r\n$9\r\nSISMEMBER\r\n$14\r\nsismember_key2\r\n$6\r\nbanana\r\n",
		ExpectedResponse: ":1\r\n",
	},
	{
		Name:             "SISMEMBER_key_exists_member_missing",
		ExecuteBefore:    "*4\r\n$4\r\nSADD\r\n$14\r\nsismember_key3\r\n$5\r\napple\r\n$6\r\nbanana\r\n",
		Message:          "*3\r\n$9\r\nSISMEMBER\r\n$14\r\nsismember_key3\r\n$6\r\norange\r\n",
		ExpectedResponse: ":0\r\n",
	},
	{
		Name:             "SISMEMBER_wrong_type_operation",
		ExecuteBefore:    "*4\r\n$5\r\nLPUSH\r\n$20\r\nsismember_wrong_type\r\n$5\r\nitem1\r\n$5\r\nitem2\r\n",
		Message:          "*3\r\n$9\r\nSISMEMBER\r\n$20\r\nsismember_wrong_type\r\n$5\r\napple\r\n",
		ExpectedResponse: "-WRONGTYPE Operation against a key holding the wrong kind of value\r\n",
	},
	{
		Name:             "SISMEMBER_wrong_number_of_arguments",
		ExecuteBefore:    "",
		Message:          "*2\r\n$9\r\nSISMEMBER\r\n$3\r\nfoo\r\n",
		ExpectedResponse: "-ERR wrong number of arguments for 'SISMEMBER' command\r\n",
	},
}

func TestSismember(t *testing.T) {
	for i, test := range SismemberTests {
		RunTest(test, i, t)
	}
}
