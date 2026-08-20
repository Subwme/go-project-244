package diff

import (
	"testing"
)

func TestBuildWithArrays(t *testing.T) {
	data1 := map[string]interface{}{
		"name":    "Ivan",
		"hobbies": []interface{}{"go", "js"},
	}

	data2 := map[string]interface{}{
		"name":    "Ivan",
		"hobbies": []interface{}{"go", "rust"},
	}

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Code crashed with panic! Error: %v", r)
		}
	}()

	nodes := Build(data1, data2)

	if len(nodes) != 2 {
		t.Errorf("Expected 2 nodes, got: %d", len(nodes))
	}
}

func TestBuildWithDifferentNumericTypes(t *testing.T) {
	data1 := map[string]interface{}{
		"timeout": float64(50),
	}

	data2 := map[string]interface{}{
		"timeout": int(50),
	}

	nodes := Build(data1, data2)

	if len(nodes) != 1 {
		t.Fatalf("Expected 1 node, got: %d", len(nodes))
	}

	if nodes[0].Type != Unchanged {
		t.Errorf("Expected type 'unchanged' for equal numbers of different types, got: '%s'", nodes[0].Type)
	}
}
