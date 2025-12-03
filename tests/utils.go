package tests

import (
	"net"
	"testing"
)

func ExecuteCommand(cmd string) (string, error) {
	conn, err := net.Dial("tcp", "localhost:6379")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	_, err = conn.Write([]byte(cmd))
	if err != nil {
		return "", err
	}

	resp := make([]byte, 1024)
	n, err := conn.Read(resp)
	if err != nil {
		return "", err
	}

	strResp := string(resp[:n])
	return strResp, nil
}

func RunTest(test TestCase, i int, t *testing.T) {
	t.Run(test.Name, func(t *testing.T) {
		if len(test.ExecuteBefore) > 0 {
			_, err := ExecuteCommand(test.ExecuteBefore)

			if err != nil {
				t.Error("Failed to execute 'ExecuteBefore' command")
			}
		}

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
