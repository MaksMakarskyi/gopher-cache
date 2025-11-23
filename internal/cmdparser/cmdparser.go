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

func (gcp *GopherCommandParser) Parse(cmd string) (string, []any, error) {
	cursor := 0

	if cursor >= len(cmd) {
		return "", nil, fmt.Errorf("ERR incomplete message")
	}

	args, err := gcp.ParseArray(&cursor, cmd)
	if err != nil {
		return "", nil, err
	}

	commandType, ok := args[0].(string)
	if !ok {
		return "", nil, fmt.Errorf("ERR Parse error invalid command type")
	}

	return commandType, args[1:], nil
}

func (gcp *GopherCommandParser) _parse(cursor *int, s string) (any, error) {
	if *cursor >= len(s) {
		return nil, fmt.Errorf("ERR incomplete message")
	}

	switch s[*cursor] {
	case byte('$'):
		return gcp.ParseBulkString(cursor, s)
	case byte('*'):
		return gcp.ParseArray(cursor, s)
	default:
		return "", fmt.Errorf("ERR Parse error at %d, expected dtype byte", cursor)
	}
}

func (gcp *GopherCommandParser) ParseArray(cursor *int, s string) ([]any, error) {
	if s[*cursor] != byte('*') {
		return nil, fmt.Errorf("ERR Parse error expected '*', but got: %c", s[*cursor])
	}
	*cursor += 1

	// Read length
	crlfStart := strings.Index(s[*cursor:], "\r\n")
	if crlfStart == -1 {
		return nil, fmt.Errorf("ERR Parse error incomplete command (missing length CRLF)")
	}

	lengthStr := s[*cursor : *cursor+crlfStart]
	length, err := strconv.Atoi(lengthStr)
	if err != nil {
		return nil, fmt.Errorf("ERR Parse error invalid length: %s", lengthStr)
	}
	*cursor += crlfStart + 2

	// Check length for safety
	if length > 1024*1024 {
		return nil, fmt.Errorf("ERR array too large")
	}

	// Read array items
	items := make([]any, length)
	for i := range length {
		item, err := gcp._parse(cursor, s)

		if err != nil {
			return nil, err
		}

		items[i] = item
	}

	return items, nil
}

func (gcp *GopherCommandParser) ParseBulkString(cursor *int, s string) (string, error) {
	// Example:
	// - Input: $3\r\nGET\r\n Output: "GET"
	// - Input: $6\r\nfoobar\r\n Output: "foobar"
	// - Input: $dfs\r\nfoobar\r\n Output: "", Error("parse error: invalid length: dfs")

	// Check datatype
	if s[*cursor] != byte('$') {
		return "", fmt.Errorf("ERR Parse error expected '$', but got: %c", s[*cursor])
	}
	*cursor += 1

	// Read length
	crlfStart := strings.Index(s[*cursor:], "\r\n")
	if crlfStart == -1 {
		return "", fmt.Errorf("ERR Parse error incomplete command (missing length CRLF)")
	}

	lengthStr := s[*cursor : *cursor+crlfStart]
	length, err := strconv.Atoi(lengthStr)
	if err != nil {
		return "", fmt.Errorf("ERR Parse error invalid length: %s", lengthStr)
	}
	*cursor += crlfStart + 2

	if length == -1 {
		return "", nil
	}

	// Read string
	if *cursor+length+2 > len(s) {
		return "", fmt.Errorf("ERR Parse error buffer too short, expected %d bytes", length)
	}

	strValue := s[*cursor : *cursor+length]
	if s[*cursor+length] != '\r' || s[*cursor+length+1] != '\n' {
		return "", fmt.Errorf("ERR Parse error expected CRLF at end of payload")
	}

	*cursor += length + 2

	return strValue, nil
}
