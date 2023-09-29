/*
  If and Else
    - Variables declared inside an "if" short statement are
      also available inside any of the "else" blocks.
*/

package main

import "fmt"
import "math"

func pow(x, n, limit float64) float64 {
  if v := math.Pow(x, n); v < limit  {
    return v
  } else {
    return limit
  }
}

func main() {
  fmt.Println("Return Value: ", pow(3, 2, 10))
  fmt.Println("Return Limit: ", pow(3, 2, 7))
}

