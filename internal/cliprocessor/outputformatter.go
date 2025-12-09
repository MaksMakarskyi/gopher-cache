package cliprocessor

import (
	"fmt"
	"strings"

	"github.com/MaksMakarskyi/gopher-cache/internal/cmdparser"
	"github.com/MaksMakarskyi/gopher-cache/internal/encodingutils"
)

type OutputFormatter struct {
	CommandParser *cmdparser.GopherCommandParser
}

func NewOutputFormatter(cp *cmdparser.GopherCommandParser) *OutputFormatter {
	return &OutputFormatter{CommandParser: cp}
}

func (of *OutputFormatter) Format(s string) (string, error) {
	switch s[0] {
	case '+':
		return of.FormatSimpleString(s), nil
	case '-':
		return of.FormatSimpleError(s), nil
	case ':':
		return of.FormatInteger(s), nil
	case '$':
		return of.FormatBulkString(s)
	case '*':
		return of.FormatArray(s)
	default:
		return s, nil
	}
}

func (of *OutputFormatter) FormatSimpleString(s string) string {
	s = strings.TrimLeft(s, "+")
	s = strings.TrimRight(s, "\r\n")

	return fmt.Sprintf("\"%s\"", s)
}

func (of *OutputFormatter) FormatSimpleError(s string) string {
	s = strings.TrimLeft(s, "-")
	s = strings.TrimRight(s, "\r\n")

	return fmt.Sprintf("(error) %s", s)
}

func (of *OutputFormatter) FormatInteger(s string) string {
	s = strings.TrimLeft(s, ":")
	s = strings.TrimRight(s, "\r\n")

	return fmt.Sprintf("(integer) %s", s)
}

func (of *OutputFormatter) FormatBulkString(s string) (string, error) {
	if s == encodingutils.GetNullBulkString() {
		return "(nil)", nil
	}

	c := 0
	str, err := of.CommandParser.ParseBulkString(&c, s)

	if err != nil {
		return str, err
	}

	return fmt.Sprintf("\"%s\"", str), nil
}

func (of *OutputFormatter) FormatArray(s string) (string, error) {
	c := 0
	values, err := of.CommandParser.ParseArray(&c, s)
	if err != nil {
		return "", err
	}

	strValues, err := cmdparser.ExpectStrings(values)
	if err != nil {
		return "", err
	}

	formattedArray := make([]string, len(strValues))
	for i, v := range strValues {
		if v == "" {
			v = "(nil)"
		} else {
			v = fmt.Sprintf("\"%s\"", v)
		}

		formattedArray[i] = fmt.Sprintf("%d) %s", i+1, v)
	}

	return strings.Join(formattedArray, "\n"), nil
}
