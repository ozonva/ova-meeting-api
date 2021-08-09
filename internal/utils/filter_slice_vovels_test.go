package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterSlice(t *testing.T) {
	assertions := assert.New(t)

	var testParams = []struct {
		testName string
		input    []string
		output   []string
	}{
		{"Empty input", []string{}, []string{}},
		{"Input with one filtered element", []string{"a", "b", "c"}, []string{"b", "c"}},
		{"Input with all filtered element", []string{"a", "e", "o"}, []string{}},
		{"Input with no filtered element", []string{"b", "c", "d"}, []string{"b", "c", "d"}},
	}
	for _, testParam := range testParams {
		assertions.Equal(testParam.output, FilterSlice(testParam.input), "Should be equal. "+testParam.testName)
	}
}
