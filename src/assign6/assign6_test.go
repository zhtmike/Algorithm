package main

import (
	"reflect"
	"testing"
)

func TestMedianMaintenance(t *testing.T) {
	arr := []int{5, 7, 1, 4, 10}
	ans := []int{5, 5, 5, 4, 5}
	medians := MedianMaintenance(arr)
	if !reflect.DeepEqual(medians, ans) {
		t.Error("Expected ", ans, " got ", medians)
	}
}
