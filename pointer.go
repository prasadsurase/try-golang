package main

import "fmt"

func zeroVal(iVal int){
  iVal = 10
}

func zeroPtr(iPtr *int) {
  *iPtr = 99
}

func main(){
  i := 1
  fmt.Println("initial: ", i)

  zeroVal(i)
  fmt.Println("value : ", i)

  zeroPtr(&i)
  fmt.Println("pointer: ", i)

  fmt.Println("pointer: ", &i)
}
