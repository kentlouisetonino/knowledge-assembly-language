/*
  Type Inference
    - When declaring a variable without specifying an explicit type (either
      by using the ":=" syntax or "var =" expression syntax, the variable's
      type is inferred from the value on the right hand side.
*/

package main

import "fmt"

// When the right hand side of the declaration is type,
// the new variable is of that same type.
func sample1() {
  var i int
  j := i

  fmt.Println("***** Sample #1 *****")
  fmt.Println("i: ", i)
  fmt.Println("j: ", j)
  fmt.Println("\n")
}

// When the right hand side contains untyped numeric constant, the new
// variable may be an int, float64, or complex128 depending ont the precesion
// of the constant.
func sample2() {
  i := 42
  f := 3.142
  g := 0.867 + 0.5i


  fmt.Println("***** Sample #2 *****")
  fmt.Printf("i = %T, %v \n", i, i)
  fmt.Printf("j = %T, %v \n", f, f)
  fmt.Printf("g = %T, %v \n", g, g)
  fmt.Println("\n")
}

func main() {
  sample1()
  sample2()
}

