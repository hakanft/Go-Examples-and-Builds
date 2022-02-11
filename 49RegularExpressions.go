/*package main

import(
  "fmt"
  "bytes"
  "regexp"
)

func main(){

  var a = fmt.Println

  match, _ := regexp.MatchString("p([a-z]+)ch","peach")
  a(match)

  r, _ := regexp.Compile("p([a-z]+)ch")

  a(r.MatchString("peach"))

  a(r.FindString("peach punch"))

  a("idx:", r.FindStringIndex("peach punch"))

  a(r.FindStringSubmatch("peach punch"))

  a(r.FindStringSubmatchIndex("peach punch"))

  a(r.FindAllString("peach punch pinch", -1))

  a("all:", r.FindAllStringSubmatchIndex("peach punch pinch", -1))
  
  a(r.FindAllString("peach punch pinch", 2))

  a(r.Match([]byte("peach")))

  r = regexp.MustCompile("p([a-z]+)ch")
  a("regexp: ", r)

  a(r.ReplaceAllString("a peach","<fruit>"))

  in := []byte("a peach")
  out := r.ReplaceAllFunc(in, bytes.ToUpper)
  a(string(out))


}*/