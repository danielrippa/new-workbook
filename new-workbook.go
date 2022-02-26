package main

import (
  "os"
  "fmt"
  "github.com/xuri/excelize/v2"
)

func main() {
  args := os.Args[1:]

  sheets := os.Args[2:]

  f := excelize.NewFile()

  for _, name := range sheets {
    f.NewSheet(name)
  }

  err := f.SaveAs(args[0])
  if err != nil {
    fmt.Println(err)
  }
}