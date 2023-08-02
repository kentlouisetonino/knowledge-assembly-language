/*
  You can also write multiple imports statements, like:
    import "fmt"
    import "math"
  
  But it is a good style to use the factored import statement.
    import {
      "fmt"
      "math"
    }

  You can use fmt.Printf() like in C.
*/

package main
import "fmt" // * Reference: https://pkg.go.dev/fmt
import "math" // * Reference: https://pkg.go.dev/math

func main() {
  fmt.Printf("Now you have %g problems. \n", math.Sqrt(7));
}

