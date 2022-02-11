/*package main

import(
  "fmt"
  "encoding/json"
  "os"

)

type response1 struct{
  Page int
  Fruits []string
}

type response2 struct{
  Page int `json: "page"`
  Fruits []string `json: "fruits"`
}

func main(){

  var a = fmt.Println

  bolB,_ := json.Marshal(true)
  a(string(bolB))

  intB,_ := json.Marshal(1)
  a(string(intB))

  fltB,_ := json.Marshal(2.34)
  a(string(fltB))

  strB,_ := json.Marshal("gopher")
  a(string(strB))

  slcD := []string{"apple","peach","pear"}
  slcB,_ := json.Marshal(slcD)
  a(string(slcB))

  mapD := map[string]int{"apple":5, "lettuce": 7}
  mapB, _ := json.Marshal(mapD)
  a(string(mapB))

  res1D := &response1{
    Page: 1,
    Fruits: []string{"apple","peach","pear"}}
  res1B,_ := json.Marshal(res1D)
  a(string(res1B))

  res2D := &response2{
    Page: 1,
    Fruits: []string{"apple","peach","pear"}}
  res2B,_ := json.Marshal(res2D)
  a(string(res2B))

  byt := []byte(`{"num":6.13, "strs":["a","b"]}`)

  var dat map[string]interface{}

  if err := json.Unmarshal(byt,&dat); err != nil{
    panic(err)
  }
  a(dat)
  
  num := dat["num"].(float64)
  a(num)

  strs := dat["strs"].([]interface{})
  str1 := strs[0].(string)
  a(str1)

  str := `{"page": 1, "fruits": ["apple","peach"]}`
  res := response2{}
  json.Unmarshal([]byte(str),&res)
  a(res)
  a(res.Fruits[0])

  enc := json.NewEncoder(os.Stdout)
  d := map[string]int{"apple":5, "lettuce":7}
  enc.Encode(d)

}*/