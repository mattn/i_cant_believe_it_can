package i_cant_beleave_it_can_test

import (
	"reflect"
	"testing"

	"github.com/mattn/i_cant_beleave_it_can"
)

func TestInts(t *testing.T) {
	v := []int{3, 5, 2, 8, 1}
	i_cant_beleave_it_can.Sort(v)
	want := []int{1, 2, 3, 5, 8}
	if !reflect.DeepEqual(v, want) {
		t.Fatalf("want %v, but %v:", want, v)
	}
}
