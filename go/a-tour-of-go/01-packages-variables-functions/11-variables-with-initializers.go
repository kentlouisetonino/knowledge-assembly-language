/*
  var
    - A declaration that can include initializers, one per variable.
    - If an initializer is present, the type can be omitted; the variable will
      take the type of the initializer.

  Documentation Reference:
    fmt   : https://pkg.go.dev/fmt
*/

package main
import "fmt"

var i, j = 1, 2

func main() {
  var c, python, java = true, false, "fuck!"

  // Types are omitted.
  fmt.Println(c, python, java)
  // Types are not omitted.
  fmt.Println(i, j)
}

