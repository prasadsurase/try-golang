package main

import "fmt"

func main(){
  nums := []int{ 1, 2, 3 }
  sum := 0

  for _, num := range nums {
    sum += num
  }

  fmt.Println("Sum: ", sum)

  kvs := map[string]string{"a": "temple", "b": "run"}
  for k, v := range kvs {
    fmt.Println("K: ", k, "V:", v)
  }
}
