package main

import (
	"reflect"
	"testing"
)

func Test_sandwich_exists(t *testing.T) {
	if sandwich != "pork chop" {
		t.Errorf("Something went wrong with sandwich declaration: %v", sandwich)
	}
}

func Test_sandwich_type(t *testing.T) {
	if tp := reflect.TypeOf(sandwich).Kind(); tp != reflect.String {
		t.Errorf("sandwich is the wrong type. Expected string, got %T", sandwich)
	}
}

func Test_rank_vars_value(t *testing.T) {
	if first != 1 {
		t.Errorf("Something went wrong with first declaration: %v", first)
	}
	if second != 2 {
		t.Errorf("Something went wrong with second declaration: %v", second)
	}
	if third != 3 {
		t.Errorf("Something went wrong with third declaration: %v", third)
	}
	if fourth != 4 {
		t.Errorf("Something went wrong with fourth declaration: %v", fourth)
	}
}

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

func Test_x_exists(t *testing.T) {
	if x != -1 {
		t.Errorf("Something went wrong with x declaration: %v", x)
	}
}

func Test_y_exists(t *testing.T) {
	if y != -3 {
		t.Errorf("Something went wrong with y declaration: %v", y)
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

func Test_char_exists(t *testing.T) {
	if char != 'Q' {
		t.Errorf("Something went wrong with char declaration: %q", char)
	}
}

func Test_char_type(t *testing.T) {
	if tp := reflect.TypeOf(char).Kind(); tp != reflect.Int32 {
		t.Errorf("char is the wrong type. Expected rune, got %T", char)
	}
}

func Test_letterFromChar_exists(t *testing.T) {
	if letterFromChar != "Q" {
		t.Errorf("Something went wrong with letterFromChar declaration: %v", letterFromChar)
	}
}

func Test_letterFromChar_type(t *testing.T) {
	if tp := reflect.TypeOf(letterFromChar).Kind(); tp != reflect.String {
		t.Errorf("letterFromChar is the wrong type. Expected string, got %T", letterFromChar)
	}
}

func Test_numberFromChar_exists(t *testing.T) {
	if numberFromChar != 81 {
		t.Errorf("Something went wrong with numberFromChar declaration: %v", numberFromChar)
	}
}

func Test_numberFromChar_type(t *testing.T) {
	if tp := reflect.TypeOf(numberFromChar).Kind(); tp != reflect.Int64 {
		t.Errorf("numberFromChar is the wrong type. Expected string, got %T", numberFromChar)
	}
}

func Test_cryBaby_exists(t *testing.T) {
	if cryBaby != "QQ" {
		t.Errorf("Something went wrong with cryBaby declaration: %v", cryBaby)
	}
}

func Test_charIntString_exists(t *testing.T) {
	if charIntString != "81" {
		t.Errorf("Something went wrong with charIntString declaration: %v", charIntString)
	}
}
