package formatter

import (
	"fmt"

	"code/internal/diff"
)

const (
	FormatStylish = "stylish"
	FormatPlain   = "plain"
	FormatJSON    = "json"
)

func Format(nodes []diff.Node, format string) (string, error) {
	switch format {
	case FormatStylish, "":
		return Stylish(nodes)
	case FormatPlain:
		return Plain(nodes)
	case FormatJSON:
		return JSON(nodes)
	default:
		return "", fmt.Errorf("unknown format: %s", format)
	}
}
