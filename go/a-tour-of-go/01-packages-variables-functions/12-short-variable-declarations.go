/*
  Short Variable Declarations
  - Inside a function, the ":=" short assignment statement can be used in place
    of a "var" declaration witch implicit type.
  - Ouside a function, every statement begins with a keyword ("var", "func",
    and so on) and so the ":=" construct is not available.
*/

package main
import "fmt"

func main() {
  // Short variable declaration.
  // Declare and assign.
  x, y := 1, 2

  fmt.Println(x, y)
}

