/*
  Numeric Constants
    - Are high precision values.
  
  Note:
    - An untyped constants takes the type needed by its context.
    - An "int" can store at maximum a 64-bit integer, and sometimes less.
*/

package main

import "fmt"
import "math/bits"


const (
  // Shifting 1 bit left 62 places.
  // The binary number that is 1 followed by 62 zeroes.
  Big = 1 << 62

  // Shifting 63 bits right 61 places.
  // Preceded by 61 zeroes.
  Small = Big >> 61
)

// The fmt package allows us to print anything on the screen.
// The bits package allows to perform bitwise operations to unsigned integer values.
func getBits(number uint) int {
   // getting the binary representation of the above chosen value
   fmt.Printf("Binary representation of %d is: %b\n", number, number)

   // initializing a variable to store the results
   var result int

   // getting the number of bits in the result variable
   result = bits.Len(number)

   // returning back the number of bits
   return result
}

func main() {
  fmt.Println(getBits(Big));
  fmt.Println(getBits(Small))
}

