/*
  Export Names
    - In Go, a name is exported if it begins with a capital letter.
      For example, "Pizza" is an exported name, as is "Pi", which
      is exported frorm the "math" package.

  Package References
    fmt   : https://pkg.go.dev/fmt
    math  : https://pkg.go.dev/math

*/

package main
import "fmt"
import "math"

func main() {
  // Pi (The ratio of a circle circumference to its diameter).
  fmt.Println("Pi   :", math.Pi);

  // Phi (The basic for the Golden Ration, Section, or Mean).
  fmt.Println("Phi  :", math.Phi);
}

