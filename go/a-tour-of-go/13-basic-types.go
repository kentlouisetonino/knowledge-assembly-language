/*
  Go's Basic Types:
    - bool
    - string
    - int int8 int16 int32 int64
      uint uint8 uint16 uint32 uint64 uintptr
    - byte (alias for uint8)
    - rune (alias for int32)
           (represents a Unicode code point)
    - float32 float64
    - complex64 complex128

  The "int", "uint", and "uintptr" types are usually 32bits wide on 32b-bit
  system and 64 bits wide on 64-bit systemss. When you need an integer value
  you should use "int" unless you have a specific reason to use a sized or
  unsigned integer type.

  The example below shows variables of several types, and also that variable
  declarations may be "factored" into blocks, as with import statements.

  Note
   %T : Getting the type.
   %v : Getting the value.
*/

package main

import (
  "fmt"
  "math/cmplx"
)

var (
  ToBe bool = false
  MaxInt uint64 = 1 << 64 - 1
  z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
  fmt.Printf("Type: %T, Value: %v \n", ToBe, ToBe)
  fmt.Printf("Type: %T, Value: %v \n", MaxInt, MaxInt)
  fmt.Printf("Type: %T, Value: %v \n", z, z)
}

