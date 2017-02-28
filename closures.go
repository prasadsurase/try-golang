package main

import "fmt"

func intSeq() func() int {
  fmt.Println("hello")
  i := 0
  return func() int {
    i += 1
    return i
  }
}

func main() {
  fmt.Println("hello")
  nextInt := intSeq()
  fmt.Println(nextInt())
  fmt.Println(nextInt())
  fmt.Println(nextInt())
  newInts := intSeq()
  fmt.Println(newInts())
}

