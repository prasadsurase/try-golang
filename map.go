package main

import "fmt"

func main(){
  m := make(map[string]int)

  fmt.Println("map: ", m)

  m["a"] = 1
  m["b"] = 2
  fmt.Println("map: ", m)

  v1 := m["a"]
  fmt.Println("m['a']: ", v1)

  n := map[string]int{"foo": 1, "bar": 2}
  fmt.Println("map: ", n)
}
