package main

import (
  "reflect"
  "testing"
)

func Test_triumverate_values(t *testing.T) {
  if Caesar != 1 {
    t.Errorf("Something went wrong with Caesar iota: %v", Caesar)
  }
  if Crassus != 2 {
    t.Errorf("Something went wrong with Crassus iota: %v", Crassus)
  }
  if Pompeius != 3 {
    t.Errorf("Something went wrong with Pompeius iota: %v", Pompeius)
  }
}
