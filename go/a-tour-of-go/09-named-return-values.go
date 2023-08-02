/*
  Named Return Values
    - Go's return values may be named, If so, they are treated as
      variables defined at the top of the function.
    - These names should be used to document the meaning of the
      return values.
    - A "return" statement without arguments returns the named
      return values. This is known as "naked" return. The naked
      return statements should be used only in short functions,
      as with the example shown below.

  Package References
    fmt   : https://pkg.go.dev/fmt
*/

package main
import "fmt"

func split(sum int) (x, y int) {
  x = sum * 4 / 9;
  y = sum - x;
  return;
}

func main() {
  int1, int2 := split(17);

  fmt.Println("1st Return =", int1);
  fmt.Println("2nd Return =", int2);
}

