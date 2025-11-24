package tests

import (
	"net"
	"testing"
)

type TestCase struct {
	Name             string
	Message          string
	ExpectedResponse string
}

var tests = []TestCase{
	{
		Name:             "valid_set_cmd",
		Message:          "*3\r\n$3\r\nSET\r\n$3\r\nfoo\r\n$3\r\nbar\r\n",
		ExpectedResponse: "OK",
	},
	{
		Name:             "valid_get_cmd",
		Message:          "*2\r\n$3\r\nGET\r\n$3\r\nfoo\r\n",
		ExpectedResponse: "bar",
	},
	{
		Name:             "valid_set_cmd_2",
		Message:          "*3\r\n$3\r\nSET\r\n$3\r\nfoo\r\n$9\r\nbarbarbar\r\n",
		ExpectedResponse: "OK",
	},
	{
		Name:             "valid_get_cmd_2",
		Message:          "*2\r\n$3\r\nGET\r\n$3\r\nfoo\r\n",
		ExpectedResponse: "barbarbar",
	},
}

func TestGopherCache(t *testing.T) {
	for i, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			conn, err := net.Dial("tcp", "localhost:6379")
			if err != nil {
				t.Errorf("Failed to connect to the server")
			}
			defer conn.Close()

			_, err = conn.Write([]byte(test.Message))
			if err != nil {
				t.Fatalf("failed to write: %v", err)
			}

			resp := make([]byte, 1024)
			n, err := conn.Read(resp)
			if err != nil {
				t.Errorf("#%d: Unexpected error for input '%s': %v", i, test.Message, err)
			}

			strResp := string(resp[:n])
			if strResp != test.ExpectedResponse {
				t.Errorf("#%d: Response mismatch. Got '%s', want '%s'", i, strResp, test.ExpectedResponse)
			}
		})

	}
}
