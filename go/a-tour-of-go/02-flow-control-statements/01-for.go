/*
  For Loop
    - Go has only one looping construct.
    - Unlike other languages, Go has no parenthesis surrounding
      the three components.

    Three Components
      01. The init statement.
        - Will often be a short variable.
        - The variables declared there are visible only in the scope
          of the for statement.

      02. The condition expression.

      03. The post statement.
*/

package main

import "fmt"

func main() {
  sum := 1
  
  for i := 1; i < 10; i++ {
    sum += i
  }
  
  fmt.Println("sum: ", sum)
}

