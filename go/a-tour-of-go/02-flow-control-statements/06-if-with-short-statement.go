/*
  If
    - Like "for", the "if" statement can start with a short
      statement to execute before the condition.
    - Variables declared by the statement are only in scope
      until the end of the "if".
*/

package main

import (
  "fmt"
  "math"
)

// If with standard flow.
func standard(x, n, limit float64) float64 {
  var v float64 = math.Pow(x, n)

  if v < limit {
    return v
  }

  return limit
}

// If with short statement.
func pow(x, n, limit float64) float64 {
  if v := math.Pow(x, n); v < limit {
    return v
  }

  return limit
}

func main() {
  fmt.Println("With Standard Statement: ", standard(3, 2, 10))
  fmt.Println("With Short Statement: ", pow(3, 2, 10))
}

