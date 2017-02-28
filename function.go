package main
import "fmt"

func location(city string) (string, string){
  var region, continent string

    switch(city){
      case "Pune":
        region, continent = "India", "Asia"
      case "NY":
          region, continent = "USA", "America"
      default:
            region, continent = "World", "Earth"
    }
  return region, continent
}

func locatn(name, city string) (region, continent string) {
  switch city {
    case "New York", "LA", "Chicago":
      continent = "North America"
    default:
        continent = "Unknown"
  }
  return
}

func main(){
  regionn, continentt := locatn("Matt", "LA")
    fmt.Printf("%s lives in %s", regionn, continentt)
}

