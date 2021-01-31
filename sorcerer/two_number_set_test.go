package main

import (
	"reflect"
	"testing"
)

func TestGetWeight(t *testing.T) {

	if !reflect.DeepEqual(GetWeight([]int{1, 2, 3, 4, 5}, 5), []int{1, 4}) {
		t.Error("Test failed")
	}
}

func TestGetWeightEmpty(t *testing.T) {

	if !reflect.DeepEqual(GetWeight([]int{1, 2}, 5), []int{}) {
		t.Error("Test failed")
	}
}
