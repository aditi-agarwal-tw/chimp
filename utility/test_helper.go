package utility

import (
	"reflect"
	"testing"
)

// AssertEqual is equality tester
func AssertEqual(t *testing.T,
	actual interface{},
	expected interface{},
	errorMessage string) {
	if actual != expected {
		t.Errorf("Test failed, got: %s, want: %s, message: %s.", actual, expected, errorMessage)
	}
}

// AssertContains checks if the object is inside the given array
func AssertContains(t *testing.T,
	arr []interface{}, obj interface{}, errorMessage string) {
	if !contains(arr, obj) {
		t.Errorf("Test failed, object: %s, is not contained in array, message: %s.", obj, errorMessage)
	}
}

func contains(s []interface{}, e interface{}) bool {
	for _, a := range s {
		if reflect.DeepEqual(a, e) {
			return true
		}
	}
	return false
}