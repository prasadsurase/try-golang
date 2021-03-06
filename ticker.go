package main

import "fmt"
import "time"

func main(){
  ticker := time.NewTicker(time.Millisecond * 500)

  go func(){
    for t:= range ticker.C {
      fmt.Println("Tick: ", t)
    }
  }()

  time.Sleep(time.Millisecond * 2600)
  ticker.Stop()
  fmt.Println("Ticker Stopped")
}
