package utils

import (
	"reflect"
	"testing"
)

func TestFilterSlice(t *testing.T) {
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
		result := FilterSlice(testParam.input)
		if !reflect.DeepEqual(result, testParam.output) {
			t.Errorf("%s: Expected: '%v'; Got: '%v'", testParam.testName, testParam.output, result)
		}
	}
}
