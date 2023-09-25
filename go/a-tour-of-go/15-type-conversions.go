/*
  Type Conversions
    - The expression T(v) converts the value v to the type T.
    - Note: Unlike in C, in Go assigment between items of different
      type requires an explicit conversion.
*/

package main

import "fmt"

func numericConversion() {
  var i int = 42
  var f float64 = float64(i) 
  var u uint = uint(f)

  fmt.Println("***** Numeric Conversion *****")
  fmt.Println("i: ", i)
  fmt.Println("f: ", f)
  fmt.Println("u: ", u)
  fmt.Println("\n")
}

func shortNumericConversion() {
  i := 100
  f := float64(i)
  u := uint(f)

  fmt.Println("***** Short Numeric Conversion *****")
  fmt.Println("i: ", i)
  fmt.Println("f: ", f)
  fmt.Println("u: ", u)
  fmt.Println("\n")

}

func main() {
  numericConversion()
  shortNumericConversion()
}

