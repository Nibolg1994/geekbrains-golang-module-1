package simplenumbers

import (
	"reflect"
	"testing"
)

func TestSimpleNumbers(t *testing.T) {
	tests := []struct {
		name  string
		input int64
		want  []int64
	}{
		{name: "one", input: 1, want: []int64{}},
		{name: "simple", input: 10, want: []int64{2, 3, 5, 7}},
		{name: "two", input: 2, want: []int64{2}},
		{name: "random", input: 38, want: []int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37}},
	}

	for _, tc := range tests {
		got, _ := simpleNumbers(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}

func TestNegativeSimpleNumbers(t *testing.T) {
	_, err := simpleNumbers(-38)
	if err == nil {
		t.Error("Expected  error")
	}
}
