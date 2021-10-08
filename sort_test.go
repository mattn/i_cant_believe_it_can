package i_cant_beleave_it_can

import (
	"reflect"
	"testing"
)

func TestInts(t *testing.T) {
	v := []int{3, 5, 2, 8, 1}
	Sort(v)
	want := []int{1, 2, 3, 5, 8}
	if !reflect.DeepEqual(v, want) {
		t.Fatalf("want %v, but %v:", want, v)
	}
}
