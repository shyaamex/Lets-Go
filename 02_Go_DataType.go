/*
	Go Datatype & Variable


Go is statically typed, meaning that ONCE a variable type is DEFINED, it can ONLY store data of that type.


Go has mainly three Datatypes:
- String
- Numeric - Float & int
- Boolean




*/

package main
import ("fmt")

func main() {
  var a bool = true     // Numeric-  Boolean
  var b int = 5         // Numeric-  Integer
  var c float32 = 3.14  // Floating point number
  var d string = "Hi!"  // String

  fmt.Println("Boolean: ", a)
  fmt.Println("Integer: ", b)
  fmt.Println("Float:   ", c)
  fmt.Println("String:  ", d)
}

