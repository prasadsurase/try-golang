package main

import "fmt"
import "time"

func main(){
  i := 2
  switch i {
    case 1:
      fmt.Println("One")
    case 2:
      fmt.Println("Two")
    case 3:
      fmt.Println("Three")
  }

  switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
      fmt.Println("Its Weekend!!!")
    default:
      fmt.Println("FML")
  }

  t := time.Now()

  switch {
    case t.Hour() < 12:
      fmt.Println("Morning")
    case t.Hour() > 12:
      fmt.Println("Evening")
  }

  whatAmI := func(i interface{}){
    switch t := i.(type){
      case int:
        fmt.Println("Integer")
      case bool:
        fmt.Println("Boolean")
      default:
        fmt.Println("Kuch to hoga ", t)
    }
  }

  whatAmI(true)
  whatAmI(1)
  whatAmI("hola")
}
