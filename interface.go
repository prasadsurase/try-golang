package main

import "fmt"
import "math"

type geometry interface {
  area() float64
  perimeter() float64
}

type rectangle struct {
  width, height float64
}

type circle struct {
  radius float64
}

func area(r rectangle) float64 {
  return r.width * r.height
}

func perimeter(r rectangle) float64 {
  return 2 * r.width + 2 * r.height
}

func area(c circle) float64 {
  return 2 * math.Pi * c.radius * c.radius
}

func perimeter(c circle) float64 {
  return 2 * math.Pi * c.radius
}

func measure(g geometry) {
  fmt.Println(g)
  fmt.Println(g.area())
  fmt.Println(g.perimeter())
}

func main(){
  r := rectangle{ width: 10.5, height: 20.7 }
  c := circle{ radius: 9.75 }
  fmt.Println("r.area(): ", r.area())
  fmt.Println("r.perimeter(): ", r.perimeter())
  measure(r)
  measure(c)
}


