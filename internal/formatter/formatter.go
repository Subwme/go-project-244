package formatter

import (
	"fmt"

	"code/internal/diff"
)

func Format(nodes []diff.Node, format string) (string, error) {
	switch format {
	case "stylish", "":
		return Stylish(nodes), nil
	case "plain":
		return Plain(nodes), nil
	case "json":
		return JSON(nodes)
	default:
		return "", fmt.Errorf("unknown format: %s", format)
	}
}
