package main

import "fmt"

func fact(n int) int {
  if(n == 0){
    return 1
  }
  return n * fact(n-1)
}

func fibonacci(n int) int{

}

func main(){
  fmt.Println("7!: ", fact(7))
}
