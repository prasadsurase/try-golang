package main

import "fmt"

func hello(s string){
  for i:= 0; i < 3; i++ {
    fmt.Println(s, ": ", i)
  }
}

func main(){
  hello("prasad")

  go hello("goroutine")

  go func(msg string){
    fmt.Println(msg)
  }("going")

  var input string
  fmt.Scanln(&input)
  fmt.Println("done")
}
