/*package main

import(
  "fmt"
  "os"
)

type point struct{
  x,y int
}

func main(){

  var a = fmt.Printf

  p := point{1,2}
  a("struct1: %v\n",p)
  
  a("struct2: %+v\n",p)
  
  a("struct3: %#v\n",p)

  a("type: %T\n",p)

  a("bool: %t\n", true)

  a("int: %d\n", 123)

  a("bin: %b\n",14)

  a("char: %c\n",33)

  a("hex: %x\n",456)

  a("float1: %f\n",78.9)

  a("float2: %e\n", 123400000.0)

  a("float3: %E\n", 123400000.0)

  a("str1: %s\n", "\"string\"")

  a("str2: %q\n", "\"string\"")

  a("str3: %x\n", "hex this")

  a("pointer: %p\n", &p)

  a("width1: |%6d|%6d|\n",12,345)
  
  a("width2: |%6.2f|%6.2f|\n",1.2,3.45)

  a("width3: |%-6.2f|%-6.2f|\n",1.2,3.45)

  a("width4: |%6s|%6s|\n","foo","b")

  a("width5: |%-6s|%-6s|\n","foo","b")

  s := fmt.Sprintf("sprintf: a %s","string")
  fmt.Println(s)

  fmt.Fprintf(os.Stderr, "io: an %s\n", "error")


}*/