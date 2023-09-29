/*
  Constants
    - Declared like variables, but the "const" keyword.
    - Can be character, string, boolean, or numeric values.
    - NOTE: constants cannot be declared using the ":=" syntax.
*/

package main

import "fmt"

const PI = 3.14
const TEST = "TEST"
const ONE = 1


func main() {
  fmt.Println("***** Constants *****")
  fmt.Printf("PI: %T, %v \n", PI, PI)
  fmt.Printf("TEST: %T, %v \n", TEST, TEST)
  fmt.Printf("ONE: %T, %v \n", ONE, ONE)
}

