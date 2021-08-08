package utils

import (
	"reflect"
	"testing"
)

func TestReverseKeys(t *testing.T) {
	var testParams = []struct {
		testName string
		input    map[string]string
		output   map[string]string
	}{
		{"Empty input", map[string]string{}, map[string]string{}},
		{"Valid input", map[string]string{"a": "z", "s": "x", "d": "c"}, map[string]string{"z": "a", "x": "s", "c": "d"}},
	}
	for _, testParam := range testParams {
		result := ReverseKeys(testParam.input)
		if !reflect.DeepEqual(result, testParam.output) {
			t.Errorf("%s: Expected: '%v'; Got: '%v'", testParam.testName, testParam.output, result)
		}
	}

	TestReverseKeysPanic(t)
}

func TestReverseKeysPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	// The following is the code under test
	ReverseKeys(map[string]string{"a": "a", "b": "a"})
}
