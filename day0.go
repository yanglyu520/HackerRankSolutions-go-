package main
import (
    "fmt"
    "bufio"
    "os"
    "log")

func main() {
 // Read input from STDIN. Print output to STDOUT
  reader := bufio.NewReader(os.Stdin)
  
  input, err := reader.ReadString('\n')
  if err != nil {
    log.Fatal(err)
}
 
  fmt.Println(input)
}
