package main

import "fmt"

// declare interface named Stringer that declares toString function
type Stringer interface {
  toString() string
}

//
type fakeString struct {
  content string
}

func (s *fakeString) toString() string {
  return s.content
}

func printString(value interface{}){
  switch str := value.(type) {
    case string:
      fmt.Println(str)
    case Stringer:
      fmt.Println(str.toString())
  }
}

func main(){
  s := fakeString{content: "Hello World"}
  printString(s)
  printString(&s)
  printString("Hola!!!")
}
