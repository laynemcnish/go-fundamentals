package main

import "strconv"

/*
  Resources: https://code.google.com/p/go-wiki/wiki/Range
             http://golang.org/pkg/strconv/

  TODO: create a function frontRange which takes slice of strings. It should return a single string which
  contains a numbered list of the strings such that ["one", "two"] would yield "1. One\n2.Two\n"
  HINT: import "strconv" and use Itoa */
func frontRange(arr []string) string {
  var finalStr string
  for i, str := range arr {
    finalStr = finalStr + strconv.Itoa(i + 1) + ". " + str + "\n"
  }
  return finalStr
}
