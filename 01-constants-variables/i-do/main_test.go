package main

import (
	"reflect"
	"testing"
)

func Test_pi_exists(t *testing.T) {
	if pi != 3.14159265359 {
		t.Errorf("Something went wrong with pi declaration: %v", pi)
	}
}

func Test_diameter_exists(t *testing.T) {
	if diameter != 10 {
		t.Errorf("Something went wrong with your diameter declaration: %v", diameter)
	}
}

func Test_diameter_type(t *testing.T) {
	if tp := reflect.TypeOf(diameter).Kind(); tp != reflect.Int {
		t.Errorf("diameter is the wrong type. Expected Int, got %T!", diameter)
	}
}

func Test_circumference(t *testing.T) {
	if circumference != float32(diameter)*pi {
		t.Errorf("Something went wrong with determining circumference: %v", circumference)
	}

}
