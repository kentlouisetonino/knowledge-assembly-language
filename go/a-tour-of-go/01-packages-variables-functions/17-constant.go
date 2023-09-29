/*
  Constants
    - Declared like variables, but the "const" keyword.
    - Can be character, string, boolean, or numeric values.
    - NOTE: constants cannot be declared using the ":=" syntax.
*/

package main

import "fmt"

const PI = 3.14

func main() {
  fmt.Println("***** Constants *****")
  fmt.Printf("PI: %T, %v \n", PI, PI)
}

