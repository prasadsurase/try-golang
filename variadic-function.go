package main

import "fmt"

func sum(nums ...int) (int) {
  total := 0
  for _, num := range nums {
    total += num
  }
  return total
}

func main(){
  fmt.Println("Sum: ", sum(1, 2, 3, 4, 5))
  fmt.Println("Sum: ", sum(3, 4, 5))
  nums := []int{ 1, 3, 5, 7, 9, 11 }
  fmt.Println("Sum: ", sum(nums ...))
}
