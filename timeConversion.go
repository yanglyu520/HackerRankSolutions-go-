package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
    "strconv"
)
 
 // consider adding error as output variables in production env
 // consider more validations of time input and output and trim string in production env

func timeConversion(s string) string {
    // divide time string input into hh, mm, ss, suffix
    tokens := strings.Split(s,":")
    hhStr := tokens[0]
    mm := tokens[1]
    ss := tokens[2][:2] 
    suffix := tokens[2][2:]
    
    // convert hh from string to int
    hhInt, err := strconv.Atoi(hhStr)
    if err != nil {
     return ""
}
    newHhInt:=hhInt
    switch suffix {
      case "AM":
        if hhInt == 12 {
          newHhInt = 0
        }
      case "PM":
        if hhInt != 12 {
            newHhInt = hhInt+12
        }
       default:
          return ""
    }
    
    return fmt.Sprintf("%02d:%s:%s", newHhInt, mm, ss)
}
