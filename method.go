package main

import "fmt"

type rect struct{
  width, height int
}

func area(r* rect) int{
  return r.width * r.height
}

func perimiter(r rect) int {
  return 2 * r.width + 2 * r.height
}

func main(){
  r := rect{width: 10, height: 25}

  fmt.Println("area: ", area(&r))
  fmt.Println("perimiter: ", perimiter(r))
}



