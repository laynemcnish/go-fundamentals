package main

import (
	"reflect"
	"testing"
)

func Test_typedFloat_type(t *testing.T) {
	if tp := reflect.TypeOf(typedFloat).Kind(); tp != reflect.Float32 {
		t.Errorf("typedFloat is the wrong type. Expected Float32, got %T", typedFloat)
	}
}

func Test_untypedFloat_type(t *testing.T) {
	if tp := reflect.TypeOf(untypedFloat).Kind(); tp != reflect.Float64 {
		t.Errorf("untypedFloat is the wrong type. Expected Float64, got %T", untypedFloat)
	}
}

func Test_bigSixtyFour_exists(t *testing.T) {
	if bigSixtyFour != 9223372036854775807 {
		t.Errorf("Something went wrong with bigSixtyFour declaration: %v", bigSixtyFour)
	}
}

func Test_bigSixtyFour_type(t *testing.T) {
	if tp := reflect.TypeOf(bigSixtyFour).Kind(); tp != reflect.Int64 {
		t.Errorf("bigSixtyFour is the wrong type. Expected int64, got %T", bigSixtyFour)
	}
}

func Test_negaNumber_exists(t *testing.T) {
	if negaNumber != -2147483648 {
		t.Errorf("Something went wrong with negaNumber declaration: %v", negaNumber)
	}
}

func Test_negaNumber_type(t *testing.T) {
	if tp := reflect.TypeOf(negaNumber).Kind(); tp != reflect.Int32 {
		t.Errorf("negaNumber is the wrong type. Expected int32, got %T", negaNumber)
	}
}
