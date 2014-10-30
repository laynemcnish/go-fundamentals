package main

import (
	"reflect"
	"testing"
)

func Test_fishMeasure_attributes(t *testing.T) {
	if fish.CommonName != "cutbow" {
		t.Errorf("fishMeasure Common Name throwing an error: %v", fish.CommonName)
	}
	if fish.ScienceName != "Oncorhynchus clarki x mykiss" {
		t.Errorf("fishMeasure Science Name throwing an error: %v", fish.ScienceName)
	}
	if fish.Length != 39 {
		t.Errorf("fishMeasure Length throwing an error: %v", fish.Length)
	}
	if fish.Weight != 0.75 {
		t.Errorf("fishMeasure Weight throwing an error: %v", fish.Weight)
	}
}

func Test_fishMeasure_type(t *testing.T) {
	if reflect.TypeOf(fish).Kind() != reflect.Struct {
		t.Errorf("expected fishMeasure to be struct, got %T", fish)
	}
}
