package main

import "fmt"
import "math"

var test = "testing"
const s string = "i am an constant"

func main(){
  fmt.Println(test)

  fmt.Println(s)

  const n = 5000000000

  const d = 3e20 / n
  fmt.Println(d)

  fmt.Println(int64(d))
  fmt.Println(math.Sin(n))
}
