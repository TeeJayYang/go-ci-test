package sort

import (
	"reflect"
	"testing"
)

func TestSortBubble(t *testing.T) {
	input := []int{4, 2, 1, 5, 3}
	expected := []int{1, 2, 3, 4, 5}
	SortBubble(input)
	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Expected sorted slice to be %v, got %v", expected, input)
	}
}
