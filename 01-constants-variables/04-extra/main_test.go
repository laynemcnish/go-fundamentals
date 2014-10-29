package main

import (
  "reflect"
  "testing"
)

func Test_nutrient_values(t *testing.T) {
  if minerals != 1 {
    t.Errorf("Something went wrong with minerals iota: %v", minerals)
  }
  if fat != 0 {
    t.Errorf("Something went wrong with fat iota: %v", fat)
  }
  if vitamins != 2 {
    t.Errorf("Something went wrong with vitamins iota: %v", vitamins)
  }
  if protein != 2 {
    t.Errorf("Something went wrong with protein iota: %v", protein)
  }
  if carbohydrates != 3 {
    t.Errorf("Something went wrong with carbohydrates iota: %v", carbohydrates)
  }
  if water != 6 {
    t.Errorf("Something went wrong with water iota: %v", water)
  }
}
