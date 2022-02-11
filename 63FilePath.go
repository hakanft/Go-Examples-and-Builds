/*package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

var a = fmt.Println

func main(){
  p := filepath.Join("dir1","dir2","filename")
  a("p:",p)

  a(filepath.Join("dir1//","filename"))
  a(filepath.Join("dir1/../dir1","filename"))

  a("Dir(p):", filepath.Dir(p))
  a("Base(p):", filepath.Base(p))

  a(filepath.IsAbs("dir/file"))
  a(filepath.IsAbs("/dir/file"))

  filename := "config.json"

  ext := filepath.Ext(filename)
  a(ext)

  a(strings.TrimSuffix(filename, ext))

  rel, err := filepath.Rel("a/b", "a/b/t/file")
  if err != nil{
    panic(err)
  }
  a(rel)

  rel, err = filepath.Rel("a/b", "a/c/t/file")
  if err != nil{
    panic(err)
  }
  a(rel)
  
}*/