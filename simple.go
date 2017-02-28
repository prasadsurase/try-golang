package main
import "fmt"
import "math"

const (
  Pi = 3.14
  Truth = true
)

func main(){
  hello := func(){
    fmt.Println("Hola World")
  }

  name, location := "Prasad", "Pune"
  age := 30
  fmt.Printf("%s (%d) of %s", name, age, location)
  fmt.Println(Pi)
  fmt.Println(Truth)
  hello()

  fmt.Println(math.Pi)
}
