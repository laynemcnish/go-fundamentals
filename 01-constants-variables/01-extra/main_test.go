package main

import (
  "reflect"
  "testing"
)

func Test_intIota_values(t *testing.T) {
  if Agape != 0 {
    t.Errorf("Something went wrong with Agape iota: %v", Agape)
  }
  if Phileo != 1 {
    t.Errorf("Something went wrong with Phileo iota: %v", Phileo)
  }
  if Storge != 2 {
    t.Errorf("Something went wrong with Storge iota: %v", Storge)
  }
  if Eros != 3 {
    t.Errorf("Something went wrong with Eros iota: %v", Eros)
  }
}
