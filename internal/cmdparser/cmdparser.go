package cmdparser

import (
	"fmt"
	"strconv"
	"strings"
)

type GopherCommandParser struct{}

func NewGopherCommandParser() *GopherCommandParser {
	return &GopherCommandParser{}
}

// func (gcp *GopherCommandParser) Parse(cmd string) (string, []string, error) {
// 	cursor := 0

// 	return "SET", []string{"foo", "bar"}, nil
// }

func (gcp *GopherCommandParser) ParseBulkString(c *int, s string) (string, error) {
	// Example:
	// - Input: $3\r\nGET\r\n Output: "GET"
	// - Input: $6\r\nfoobar\r\n Output: "foobar"
	// - Input: $dfs\r\nfoobar\r\n Output: "", Error("parse error: invalid length: dfs")

	// Check datatype
	if s[*c] != byte('$') {
		return "", fmt.Errorf("parse error: expected '$', but got: %c", s[*c])
	}
	*c += 1

	// Read length
	crlfStart := strings.Index(s[*c:], "\r\n")
	if crlfStart == -1 {
		return "", fmt.Errorf("parse error: incomplete command (missing length CRLF)")
	}

	lengthStr := s[*c : *c+crlfStart]
	length, err := strconv.Atoi(lengthStr)
	if err != nil {
		return "", fmt.Errorf("parse error: invalid length: %s", lengthStr)
	}
	*c += crlfStart + 2

	if length == -1 {
		return "", nil
	}

	// Read string
	if *c+length+2 > len(s) {
		return "", fmt.Errorf("parse error: buffer too short, expected %d bytes", length)
	}

	strValue := s[*c : *c+length]
	if s[*c+length] != '\r' || s[*c+length+1] != '\n' {
		return "", fmt.Errorf("parse error: expected CRLF at end of payload")
	}

	*c += length + 2

	return strValue, nil
}
