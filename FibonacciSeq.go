package main

import (
"fmt"
"os"
"strconv")


func main() {
arg := os.Args
if len(arg) != 2 {
fmt.Println("No argument given, Expecting one argument as size of series")
os.Exit(1)
}
fmt.Println("aa", len(arg))
val, err := strconv.Atoi(arg[1])
if err != nil {
fmt.Println("Invalid argument received, expected integer, received",arg[1])
os.Exit(1)
}
f := Fibonacci()
for i := 0 ; i < val ; i++{
fmt.Println(f())
}
}

//Fibonacci return a clousre which handles the fibonacci sequence at ith index & this will work for int32 as limit
func Fibonacci() func() int {
a, b := -1, 1
return func() int {
a, b = b, a+b
return b
}
}

//To run the code : " go run filename.go count" where filename.go is filename of your code and count must be integer value wihtout "", eg: go run temp.go 10,here 10 is size of series