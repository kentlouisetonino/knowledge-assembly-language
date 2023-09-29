/*
  Variables
    - The "var" statement declares a list variables; as in function
      argument lists, the type is last.
    - A "var" statement can be at package or function level.

  Package References
    fmt   : https://pkg.go.dev/fmt
*/

package main
import "fmt"

var c, python, java bool;

func main() {
  var golang int;
  
  fmt.Println("Variable (go) =", golang); 
  fmt.Println("Variable (c) =", c);
  fmt.Println("Variable (python) =", python);
  fmt.Println("Variable (java) =", java);
}

