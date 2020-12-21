package rules

import (
	"bytes"
	"github.com/thedevsaddam/gojsonq/v2"
)

func LimitsMemory(json []byte) int {
	spec := getSpecSelector(json)
	found := 0

	paths := gojsonq.New().Reader(bytes.NewReader(json)).
		From(spec + ".containers").
		Only("resources.limits.memory")

	if paths != nil {
		found += len(paths.([]interface{}))
	}

	return found
}
