package reflection

import (
	"fmt"
	"testing"
)

func TestPutToMap(t *testing.T) {
	asMap := map[string]any{
		"a": 1,
		"b": []int{1, 2, 3, 4, 5, 6},
		"c": map[string]any{
			"c1": 1,
			"c2": 2,
		},
	}
	fmt.Printf("asMap: %#v\n", asMap)
}
