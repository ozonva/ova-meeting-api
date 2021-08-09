package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitSlice(t *testing.T) {
	assertions := assert.New(t)
	var testParams = []struct {
		testName string
		input    []string
		chunk    int
		output   [][]string
	}{
		{"Empty splice", []string{}, 1, [][]string{}},
		{"Split chunk size 1", []string{"a", "b", "c"}, 1, [][]string{{"a"}, {"b"}, {"c"}}},
		{"Split chunk size 2", []string{"a", "b", "c"}, 2, [][]string{{"a", "b"}, {"c"}}},
		{"Split chunk size 3", []string{"a", "b", "c"}, 3, [][]string{{"a", "b", "c"}}},
		{"Split chunk size 4", []string{"a", "b", "c"}, 4, [][]string{{"a", "b", "c"}}},
	}
	for _, testParam := range testParams {
		assertions.Equal(testParam.output, SplitSlice(testParam.input, uint(testParam.chunk)), "Should be equal. "+testParam.testName)
	}

	assert.Panics(t, func() { SplitSlice([]string{"a", "b", "c"}, 0) })
}
