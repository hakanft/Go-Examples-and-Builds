/*package main

import "fmt"


func main() {

  messages := make(chan string)
  singnals := make(chan bool)

  select{
    case msg := <- messages:
    fmt.Println("recieved message", msg)
    default:
    fmt.Println("No recieved message") 
  }

  msg := "hi"
  select{

    case messages <- msg:
    fmt.Println("sent message", msg)
    default:
    fmt.Println("no message sent")
  }

  select{
    case msg := <- messages:
    fmt.Println("recieved message", msg)
    case sig := <-singnals:
    fmt.Println("recieved signal", sig)
    default:
    fmt.Println("no activty")
  }
  
}*/