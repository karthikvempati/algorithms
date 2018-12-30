package routine

import (
  "fmt"
  "time"
)

func PrintInt(n int) {
  for i:=0;i<10;i++ {
    fmt.Println(n, ":", i)
  }
}

func main() {
  for i:=0;i<10;i++ {
    go PrintInt(i)
  }
  amt := time.Duration(250)
  time.Sleep(time.Millisecond * amt)
}
