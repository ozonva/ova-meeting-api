package utils

import (
	"reflect"
	"testing"
)

func TestSplitSlice(t *testing.T) {
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
		result := SplitSlice(testParam.input, uint(testParam.chunk))
		if !reflect.DeepEqual(result, testParam.output) {
			t.Errorf("%s: Expected: '%v'; Got: '%v'", testParam.testName, testParam.output, result)
		}
	}

	TestSplitSlicePanic(t)
}

func TestSplitSlicePanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	// The following is the code under test
	_ = SplitSlice([]string{"a", "b", "c"}, 0)
}
