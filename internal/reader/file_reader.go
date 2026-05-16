package reader

import (
	"fmt"
	"os"
  "bufio"
)

func ReadFile(file_path string) ([]byte, error) {
  data, err := os.ReadFile(file_path)

  if err != nil {
    fmt.Printf("[Error]: File read error %s\n", err)
    return nil, err
  }

	return data, nil
}

func CountFileLines(file_path string) (int ,error){
  
  var lines int

  file, err := os.Open(file_path)
  if err != nil {
    fmt.Printf("[Error]: CountFileLines -> Open file error: %s\n", err)
    return 0, err
  }

  defer file.Close()
  
  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    lines++
  }

  if err := scanner.Err(); err != nil {
    fmt.Printf("[Error]: CountFileLines -> Scanner error: %s\n ", err)
    return 0, err
  }

  return lines, nil

}
