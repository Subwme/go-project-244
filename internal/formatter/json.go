package formatter

import (
	"encoding/json"

	"code/internal/diff"
)

func JSON(nodes []diff.Node) (string, error) {
	result := nodesToMap(nodes)
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func nodesToMap(nodes []diff.Node) map[string]interface{} {
	result := make(map[string]interface{})
	for _, node := range nodes {
		switch node.Type {
		case diff.Nested:
			result[node.Key] = map[string]interface{}{
				"type":     "nested",
				"children": nodesToMap(node.Children),
			}
		case diff.Added:
			result[node.Key] = map[string]interface{}{
				"type":  "added",
				"value": node.NewValue,
			}
		case diff.Removed:
			result[node.Key] = map[string]interface{}{
				"type":  "removed",
				"value": node.OldValue,
			}
		case diff.Unchanged:
			result[node.Key] = map[string]interface{}{
				"type":  "unchanged",
				"value": node.Value,
			}
		case diff.Changed:
			result[node.Key] = map[string]interface{}{
				"type": "changed",
				"old":  node.OldValue,
				"new":  node.NewValue,
			}
		}
	}
	return result
}
