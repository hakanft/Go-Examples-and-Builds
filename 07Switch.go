/*package main

import (
	"fmt"
	"time"
)

func main(){

  i := 2

  fmt.Print("Write ", i," as ")

  switch i{
    case 1:
    fmt.Println("One")
    case 2:
    fmt.Println("Two")
    case 3:
    fmt.Println("Three")
  } 

  switch time.Now().Weekday(){
    case time.Saturday, time.Sunday:
    fmt.Println("It is the weekend")
    default:
    fmt.Println("It is a weekday")
  }

  t := time.Now()

  switch{
    case t.Hour() < 12:
    fmt.Println("It is before noon")
    default:
    fmt.Println("It is after noon")
  }

  WhatAmI := func(i interface{}) {

    switch t := i.(type) {
      case bool:
      fmt.Println("I am bool")
      case int:
      fmt.Println("I am integer")
      default:
      fmt.Printf("Don't know type %T\n", t) 

    }

  }
  WhatAmI(true)
  WhatAmI(1)
  WhatAmI("Hey")
}*/