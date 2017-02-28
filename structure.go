package main
import "fmt"

type User struct {
  Name string
  Age int
}

func newKid(u User) int {
  u.Age++
  return u.Age
}

func newWife(u *User) int {
  u.Age++
  return u.Age
}

func main(){
  me := User{Name: "Prasad", Age: 30}
  fmt.Printf("%s is %d years old\n", me.Name, newKid(me))
  fmt.Printf("%s is actually %d years old \n", me.Name, me.Age)
  fmt.Printf("%s is actually %d years old", me.Name, newWife(&me))
}
