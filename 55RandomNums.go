/*package main

import (
	"fmt"
	"math/rand"
  "time"
)

func main(){
  p1 := fmt.Print
  p2 := fmt.Println

  p1(rand.Intn(100),",")
  p1(rand.Intn(100))
  p2()

  p1((rand.Float64()*5)+5,",")
  p1((rand.Float64() * 5) + 5)
  p2()

  s1 := rand.NewSource(time.Now().UnixNano())
  r1 := rand.New(s1)

  
  p1(r1.Intn(100),",")
  p1(r1.Intn(100))
  p2()

  s2 := rand.NewSource(42)
  r2 := rand.New(s2)
  p1(r2.Intn(100),",")
  p1(r2.Intn(100))
  p2()
  s3 := rand.NewSource(42)
  r3 := rand.New(s3)
  p1(r3.Intn(100),",")
  p1(r3.Intn(100))
  p2()

}*/