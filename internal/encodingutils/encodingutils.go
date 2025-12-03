package encodingutils

import (
	"fmt"
	"strings"
)

func FormatSimpleString(s string) string {
	return fmt.Sprintf("+%s\r\n", s)
}

func FormatSimpleError(s string) string {
	return fmt.Sprintf("-%s\r\n", s)
}

func FormatInteger(i int) string {
	return fmt.Sprintf(":%d\r\n", i)
}

func FormatBulkString(s string) string {
	return fmt.Sprintf("$%d\r\n%s\r\n", len(s), s)
}

func FormatArray(strs []string) string {
	arrayPrefix := fmt.Sprintf("*%d\r\n", len(strs))

	formattedStrings := make([]string, len(strs)+1)
	formattedStrings[0] = arrayPrefix
	for i, s := range strs {
		if s == "" {
			formattedStrings[i] = GetNullBulkString()
		} else {
			formattedStrings[i] = FormatBulkString(s)
		}
	}

	return strings.Join(formattedStrings, "")
}

func GetNullBulkString() string {
	return "$-1\r\n"
}
