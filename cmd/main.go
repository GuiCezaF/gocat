package main

import (
	"fmt"
	"os"

  "github.com/GuiCezaF/gocat/internal/reader"
  	
)

func main(){
  file_path := "./test.txt"
  data, err := reader.ReadFile(file_path)
  
  if err != nil {
    os.Exit(1)
  }
  fmt.Println(string(data))

  n_lines, err := reader.CountFileLines(file_path)
  fmt.Println(n_lines)
}
