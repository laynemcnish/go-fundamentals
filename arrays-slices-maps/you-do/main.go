package main

import "fmt"
/*
  Resources: http://blog.golang.org/go-slices-usage-and-internals
             https://blog.golang.org/go-maps-in-action
*/

func main() {
	// create an array called fours with a length 10 with integers holding the values
	// divisible by four from 4 to 40 (ie 4, 8, 12, etc)
  fours := [10]int{}
    var x, i int
    for  {
          if i % 4 == 0 {
            fours[x] = i
            x += 1
          }
          i++
          if x == len(fours) {
            break
          }
        }
  fmt.Println(fours)

	// print the slice of fours holding all the numbers which have a 2 in the tens digit
  slice := make([]int, len(fours))

  for index,element := range fours {
    if element % 10 == 2 {
      slice[index] = element
    }
}
  fmt.Println(slice)

	// create a slice named s with the make keyword- give it int type with a length of 1
  s := make([]int, 1)
  fmt.Println(s)
	// append onto s the integers 1, 2, and 3
  s = append(s, 1, 2, 3)
  fmt.Println(s)
	// print s- print it before the append and after to see the difference. Can you describe what is happening?

	/* create two slices aTeam and bTeam, both string types. Inside aTeam declare the strings "Murdock", "Baracus"
	"Hannibal", "Peck". In bTeam declare "Dunning" and "Kruger". Then append bTeam onto the end of aTeam
	in a variable named Team. Print "Full Roster:" followed by Team
	HINT: You'll find '...' to be more than a trailing sentiment in a morose email... */
  aTeam := []string{"Murdock", "Baracus", "Hannibal", "Peck"}
  bTeam := []string{"Dunning", "Kruger"}
  Team := append(aTeam, bTeam...)
  fmt.Println(Team)


	// create a map classes, make it string keys with integer values. Set "History" to 4, "Biology" to 2,
	// "Calculus" to 5, and "Spanish" to 3.
  classes := map[string]int{"History":4,"Biology":2,"Calculus":5,"Spanish":3}

	// Print the value of classes["Earth Sciences"] - can you explain your result?
  fmt.Println(classes["Earth Sciences"])

	// declare a boolean called ok which will be true only if classes["Earth Sciences"] exists
_, ok := classes["Earth Sciences"]
fmt.Println("Ok", ok)
	// Print the value of ok

}
