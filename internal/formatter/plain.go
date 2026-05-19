package formatter

import (
	"fmt"
	"strings"

	"code/internal/diff"
)

func Plain(nodes []diff.Node) string {
	return strings.Join(formatPlainNodes(nodes, ""), "\n")
}

func formatPlainNodes(nodes []diff.Node, path string) []string {
	var lines []string

	for _, node := range nodes {
		currentPath := node.Key
		if path != "" {
			currentPath = path + "." + node.Key
		}

		switch node.Type {
		case diff.Nested:
			lines = append(lines, formatPlainNodes(node.Children, currentPath)...)
		case diff.Added:
			lines = append(lines, fmt.Sprintf("Property '%s' was added with value: %s", currentPath, plainValue(node.NewValue)))
		case diff.Removed:
			lines = append(lines, fmt.Sprintf("Property '%s' was removed", currentPath))
		case diff.Changed:
			lines = append(lines, fmt.Sprintf("Property '%s' was updated. From %s to %s", currentPath, plainValue(node.OldValue), plainValue(node.NewValue)))
		}
	}

	return lines
}

func plainValue(value interface{}) string {
	if value == nil {
		return "null"
	}
	if _, ok := value.(map[string]interface{}); ok {
		return "[complex value]"
	}
	if s, ok := value.(string); ok {
		return fmt.Sprintf("'%s'", s)
	}
	return fmt.Sprintf("%v", value)
}
