package diff

import (
	"reflect"
	"sort"
)

type Type string

const (
	Added     Type = "added"
	Removed   Type = "removed"
	Unchanged Type = "unchanged"
	Changed   Type = "changed"
	Nested    Type = "nested"
)

type Node struct {
	Key      string      `json:"key"`
	Type     Type        `json:"type"`
	Value    interface{} `json:"value"`
	OldValue interface{} `json:"oldValue"`
	NewValue interface{} `json:"newValue"`
	Children []Node      `json:"children"`
}

func Build(data1, data2 map[string]interface{}) []Node {
	keys := collectKeys(data1, data2)
	sort.Strings(keys)

	var nodes []Node
	for _, key := range keys {
		val1, in1 := data1[key]
		val2, in2 := data2[key]

		map1, isMap1 := val1.(map[string]interface{})
		map2, isMap2 := val2.(map[string]interface{})

		switch {
		case in1 && in2 && isMap1 && isMap2:
			nodes = append(nodes, Node{Key: key, Type: Nested, Children: Build(map1, map2)})
		case in1 && in2 && !isMap1 && !isMap2 && areEqual(val1, val2):
			nodes = append(nodes, Node{Key: key, Type: Unchanged, Value: val1})
		case in1 && in2:
			nodes = append(nodes, Node{Key: key, Type: Changed, OldValue: val1, NewValue: val2})
		case in1:
			nodes = append(nodes, Node{Key: key, Type: Removed, OldValue: val1})
		default:
			nodes = append(nodes, Node{Key: key, Type: Added, NewValue: val2})
		}
	}
	return nodes
}

func collectKeys(data1, data2 map[string]interface{}) []string {
	seen := make(map[string]struct{})
	for k := range data1 {
		seen[k] = struct{}{}
	}
	for k := range data2 {
		seen[k] = struct{}{}
	}
	keys := make([]string, 0, len(seen))
	for k := range seen {
		keys = append(keys, k)
	}
	return keys
}

func areEqual(v1, v2 interface{}) bool {
	f1, isNum1 := convertToFloat(v1)
	f2, isNum2 := convertToFloat(v2)

	if isNum1 && isNum2 {
		return f1 == f2
	}

	return reflect.DeepEqual(v1, v2)
}

func convertToFloat(v interface{}) (float64, bool) {
	switch val := v.(type) {
	case float64:
		return val, true
	case int:
		return float64(val), true
	case int64:
		return float64(val), true
	case float32:
		return float64(val), true
	}
	return 0, false
}
