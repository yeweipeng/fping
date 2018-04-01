package main

import (
  "os"
  "fmt"
  "bufio"
  "log"
)

func main() {
  ip_file, err := os.Open("./ip")
  if err != nil {
    log.Fatal(err)
  }
  defer ip_file.Close()

  scanner := bufio.NewScanner(ip_file)
  for scanner.Scan() {
    fmt.Println(scanner.Text())
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }
}
