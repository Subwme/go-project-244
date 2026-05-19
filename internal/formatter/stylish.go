package formatter

import (
	"fmt"
	"sort"
	"strings"

	"code/internal/diff"
)

func Stylish(nodes []diff.Node) string {
	return formatNodes(nodes, 1)
}

func formatNodes(nodes []diff.Node, depth int) string {
	var lines []string
	lines = append(lines, "{")

	for _, node := range nodes {
		signIndent := strings.Repeat(" ", depth*4-2)
		plainIndent := strings.Repeat(" ", depth*4)

		switch node.Type {
		case diff.Nested:
			lines = append(lines, fmt.Sprintf("%s%s: %s", plainIndent, node.Key, formatNodes(node.Children, depth+1)))
		case diff.Unchanged:
			lines = append(lines, fmt.Sprintf("%s%s: %s", plainIndent, node.Key, formatValue(node.Value, depth)))
		case diff.Added:
			lines = append(lines, fmt.Sprintf("%s+ %s: %s", signIndent, node.Key, formatValue(node.NewValue, depth)))
		case diff.Removed:
			lines = append(lines, fmt.Sprintf("%s- %s: %s", signIndent, node.Key, formatValue(node.OldValue, depth)))
		case diff.Changed:
			lines = append(lines, fmt.Sprintf("%s- %s: %s", signIndent, node.Key, formatValue(node.OldValue, depth)))
			lines = append(lines, fmt.Sprintf("%s+ %s: %s", signIndent, node.Key, formatValue(node.NewValue, depth)))
		}
	}

	closingIndent := strings.Repeat(" ", (depth-1)*4)
	lines = append(lines, closingIndent+"}")
	return strings.Join(lines, "\n")
}

func formatValue(value interface{}, depth int) string {
	if value == nil {
		return "null"
	}
	if m, ok := value.(map[string]interface{}); ok {
		return formatMap(m, depth)
	}
	return fmt.Sprintf("%v", value)
}

func formatMap(m map[string]interface{}, depth int) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var lines []string
	lines = append(lines, "{")
	for _, k := range keys {
		indent := strings.Repeat(" ", (depth+1)*4)
		lines = append(lines, fmt.Sprintf("%s%s: %s", indent, k, formatValue(m[k], depth+1)))
	}
	closingIndent := strings.Repeat(" ", depth*4)
	lines = append(lines, closingIndent+"}")
	return strings.Join(lines, "\n")
}
