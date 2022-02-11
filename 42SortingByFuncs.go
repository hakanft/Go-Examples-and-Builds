/*package main

import(
  "fmt"
  "sort"
)

type byLenght []string

func(s byLenght) Len() int{
  return len(s)
}

func (s byLenght) Swap(i,j int){
  s[i], s[j] = s[j], s[i]

}

func(s byLenght) Less(i, j int) bool {
  return len(s[i]) <len(s[j])

}

func main(){

  fruits := []string{"peach", "banana", "kiwi"}
  sort.Sort(byLenght(fruits))
  fmt.Println(fruits)

}*/