package formatter

import (
	"encoding/json"

	"code/internal/diff"
)

func JSON(nodes []diff.Node) (string, error) {
	data, err := json.MarshalIndent(nodes, "", "    ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}
