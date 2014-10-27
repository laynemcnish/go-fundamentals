package main

import (
	"reflect"
	"testing"
)

func Test_goStart_exists(t *testing.T) {
	if goStart != 2007 {
		t.Errorf("Something went wrong with goStart declaration: %v", goStart)
	}
}

func Test_spareChange_exists(t *testing.T) {
	if spareChange != 3.50 {
		t.Errorf("Something went wrong with your spareChange declaration: %v", spareChange)
	}
}

func Test_spareChange_type(t *testing.T) {
	if tp := reflect.TypeOf(spareChange).Kind(); tp != reflect.Float32 {
		t.Errorf("spareChange is the wrong type. Expected Float32, got %T!", spareChange)
	}
}

func Test_cost_exists(t *testing.T) {
	if cost != 0.25 {
		t.Errorf("Something went wrong with your cost declartaion: %v", cost)
	}
}

func Test_cost_type(t *testing.T) {
	if tp := reflect.TypeOf(cost).Kind(); tp != reflect.Float32 {
		t.Errorf("cost is the wrong type. Expected Float32, got %T!", cost)
	}
}

func Test_purchased_exists(t *testing.T) {
	if purchased != 14 {
		t.Errorf("Something went wrong with your purchased declaration %v", purchased)
	}
}

// func Test_circumference(t *testing.T) {
// 	if circumference != float32(diameter)*pi {
// 		t.Errorf("Something went wrong with determining circumference: %v", circumference)
// 	}
//
// }
