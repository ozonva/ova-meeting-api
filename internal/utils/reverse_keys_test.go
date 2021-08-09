package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseKeys(t *testing.T) {
	assertions := assert.New(t)
	var testParams = []struct {
		testName string
		input    map[string]string
		output   map[string]string
	}{
		{"Empty input", map[string]string{}, map[string]string{}},
		{"Valid input", map[string]string{"a": "z", "s": "x", "d": "c"}, map[string]string{"z": "a", "x": "s", "c": "d"}},
	}
	for _, testParam := range testParams {
		assertions.Equal(testParam.output, ReverseKeys(testParam.input), "Should be equal. "+testParam.testName)
	}
	assert.Panics(t, func() { ReverseKeys(map[string]string{"a": "a", "b": "a"}) })
}
